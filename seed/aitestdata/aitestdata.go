package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/go-github/github"
	"github.com/pkg/errors"
	"golang.org/x/oauth2"
)

var helpText = `
This application downloads readmes from top github repositories and indexes them
with the searchd api to use as test data for building a semantic search index.

Usage:
	testdata [arguments]

Example:
	testdata -l javascript -o ./test-data

The following flags are required:
	-l		search top repositories on github tagged with this language
	-o		the relative path to save data files
	-h		show this help text

Possible arguments are:
	help 	show this help text
`

// Readme represents the readme file in a repository
type Readme struct {
	Repo    string `json:"repo"`
	Content string `json:"content"`
}

var client = &http.Client{
	Timeout: time.Second * 30,
}

func main() {
	var lang, outputDir string
	var showHelp bool

	flag.StringVar(&lang, "l", "", "search top repositories on github tagged with this language")
	flag.StringVar(&outputDir, "o", "", "the relative path to save data files")
	flag.BoolVar(&showHelp, "h", false, "show help text")
	flag.Parse()

	var cmd string
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}

	if cmd == "help" || cmd == "Help" || showHelp {
		fmt.Printf("%v\n", helpText)
		return
	}

	if lang == "" || outputDir == "" {
		fmt.Println("Language and output directory required. Run `testdata help` for more info.")
		return
	}

	// authenticate with github oauth2 api
	token := os.Getenv("TOKEN")
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// search first 4 pages for $$lang repos with highest stars
	var results []github.Repository

	query := fmt.Sprintf("language:%v&sort=stars", lang)
	result, resp := RunSearch(client, query)
	if resp.NextPage != resp.LastPage {
		query = fmt.Sprintf("%v&page=%v", query, resp.NextPage)
	}
	results = append(results, result.Repositories...)

	result2, resp := RunSearch(client, query)
	if resp.NextPage != resp.LastPage {
		query = fmt.Sprintf("%v&page=%v", query, resp.NextPage)
	}
	results = append(results, result2.Repositories...)

	result3, resp := RunSearch(client, query)
	if resp.NextPage != resp.LastPage {
		query = fmt.Sprintf("%v&page=%v", query, resp.NextPage)
	}
	results = append(results, result3.Repositories...)

	result4, resp := RunSearch(client, query)
	if resp.NextPage != resp.LastPage {
		query = fmt.Sprintf("%v&page=%v", query, resp.NextPage)
	}
	results = append(results, result4.Repositories...)
	// finish searching

	log.Printf("Processing %v repositories.\n", len(results))
	var readmes []Readme

	// get readme for each unique repository and save to data dir
	uniqueRepo := map[string]bool{}

	for i := 0; i < len(results); i++ {
		r := results[i]

		var owner string
		var repoName string
		var user *string

		user = r.Owner.Name
		if user != nil {
			owner = *user
		} else {
			user := r.Owner.Login
			if user != nil {
				owner = *user
			}
		}

		repo := r.Name
		if repo != nil {
			repoName = *repo
		} else {
			repoName = ""
		}

		readmeCtx := context.Background()
		readme, resp, err := client.Repositories.GetReadme(readmeCtx, owner, repoName, nil)
		if err != nil {
			if _, ok := err.(*github.RateLimitError); ok {
				HandleRateLimit(resp)

				// rerun current request
				if i == 0 {
					i = -1
				} else {
					i = i - 1
				}

				continue
			} else {
				log.Fatal(errors.Wrap(err, "getting readme from repo"))
			}
		}

		content, err := readme.GetContent()
		if err != nil {
			log.Fatal(errors.Wrap(err, "getting readme content"))
		}

		// create isolated scope
		{
			readme := Readme{
				Repo:    fmt.Sprintf("%v/%v", owner, repoName),
				Content: content,
			}

			if _, ok := uniqueRepo[readme.Repo]; ok {
				continue
			} else {
				uniqueRepo[readme.Repo] = true
			}

			readmes = append(readmes, readme)
		}
	}

	// get pwd of this program
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(errors.Wrap(err, "error getting pwd"))
	}

	// pretty marshal readmes to json
	b, err := json.MarshalIndent(readmes, "", "  ")
	if err != nil {
		log.Fatal(errors.Wrap(err, "error marshaling data into json"))
	}

	// write to json file using filepath given from argument
	saveFilePath := outputDir
	dir := filepath.Join(wd, saveFilePath, lang+"-readmes.json")
	err = ioutil.WriteFile(dir, b, 0644)
	if err != nil {
		log.Fatal(errors.Wrap(err, "error writing file"))
	}

	for _, r := range readmes {
		err := Index(strings.NewReader(r.Content), r.Repo)
		if err != nil {
			log.Print(err)
			log.Print("error indexing README " + r.Repo)
			continue
		}
	}
}

// RunSearch runs a github repository search
func RunSearch(client *github.Client, query string) (*github.RepositoriesSearchResult, *github.Response) {
	searchCtx := context.Background()
	result, resp, err := client.Search.Repositories(searchCtx, query, nil)
	if err != nil {
		if _, ok := err.(*github.RateLimitError); ok {
			HandleRateLimit(resp)

			// rerun request
			result, _, _ = client.Search.Repositories(searchCtx, query, nil)
		} else {
			log.Fatal(errors.Wrap(err, "searching repositories"))
		}
	}

	return result, resp
}

// HandleRateLimit sleeps until the github API rate limit resets
func HandleRateLimit(resp *github.Response) {
	log.Printf("Hit rate limit. Waiting until %v to resume operation.\n", resp.Rate.Reset.Time)
	until := time.Until(resp.Rate.Reset.Time)
	time.Sleep(until)
}

// Index indexes a readme file with the searchd api
func Index(file io.Reader, path string) error {
	// add multipart form fields
	var buf bytes.Buffer
	encoder := multipart.NewWriter(&buf)
	field, err := encoder.CreateFormFile("content", "README.md")
	if err != nil {
		err = errors.Wrap(err, "creating content form field for searchd request")
		return err
	}

	_, err = io.Copy(field, file)
	if err != nil {
		err = errors.Wrap(err, "copying file to searchd request")
		return err
	}

	pathField, err := encoder.CreateFormField("path")
	if err != nil {
		err = errors.Wrap(err, "creating path form field for index request")
		return err
	}

	_, err = pathField.Write([]byte(path))
	if err != nil {
		err = errors.Wrap(err, "writing index path to index request")
		return err
	}
	encoder.Close()

	// perform index request
	endpoint := "http://localhost:8080/v1/index"
	req, err := http.NewRequest(http.MethodPost, endpoint, &buf)
	if err != nil {
		err = errors.Wrap(err, "preparing searchd request")
		return err
	}
	req.Header.Set("Content-Type", encoder.FormDataContentType())

	res, err := client.Do(req)
	if err != nil {
		err = errors.Wrap(err, "performing searchd request")
		return err
	}

	if res.StatusCode != http.StatusNoContent {
		err = errors.New("searchd request failed")
		return err
	}

	return nil
}
