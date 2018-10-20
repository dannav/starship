package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/google/go-github/v18/github"
	"github.com/pkg/errors"
	stripmd "github.com/writeas/go-strip-markdown"
)

var helpText = `
This application downloads readmes from top github repositories to use as test data for building
a semantic search index.

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

	client := github.NewClient(nil)
	searchCtx := context.Background()

	// search for $$lang repos with highest stars
	query := fmt.Sprintf("language:%v&sort=stars", lang)
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

	log.Printf("Processing %v repositories.\n", len(result.Repositories))
	var readmes []Readme

	// get readme for each repository and save to data dir
	for i := 0; i < len(result.Repositories); i++ {
		r := result.Repositories[i]

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
			// convert markdown to plain text
			stripped := stripmd.Strip(content)

			readme := Readme{
				Repo:    fmt.Sprintf("%v/%v", owner, repoName),
				Content: stripped,
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
}

// HandleRateLimit sleeps until the github API rate limit resets
func HandleRateLimit(resp *github.Response) {
	log.Printf("Hit rate limit. Waiting until %v to resume operation.\n", resp.Rate.Reset.Time)
	until := time.Until(resp.Rate.Reset.Time)
	time.Sleep(until)
}
