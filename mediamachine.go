package mediamachine

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

// SDKSettings - callers can override the URL if needed. No change needed by default.
type SDKSettings struct {
	URL       string
	userAgent string
}

// Error - StackRock library error usually occurs if StackRock server was unable to accept the request.
type Error struct {
	Code    string `json:"error_code"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

// Error returns a string representation of the error
func (e Error) Error() string {
	return fmt.Sprintf("%s (code: %s)", e.Message, e.Code)
}

var Settings = SDKSettings{
	URL:       getEnv("STACKROCK_URL", "https://api.stackrock.io"),
	userAgent: "Stackrock/1.0.0 [Go]",
}

// Job is the struct that represent a thumbnail/summary/transcode job running on our server.
// a Job can be on one of the three following states: `queued`, `done`, `errored`.
// If the Job is on `queued` state, it means that it hasn't been processed yet.
// If the Job is on `done` state, it means that the job is done and the output is located on the provided output provider.
// If the Job is on `errored` state, it means that the processing failed.
type Job struct {
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
}

// Status fetch the current status of the Job and returns a string and error.
// The string will contain some of the following values: `queued`, `done`, `errored`.
// error will only contains something if the communication with MediaMachine server fails.
func (job Job) Status() (string, error) {
	if job.Id == "" {
		return "", fmt.Errorf("cannot fetch job status: job Id is not set")
	}
	resp, err := http.Get(Settings.URL + "/job/status?reqId=" + job.Id)
	if err != nil {
		return "", err
	}

	payload := map[string]string{}
	err = json.NewDecoder(resp.Body).Decode(&payload)
	if err != nil {
		return "", err
	}
	return payload["status"], nil
}

const (
	SvcThumbnail  = "thumbnail"
	SvcSummaryGIF = "summary/gif"
	SvcSummaryMP4 = "summary/mp4"
)

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
