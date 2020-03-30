package healthcheck

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

// availableModelState is the state returned for a model when it is ready to be used
const availableModelState = "AVAILABLE"

type modelStatusBody struct {
	Code    string `json:"error_code"`
	Message string `json:"error_message"`
}

type modelBody struct {
	Version string          `json:"version"`
	State   string          `json:"state"`
	Status  modelStatusBody `json:"status"`
}

// ModelReady checks that the ml model API and our AI model is available
func ModelReady(modelURL string) (bool, error) {
	u, err := url.Parse(modelURL)
	if err != nil {
		return false, err
	}

	res, err := http.Get(u.String())
	if err != nil {
		return false, err
	}

	if res.StatusCode != http.StatusOK {
		return false, errors.New("ml model api returned invalid status code")
	}

	var body struct {
		Results []modelBody `json:"model_version_status"`
	}

	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		err = errors.Wrap(err, "decoding ml model api root model response")
		return false, err
	}

	if len(body.Results) != 0 {
		result := body.Results[0]
		if result.State == availableModelState {
			return true, nil
		}

		return false, errors.New("ml model not in available state")
	}

	return false, errors.New("ml model response returned no model results")
}

// TikaServiceReady checks that the ready endpoint of tikad is responding correctly
func TikaServiceReady(tikaURL string) (bool, error) {
	endpoint := fmt.Sprintf("%v/ready", tikaURL)
	u, err := url.Parse(endpoint)
	if err != nil {
		return false, err
	}

	res, err := http.Get(u.String())
	if err != nil {
		return false, err
	}

	if res.StatusCode != http.StatusOK {
		return false, errors.New("tika service returned invalid status code")
	}

	return true, nil
}
