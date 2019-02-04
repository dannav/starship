package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/dannav/starship/cli/internal/engine"
	"github.com/dannav/starship/cli/internal/help"
	"github.com/dannav/starship/cli/internal/log"
	"github.com/dannav/starship/cli/internal/output"
	"github.com/dannav/starship/cli/internal/uerrors"
	tb "github.com/nsf/termbox-go"
)

// Add timeout of 60secs on web requests
var client = &http.Client{
	Timeout: time.Second * 60,
}

func main() {
	var cmd string
	var args []string

	// flags
	var showHelp bool
	var indexPath string

	flag.BoolVar(&showHelp, "h", false, "show help text")
	flag.BoolVar(&showHelp, "help", false, "show help text")
	flag.StringVar(&indexPath, "p", "/", "path to store document in index")
	flag.Parse()

	args = flag.Args()

	// if help flags passed show help and exit
	if showHelp || len(args) == 0 {
		help.ShowHelp()
		return
	}

	// parse command and arguments
	args = args[1:]
	cmd = flag.Arg(0)

	// check to see if the API is functioning
	e := engine.NewEngine(client)
	ready := e.Ready()
	if ready != true {
		log.Error(uerrors.APINotAvailable)
		return
	}

	// cli main logic
	switch cmd {
	case "help":
		var subcommand string
		if len(args) > 0 {
			subcommand = args[0]
		}

		switch subcommand {
		case "index":
			help.ShowIndex()
			return
		case "search":
			help.ShowSearch()
			return
		}

		help.ShowHelp()
		return
	case "index":
		// ensure a path was provided
		if len(args) == 0 {
			log.Error(uerrors.IndexPathNotProvided)
			return
		}

		// format index path to always have leading '/'
		if indexPath[0] != '/' {
			indexPath = fmt.Sprintf("/%v", indexPath)
		}

		// index file passed in argument
		filePath := args[0]

		// check if path provided exists
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			log.Error(uerrors.IndexFileDoesNotExist)
			return
		}

		// create valid path /foo/bar/readme.md
		var path string
		if indexPath[len(indexPath)-1] != '/' {
			path = fmt.Sprintf("%v/%v", indexPath, filepath.Base(filePath))
		} else {
			path = fmt.Sprintf("%v%v", indexPath, filepath.Base(filePath))
		}

		// check if a file exists at this index path with the same name
		exists, err := e.ExistsAtIndexPath(path)
		if err != nil {
			log.Error(err.Error())
			return
		}

		// if file with same name exists at index path ask if want to overwrite it, cancel index if no
		if exists {
			fmt.Printf(" Overwrite file that exists at index path: %v (y/n) ", path)

			var input string
			fmt.Scanln(&input)

			switch input {
			case "y", "yes", "Y", "YES":
				break
			case "n", "no", "NO", "No":
				return
			default:
				return
			}
		}

		// start a spinner while we index the file
		s := spinner.New(spinner.CharSets[24], 100*time.Millisecond)
		s.Prefix = " Indexing " + filePath + " ... "
		s.Start()

		err = e.Index(filePath, filepath.Base(filePath), indexPath)
		if err != nil {
			log.Error(err.Error())
			return
		}
		s.Stop()

		return
	case "search":
		// ensure search text was provided
		if len(args) == 0 {
			log.Error(uerrors.SearchNoTextGiven)
			return
		}

		searchText := strings.Join(args, " ")

		// start a spinner while we perform a search
		s := spinner.New(spinner.CharSets[24], 100*time.Millisecond)
		s.Prefix = "Searching ... "
		s.Start()

		searchResults, err := e.Search(searchText)
		if err != nil {
			log.Error(uerrors.SearchRequestFailed)
			return
		}
		s.Stop()

		if len(searchResults.Documents) == 0 {
			log.Error(uerrors.SearchNoResults)
			return
		}

		for i := range searchResults.Documents {
		searchLoop:
			// create termbox session so we can read single key
			err = tb.Init()
			if err != nil {
				log.Error(err.Error())
				return
			}

			// trim leading whitespace from search result text
			searchResults.Documents[i].Text = strings.TrimSpace(searchResults.Documents[i].Text)

			// add new line after words around 75 chars to keep terminal output from being too long
			var text string
			words := strings.Fields(searchResults.Documents[i].Text)
			letterCount := 0
			for _, w := range words {
				letterCount = letterCount + len(w)

				if letterCount-75 >= 0 {
					text = text + "\n"
					letterCount = 0
				}

				text = text + " " + w
			}
			searchResults.Documents[i].Text = text

			// clean indexPath of search result
			path := searchResults.Documents[i].Path

			// replace all '.' with '/'
			path = strings.Replace(path, ".", "/", -1)

			// replace last '/' with a '.' because that marks the extension of the file
			lastIndex := strings.LastIndex(path, "/")
			firstPart := path[:lastIndex]
			secondPart := path[lastIndex:]
			secondPart = strings.Replace(secondPart, "/", ".", 1)
			path = firstPart + secondPart

			// remove _rootfolder_ from beginning of path
			path = strings.Replace(path, "_rootfolder_", "", 1)

			searchResults.Documents[i].Path = path

			// write search result template
			o := output.NewTemplate(output.SearchType, searchResults.Documents[i])
			err = o.Write()
			if err != nil {
				log.Error(err.Error())
				tb.Close()
				return
			}

			// wait for character input
			for {
			restart:
				// read single key press
				event := tb.PollEvent()
				char := event.Ch

				switch char {
				case 'p':
					goto previousResult
				case 'n':
					goto nextResult
				case 'd':
					goto download
				case 'q':
					goto quit
				default:
					goto restart
				}

			previousResult:
				if i != 0 {
					i = i - 1
					tb.Close()
					goto searchLoop
				}

				goto restart

			nextResult:
				if i == len(searchResults.Documents)-1 {
					goto quit
				}

				i = i + 1
				tb.Close()
				goto searchLoop

			quit:
				tb.Close()
				return

			download:
				// start a spinner on a new line while we download the file
				fmt.Println("")
				s := spinner.New(spinner.CharSets[24], 100*time.Millisecond)
				s.Prefix = "Downloading ... "
				s.Start()

				err = e.DownloadFile(searchResults.Documents[i].DownloadURL)
				if err != nil {
					log.Error(err.Error())
					tb.Close()
					return
				}
				s.Stop()

				// wait for keypress and continue loop
				fmt.Printf("\033[1;32mFile downloaded successfully.\033[0m")
				for {
					event := tb.PollEvent()
					char := event.Ch

					switch char {
					case 'n':
						goto nextResult
					case 'd':
						break // we already downloaed a file on this iteration don't do anything
					case 'q':
						goto quit
					default:
						goto restart
					}

				}
			}
		}
	default:
		help.ShowHelp()
		return
	}
}
