package embedding

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// ModelResponse is the prediction response from the ml model api
type ModelResponse struct {
	Outputs [][]float32 `json:"outputs"`
}

// Generate creates word embeddings through universal sentence encoder for a group of text
func Generate(inputs []string, url string, client *http.Client) ([][]float32, error) {
	var resp ModelResponse

	modelReq := struct {
		Inputs []string `json:"inputs"`
	}{
		Inputs: inputs,
	}

	// create request to ml model
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(modelReq); err != nil {
		return nil, errors.Wrap(err, "could not marshal model request body")
	}

	endpoint := fmt.Sprintf("%v:predict", url)
	req, err := http.NewRequest(http.MethodPost, endpoint, &body)
	if err != nil {
		return nil, errors.Wrap(err, "create model request")
	}

	// make word embedding request
	res, err := client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "perform model request")
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("model request failed")
	}

	// get response from ml model api
	if err := json.NewDecoder(res.Body).Decode(&resp); err != nil {
		return nil, errors.Wrap(err, "decode model response")
	}

	return resp.Outputs, nil
}
