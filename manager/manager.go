package manager

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

// New creates a new nexus script manager
func New(url string, username string, password string) *Manager {
	client := &http.Client{
		Timeout: time.Second * 10,
	}

	scriptBaseUrl := url + "/service/rest/v1/script"

	return &Manager{
		scriptBaseUrl: scriptBaseUrl,
		username:      username,
		password:      password,
		client:        client,
	}
}

// Manager manages nexus scripting
type Manager struct {
	scriptBaseUrl string
	username      string
	password      string

	client *http.Client
}

func (manager *Manager) WithTimeout(timeoutInSeconds int) {
	manager.client.Timeout = time.Second * time.Duration(timeoutInSeconds)
}

// Get returns a script instance of the given name
func (manager *Manager) Get(name string) (*Script, error) {
	exists, err := manager.Exists(name)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to check if %s exists", name)
	}

	if !exists {
		return nil, errors.Wrapf(err, "script %s does not exists", name)
	}

	return manager.getScript(name), nil
}

func (manager *Manager) getScript(name string) *Script {
	return &Script{
		url:      manager.scriptBaseUrl + "/" + name + "/run",
		username: manager.username,
		password: manager.password,
		client:   manager.client,
	}
}

// CreateFromFile creates a new script at nexus with the content from the given file
func (manager *Manager) CreateFromFile(name string, filename string) (*Script, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read file %s", filename)
	}

	return manager.Create(name, string(data))
}

// Create creates a new script at nexus
func (manager *Manager) Create(name string, script string) (*Script, error) {
	exists, err := manager.Exists(name)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to check if script %s exists", name)
	}

	method := "POST"
	url := manager.scriptBaseUrl
	if exists {
		method = "PUT"
		url = manager.scriptBaseUrl + "/" + name
	}

	payload, err := manager.createPayload(name, script)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create payload")
	}

	err = manager.executeUpload(method, url, payload)
	if err != nil {
		return nil, err
	}

	return manager.getScript(name), nil
}

// Exists returns true if a script exists with the given name
func (manager *Manager) Exists(name string) (bool, error) {
	request, err := http.NewRequest("GET", manager.scriptBaseUrl+"/"+name, nil)
	if err != nil {
		return false, errors.Wrapf(err, "failed to create request")
	}
	request.SetBasicAuth(manager.username, manager.password)

	response, err := manager.client.Do(request)
	if err != nil {
		return false, errors.Wrapf(err, "failed to check if script exists")
	}

	if response.StatusCode == http.StatusOK {
		return true, nil
	} else if response.StatusCode == http.StatusNotFound {
		return false, nil
	}

	return false, errors.Errorf("nexus respond with %v", response.StatusCode)
}

func (manager *Manager) createPayload(name string, script string) (io.Reader, error) {
	payload := make(map[string]string)

	payload["name"] = name
	payload["type"] = "groovy"
	payload["content"] = script

	data, err := json.Marshal(payload)
	if err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal payload")
	}

	return bytes.NewBuffer(data), nil
}

func (manager *Manager) executeUpload(method string, url string, payload io.Reader) error {
	request, err := http.NewRequest(method, url, payload)
	if err != nil {
		return errors.Wrapf(err, "failed to create request")
	}
	request.Header.Set("Content-Type", "application/json")
	request.SetBasicAuth(manager.username, manager.password)

	response, err := manager.client.Do(request)
	if err != nil {
		return errors.Wrapf(err, "failed to create script")
	}

	if response.StatusCode != 204 {
		return errors.Errorf("creation of user replication script failed, server returned %v", response.StatusCode)
	}

	return nil
}
