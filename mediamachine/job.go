package mediamachine

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

const (
	JobStatusQueued  = "queued"
	JobStatusErrored = "errored"
	JobStatusDone    = "done"
)

/* Job is a handle for a job submitted to the MediaMachine API

The Job ID can be stored and re-used by the caller to fetch status later.
Jobs are asynchronously processed by the StackRock API.
*/
type Job struct {
	ID        string // Unique Job ID
	CreatedAt time.Time

	lastStatusFetch time.Time
	minFresh        time.Duration
	status          string
}

// FetchStatus queries the MediaMachine API backend for the latest status for this job
func (j Job) FetchStatus() (string, error) {
	if j.ID == "" {
		return "", fmt.Errorf("cannot fetch job status: ID is not set")
	}

	if time.Now().Sub(j.lastStatusFetch) < j.minFresh {
		// don't re-check so soon, return cached status
		return j.status, nil
	}

	resp, err := httpClient.Get(fmt.Sprintf("%s/job/status?reqId=%s", apiEndpoint, j.ID))
	if err != nil {
		return "", err
	}
	j.lastStatusFetch = time.Now()

	// cache status response for new min-fresh duration
	cc := resp.Header.Get("X-Cache-Min-Fresh-Sec")
	if cc != "" {
		sec, parseErr := strconv.ParseInt(cc, 10, 32)
		if parseErr != nil {
			// ignore, assume 60 sec by default
			j.minFresh = time.Minute
		} else {
			j.minFresh = time.Second * time.Duration(sec)
		}
	}

	payload := map[string]string{}
	err = json.NewDecoder(resp.Body).Decode(&payload)
	if err != nil {
		return "", err
	}
	if payload["error"] != "" {
		// set the new job status on the status obj as well
		j.status = JobStatusErrored
		return j.status, fmt.Errorf("job %s errored", j.ID)
	}
	// set the new job status on the status obj as well
	j.status = payload["status"]
	return j.status, nil
}
