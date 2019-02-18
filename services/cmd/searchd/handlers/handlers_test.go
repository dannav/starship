package handlers_test

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/dannav/starship/services/cmd/searchd/handlers"
	"github.com/dannav/starship/services/cmd/searchd/healthcheck"
	"github.com/dannav/starship/services/internal/embedding"
	"github.com/dannav/starship/services/internal/platform/db"
	"github.com/dannav/starship/services/internal/platform/web"
	"github.com/dannav/starship/services/internal/shared"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // postgres sql driver
	"github.com/pkg/errors"
)

const (
	// readTimeout is timeout for reading the request
	readTimeout = 30 * time.Second

	// writeTimeout is timeout for reading the response
	writeTimeout = 30 * time.Second

	// shutdownTimeout is the timeout for shutdown
	shutdownTimeout = 30 * time.Second

	// defaultAppPort is the port that the application listens on if no PORT env is set
	defaultAppPort = 8080

	// vDims is how many dimensions the ml model result embedding is
	vDims = 512

	// modelEndpoint is the endpoint to the ml model
	modelEndpoint = "/v1/models/modelname"

	// postgresDockerImage is the postgres docker image to use
	postgresDockerImage = "postgres:11.0-alpine"
)

// make app and db available to all tests in this package
var ts testSuite

type testSuite struct {
	App *handlers.App
	DB  *sqlx.DB
}

func TestMain(m *testing.M) {
	// if there was an error, exit main with non-zero status code.
	var mainErr error
	defer func() {
		if mainErr != nil {
			log.Print(mainErr.Error())
			os.Exit(1)
		}
	}()

	// create a mock tikad service
	tikad := tikadMock()
	tikadURL := tikad.URL
	defer tikad.Close()

	// set storage path to project directory + testdata
	pwd, err := os.Getwd()
	if err != nil {
		mainErr = errors.Wrap(err, "setting storage path")
		return
	}

	// create storage path in root of services repo
	storagePath := filepath.Join(pwd, "../../../", "testdata")

	// create storagePath deepest dir (indexes)
	err = os.MkdirAll(filepath.Join(storagePath, "indexes"), os.ModePerm)
	if err != nil {
		mainErr = errors.Wrap(err, "could not create storage directory")
		return
	}

	// setup ml api mock
	mlapi := mlAPIMock()
	modelURL := mlapi.URL + modelEndpoint
	defer mlapi.Close()

	_, err = healthcheck.ModelReady(modelURL)
	if err != nil {
		mainErr = errors.Wrap(err, "ml model api did not return ready response")
		return
	}

	// start a postgres container
	cleanupDocker, err := startPostgresDocker()
	if err != nil {
		mainErr = errors.Wrap(err, "starting postgres container")
		return
	}

	// create new postgres connection
	dba, err := sqlx.Open("postgres", "user=postgres password=password host=localhost dbname=postgres sslmode=disable")
	if err != nil {
		mainErr = errors.Wrap(err, "open postgres database")
		return
	}

	// wait till we get a connection to the database started with docker
	var connected bool
	if err := dba.Ping(); err != nil {
		for {
			log.Print("waiting for database connection")
			ticker := time.NewTicker(time.Second)

			select {
			case <-ticker.C:
				if err := dba.Ping(); err == nil {
					connected = true
				}
			}

			if connected {
				break
			}
		}
	}

	// create db schema
	if _, err := dba.Exec(db.Schema); err != nil {
		mainErr = errors.Wrap(err, "creating database schema")
		return
	}

	// create db indexes
	if _, err := dba.Exec(db.Indexes); err != nil {
		mainErr = errors.Wrap(err, "creating database indexes")
		return
	}

	// run db seed
	for _, s := range db.Seed {
		if _, err := dba.Exec(s); err != nil {
			mainErr = errors.Wrap(err, "seeding database")
			return
		}
	}

	// force 30 second timeouts on all http requests
	client := &http.Client{
		Timeout: time.Second * 30,
	}

	cfg := handlers.Cfg{
		ModelURL:            modelURL,
		ModelVectorDims:     vDims,
		TikaURL:             tikadURL,
		StoragePath:         storagePath,
		IndexKey:            "",
		AccessKey:           "",
		ObjectStorageConfig: handlers.ObjectStorageCfg{}, // TODO :- test object storage integration
	}

	// create the app so we can get access to handlers
	app := handlers.NewApp(cfg, dba, client)

	// add to global test suite so we can use in handler tests
	ts.App = app
	ts.DB = dba

	// run all tests
	code := m.Run()

	// cleanup docker containers
	if err := cleanupDocker(); err != nil {
		mainErr = errors.New("could not clean up postgres docker container")
		return
	}

	// cleanup storageDir path
	if err := os.RemoveAll(storagePath); err != nil {
		mainErr = errors.New(err.Error())
		return
	}

	os.Exit(code)
}

