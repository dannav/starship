// fileop manages manipulating files on the local machine

package fileop

import (
	"bytes"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
)

// WriteFile handles writing a multipart file to a location on disk
func WriteFile(file multipart.File, location string) error {
	// create parent file paths first so we don't get dne errors
	dir := filepath.Dir(location)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return err
	}

	err := ioutil.WriteFile(location, buf.Bytes(), 0644)
	if err != nil {
		return err
	}

	return nil
}

// DeleteFile handles deleting a file at the given location
func DeleteFile(location string) error {
	if err := os.Remove(location); err != nil {
		return err
	}

	return nil
}
