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

var client = &http.Client{
	Timeout: time.Second * 30,
}

func main() {
	var cmd string
	var args []string
	var showHelp bool

	flag.BoolVar(&showHelp, "h", false, "show help text")
	flag.BoolVar(&showHelp, "help", false, "show help text")
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
		help.ShowHelp()
		return
	case "index":
		// ensure a path was provided
		if len(args) == 0 {
			log.Error(uerrors.IndexPathNotProvided)
			return
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

		// start progress bar for upload - proxy is used to increment progress bar
		fileSize := int(fi.Size())
		bar := pb.New(fileSize).SetUnits(pb.U_BYTES)
		proxyReader := bar.NewProxyReader(f)

		// index the file in search engine
		e := engine.NewEngine(client)
		err = e.Index(proxyReader, filepath.Base(filePath))
		if err != nil {
			log.Error(err.Error())
			return
		}

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

			o := output.NewTemplate(output.SearchType, searchResults.Documents[i])
			err = o.Write()
			if err != nil {
				log.Error(err.Error())
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
				fmt.Printf("\n")
				s := spinner.New(spinner.CharSets[24], 100*time.Millisecond)
				s.Prefix = "Downloading... "
				s.Start()

				err = e.DownloadFile(searchResults.Documents[i].DownloadURL)
				if err != nil {
					log.Error(err.Error())
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
