package embedding

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

const (
	// vectorDimensions represents the size of the vector returned from tensorflow for a word embedding
	vectorDimensions int = 512
)

// ServingResponse is the prediction response from the tensorflow serving hosted universal sentence encoder
type ServingResponse struct {
	Outputs [][]float32 `json:"outputs"`
}

// Generate creates word embeddings through universal sentence encoder for a group of text
func Generate(inputs []string, url string, client *http.Client) ([][]float32, error) {
	var resp ServingResponse

	servingReq := struct {
		Inputs []string `json:"inputs"`
	}{
		Inputs: inputs,
	}

	// create request to serving universal sentence encoder
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(servingReq); err != nil {
		return nil, errors.Wrap(err, "could not marshal serving request body")
	}

	endpoint := url
	req, err := http.NewRequest(http.MethodPost, endpoint, &body)
	if err != nil {
		return nil, errors.Wrap(err, "create serving request")
	}

	// make word embedding request
	res, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "perform serving request")
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("serving request failed")
	}

	// get response from serving
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return nil, errors.Wrap(err, "decode serving response")
	}

	return resp.Outputs, nil
}
