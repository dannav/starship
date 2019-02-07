package healthcheck

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/pkg/errors"
)

// availableTFServingState is the state returned for a model when it is ready to be used
const availableTFServingState = "AVAILABLE"

type servingStatusBody struct {
	Code    string `json:"error_code"`
	Message string `json:"error_message"`
}

type servingBody struct {
	Version string            `json:"version"`
	State   string            `json:"state"`
	Status  servingStatusBody `json:"status"`
}

// ServingReady checks that tensorflow serving and our AI model is available
func ServingReady(servingURL string) (bool, error) {
	modelPathIndex := strings.LastIndex(servingURL, "/") + 1
	modelEndpoint := servingURL[:modelPathIndex] + "universal_encoder"

	u, err := url.Parse(modelEndpoint)
	if err != nil {
		return false, err
	}

	res, err := http.Get(u.String())
	if err != nil {
		return false, err
	}

	if res.StatusCode != http.StatusOK {
		return false, errors.New("tfserving returned invalid status code")
	}

	var body struct {
		Results []servingBody `json:"model_version_status"`
	}

	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		err = errors.Wrap(err, "decoding tfserving model response")
		return false, err
	}

	if len(body.Results) != 0 {
		result := body.Results[0]
		if result.State == availableTFServingState {
			return true, nil
		}

		return false, errors.New("tfserving not in available state")
	}

	return false, errors.New("tfserving response returned no model results")
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
