package mediamachine

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const (
	version     = "1.0.0"
	ua          = "Stackrock/MediaMachine/Go/" + version
	apiEndpoint = "https://api.stackrock.io"
)

// MediaMachine gives you access to the various operations you can perform using the StackRock API.
type MediaMachine struct {
	APIKey string // Your stackrock API key goes here
}

var httpClient = http.Client{Timeout: time.Second * 10}

func (m MediaMachine) submit(path string, body io.Reader) (Job, error) {
	j := Job{}
	req, err := http.NewRequest("POST", apiEndpoint+path, body)
	if err != nil {
		return j, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", ua)

	resp, err := httpClient.Do(req)
	if err != nil {
		return j, err
	}

	// read body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return j, err
	}

	// parse body
	payload := make(map[string]interface{})
	if err = json.Unmarshal(respBody, &payload); err != nil {
		log.Printf("unexpected server response: %s", respBody)
		return j, err
	}

	// errored
	if payload["error"] != nil {
		return j, fmt.Errorf("failed to submit job to MediaMachine API. Error: %s", payload["error"])
	}

	if resp.StatusCode == http.StatusOK {
		err = json.Unmarshal(respBody, &j)
		return j, err
	}

	return j, fmt.Errorf("unexpected server response: %s", payload)
}