// tikadMock starts a mocked tikad service
func tikadMock() *httptest.Server {
	parseHandler := func(w http.ResponseWriter, r *http.Request) {
		resp := shared.TikaResponse{
			Body:         "hello world",
			DocumentType: "text/plain",
		}

		if r.URL.Path == "/v1/parse" {
			web.Respond(w, r, http.StatusOK, resp)
		} else {
			web.RespondError(w, r, http.StatusNotFound, errors.New("not found"))
		}
	}

	srv := httptest.NewServer(http.HandlerFunc(parseHandler))

	return srv
}

// mlAPIMock starts a mocked ml model api service
func mlAPIMock() *httptest.Server {
	parseHandler := func(w http.ResponseWriter, r *http.Request) {
		type modelBody struct {
			Version string `json:"version"`
			State   string `json:"state"`
		}

		resp := struct {
			Results []modelBody `json:"model_version_status"`
		}{
			Results: []modelBody{
				modelBody{
					State: "AVAILABLE",
				},
			},
		}

		// fake sentence embedding with 512 dimensions
		var em embedding.ModelResponse
		var emResult []float32
		for i := 0; i < vDims; i++ {
			emResult = append(emResult, 0)
		}

		em.Outputs = [][]float32{
			emResult,
		}

		if r.URL.Path == "/v1/models/modelname" {
			b, err := json.Marshal(resp)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write(b)
		} else if r.URL.Path == "/v1/models/modelname:predict" {
			b, err := json.Marshal(em)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write(b)
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("not found"))
		}
	}

	srv := httptest.NewServer(http.HandlerFunc(parseHandler))

	return srv
}

// startPostgresDocker starts a new container running postgres and binds it to localhost
// the postgres container is started with user=postgres password=password
func startPostgresDocker() (func() error, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return nil, err
	}

	cfg := &container.Config{
		Image: postgresDockerImage,
		ExposedPorts: nat.PortSet{
			"5432": struct{}{},
		},
		Env: []string{
			"POSTGRES_PASSWORD=password",
		},
	}

	hostCfg := &container.HostConfig{
		PortBindings: map[nat.Port][]nat.PortBinding{
			nat.Port("5432"): {
				{
					HostIP:   "127.0.0.1",
					HostPort: "5432",
				},
			},
		},
	}

	ctx := context.Background()
	resp, err := cli.ContainerCreate(ctx, cfg, hostCfg, nil, "postgres_test")
	if err != nil {
		return nil, err
	}

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return nil, err
	}

	cleanup := func() error {
		// wait 5s to stop container
		duration, _ := time.ParseDuration("5s")
		if err := cli.ContainerStop(ctx, resp.ID, &duration); err != nil {
			return err
		}

		// remove the container and volume data
		opts := types.ContainerRemoveOptions{
			RemoveVolumes: true,
		}

		if err := cli.ContainerRemove(ctx, resp.ID, opts); err != nil {
			return err
		}

		return nil
	}

	return cleanup, nil
}
