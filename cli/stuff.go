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
	pb "gopkg.in/cheggaaa/pb.v1"
)

// Add timeout of 30secs on web requests
var client = &http.Client{
	Timeout: time.Second * 30,
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

	// if help flags passed show help and exit
	if showHelp {
		help.ShowHelp()
		return
	}

	// parse command and arguments
	argLen := len(os.Args)
	if argLen > 1 {
		cmd = strings.ToLower(os.Args[1])

		if argLen >= 2 {
			args = os.Args[2:]
		}
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

		// format index path
		if indexPath[0] != '/' {
			indexPath = "/"
		}

		// index file passed in argument
		filePath := args[0]

		// check if path provided exists
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			log.Error(uerrors.IndexFileDoesNotExist)
			return
		}

		// start reading the file
		f, err := os.Open(filePath)
		if err != nil {
			log.Error(err.Error())
			return
		}
		defer f.Close()

		fi, err := f.Stat()
		if err != nil {
			log.Error(err.Error())
			return
		}

		e := engine.NewEngine(client)

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
			fmt.Printf("\n Overwrite file that exists at index path: %v (y/n) ", path)

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

		fmt.Printf("\n Uploading %v\n", filePath)

		// start progress bar for upload - proxy is used to increment progress bar
		fileSize := int(fi.Size())
		bar := pb.New(fileSize).SetUnits(pb.U_BYTES)
		proxyReader := bar.NewProxyReader(f)

		// render progress bar
		bar.SetMaxWidth(80)
		bar.Start()

		// index the file in search engine
		err = e.Index(bar, proxyReader, filepath.Base(filePath), indexPath)
		if err != nil {
			log.Error(err.Error())
			return
		}

		fmt.Println("")

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
		s.Prefix = "Searching... "
		s.Start()

		e := engine.NewEngine(client)
		searchResults, err := e.Search(searchText)
		if err != nil {
			log.Error(uerrors.SearchRequestFailed)
			return
		}
		s.Stop()

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
				fmt.Print("\n")
				s := spinner.New(spinner.CharSets[24], 100*time.Millisecond)
				s.Prefix = "Downloading... "
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
