package command

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/dannav/starship/cli/internal/engine"
	"github.com/dannav/starship/cli/internal/output"
	"github.com/dannav/starship/cli/internal/uerrors"
	tb "github.com/nsf/termbox-go"
)

// Search performs a search against all documents in the Starship API
func Search(args []string, e *engine.Engine) error {
	// ensure search text was provided
	if len(args) == 0 {
		return errors.New(uerrors.SearchNoTextGiven)
	}

	searchText := strings.Join(args, " ")

	// start a spinner while we perform a search
	s := spinner.New(spinner.CharSets[24], 100*time.Millisecond)
	s.Prefix = "Searching ... "
	s.Start()

	searchResults, err := e.Search(searchText)
	if err != nil {
		return errors.New(uerrors.SearchRequestFailed)
	}
	s.Stop()

	if len(searchResults.Documents) == 0 {
		return errors.New(uerrors.SearchNoResults)
	}

	for i := range searchResults.Documents {
	searchLoop:
		// create termbox session so we can read single key
		err = tb.Init()
		if err != nil {
			return err
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
			tb.Close()
			return err
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
			return nil

		download:
			// start a spinner on a new line while we download the file
			fmt.Println("")
			s := spinner.New(spinner.CharSets[24], 100*time.Millisecond)
			s.Prefix = "Downloading ... "
			s.Start()

			err = e.DownloadFile(searchResults.Documents[i].DownloadURL)
			if err != nil {
				tb.Close()
				return err
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

	return nil
}
