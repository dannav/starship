package main

import (
	"bytes"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// File represents a file in the ardan-files directory
type File struct {
	Repo    string `json:"repo"`
	Content string `json:"content"`
}

var client = &http.Client{
	Timeout: time.Second * 30,
}

func main() {
	var files []string

	root := "./ardan-files"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".md") || strings.Contains(path, ".pdf") {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}

		// remove filename from path
		path := strings.Replace(file, filepath.Base(file), "", 1)
		err = Index(f, filepath.Base(file), path)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Index indexes a readme file with the searchd api
func Index(file io.Reader, filename, path string) error {
	var buf bytes.Buffer
	encoder := multipart.NewWriter(&buf)
	field, err := encoder.CreateFormFile("content", filename)
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
