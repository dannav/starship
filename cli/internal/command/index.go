package command

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/briandowns/spinner"
	"github.com/dannav/starship/cli/internal/engine"
	"github.com/dannav/starship/cli/internal/uerrors"
)

// Index is the document indexing command
func Index(args []string, indexPath string, e engine.Engine) error {
	// format index path to always have leading '/'
	if indexPath[0] != '/' {
		indexPath = "/" + indexPath
	}

	// ensure a path was provided, the first arg is the path
	if len(args) == 0 {
		return errors.New(uerrors.IndexPathNotProvided)
	}

	// get filepath to index from args
	filePath := args[0]

	// check if path provided exists on the system
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return errors.New(uerrors.IndexFileDoesNotExist)
	}

	// create valid index path [ /foo/bar/readme.md ] that contains filename
	var path string
	if indexPath[len(indexPath)-1] != '/' {
		path = fmt.Sprintf("%v/%v", indexPath, filepath.Base(filePath))
	} else {
		path = fmt.Sprintf("%v%v", indexPath, filepath.Base(filePath))
	}

	// check if a file already exists at this index path on starship
	exists, err := e.ExistsAtIndexPath(path)
	if err != nil {
		return err
	}

	// if file with same name exists at index path ask if user wants to overwrite it, cancel index if no
	if exists {
		fmt.Printf(" Overwrite file that exists at index path: %v (y/n) ", path)

		var input string
		fmt.Scanln(&input)

		switch input {
		case "y", "yes", "Y", "YES":
			break
		case "n", "no", "NO", "No":
			return nil
		default:
			return nil
		}
	}

	// start a spinner while we index the file
	s := spinner.New(spinner.CharSets[24], 100*time.Millisecond)
	s.Prefix = " Indexing " + filePath + " ... "
	s.Start()

	err = e.Index(filePath, filepath.Base(filePath), indexPath)
	if err != nil {
		return err
	}
	s.Stop()

	return nil
}
