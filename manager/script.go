package manager

import (
	"bytes"
	"encoding/json"
	"net/http"

	"io/ioutil"

	"io"

	"strings"

	"github.com/pkg/errors"
)

// Script represents a nexus script
type Script struct {
	url      string
	username string
	password string
	client   *http.Client
}

// ExecuteWithFilePayload execute the script with the payload from given file
func (script *Script) ExecuteWithFilePayload(filename string) (string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", errors.Wrapf(err, "failed to read file %s", filename)
	}

	return script.Execute(bytes.NewBuffer(data))
}

// ExecuteWithJSONPayload execute the script with the given payload transformed to json
func (script *Script) ExecuteWithJSONPayload(payload interface{}) (string, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return "", errors.Wrap(err, "failed to marshal json payload")
	}

	return script.Execute(bytes.NewBuffer(data))
}

// ExecuteWithStringPayload
func (script *Script) ExecuteWithStringPayload(payload string) (string, error) {
	return script.Execute(strings.NewReader(payload))
}

// ExecuteWithoutPayload executes the script without any payload
func (script *Script) ExecuteWithoutPayload() (string, error) {
	return script.Execute(nil)
}

// Execute executes the script with the given payload
func (script *Script) Execute(reader io.Reader) (string, error) {
	req, err := http.NewRequest("POST", script.url, reader)
	if err != nil {
		return "", errors.Wrapf(err, "failed to create request")
	}
	req.Header.Set("Content-Type", "text/plain")
	req.SetBasicAuth(script.username, script.password)

	resp, err := script.client.Do(req)
	if err != nil {
		return "", errors.Wrapf(err, "failed to create script")
	}

	if resp.StatusCode != 200 {
		return "", errors.Errorf("creation of script failed, server returned %v", resp.StatusCode)
	}

	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "failed to read response body")
	}

	result := make(map[string]string)
	err = json.Unmarshal(data, &result)
	if err != nil {
		return "", errors.Wrap(err, "failed to unmarshal response body")
	}

	return result["result"], nil
}
