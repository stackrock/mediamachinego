package mediamachine

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

// jobRequest represents the new mediamachine job request.
// Use the convenience methods to invoke api calls. Example:
// job, err := mediamachine.Thumbnail(
//	mediamachine.APIKey("SECRET_API_KEY"),
//	mediamachine.Webhooks("http://mytestserver/success", "http://mytestserver/failure"),
//	mediamachine.AzureInput("testbucket", "/foo/bar/testvideo.mp4", azureCreds),
//	mediamachine.AzureOutput("testbucket", "/foo/bar/testvideo.jpg", azureCreds),
//)
type jobRequest struct {
	APIKey string `json:"apiKey"`

	// Web-hook urls
	SuccessURL string `json:"successUrl,omitempty"`
	FailureURL string `json:"failureUrl,omitempty"`

	// One of these two should be set
	OutputURL  string     `json:"outputUrl,omitempty"`
	OutputBlob *BlobStore `json:"outputBlob,omitempty"`

	// One of these two should be set
	InputURL  string     `json:"inputUrl,omitempty"`
	InputBlob *BlobStore `json:"inputBlob,omitempty"`
}

// BlobStore represents a bucket store like S3/Google Bucket/Azure Bucket
// where objects are presented by a combination of bucket + (prefixed) keys
type BlobStore struct {
	// One of: s3,gcp,azure
	Store  string `json:"store"`
	Bucket string `json:"bucket"`
	Key    string `json:"key"`

	// Any one of these three
	AwsCreds   *AWSCreds       `json:"awsCreds"`
	AzureCreds *AzureCreds     `json:"azureCreds"`
	GCPCreds   json.RawMessage `json:"gcpCreds"`
}

// AWSCreds - re-usable aws credentials that can be attached to a BlobStore
// Allows callers to mix-and-match blob stores. Callers are encouraged to
// provide the smallest surface-area credentials.
type AWSCreds struct {
	AccessKeyID     string `json:"accessKeyId"`
	SecretAccessKey string `json:"secretAccessKey"`
	Region          string `json:"region,omitempty"`
}

// AzureCreds - re-usable azure credentials that can be attached to a BlobStore
// Allows callers to mix-and-match blob stores. Callers are encouraged to
// provide the smallest surface-area credentials.
type AzureCreds struct {
	AccountName string `json:"accountName"`
	AccountKey  string `json:"accountKey"`
}

// JobOpt is a convenience helper for fluently building job requests
type JobOpt = func(tr jobRequest)

func APIKey(apiKey string) JobOpt {
	return func(jr jobRequest) { jr.APIKey = apiKey }
}

func Webhooks(success, failure string) JobOpt {
	return func(jr jobRequest) { jr.SuccessURL = success; jr.FailureURL = failure }
}
func OutputURL(url string) JobOpt {
	return func(jr jobRequest) { jr.OutputURL = url }
}
func InputURL(url string) JobOpt {
	return func(tr jobRequest) { tr.InputURL = url }
}

func S3Input(bucket, key string, creds AWSCreds) JobOpt {
	return inputBlob(BlobStore{Store: "s3", Bucket: bucket, Key: key, AwsCreds: &creds})
}
func S3Output(bucket, key string, creds AWSCreds) JobOpt {
	return outputBlob(BlobStore{Store: "s3", Bucket: bucket, Key: key, AwsCreds: &creds})
}

func GCPInput(bucket, key string, creds json.RawMessage) JobOpt {
	return inputBlob(BlobStore{Store: "gcp", Bucket: bucket, Key: key, GCPCreds: creds})
}
func GCPOutput(bucket, key string, creds json.RawMessage) JobOpt {
	return outputBlob(BlobStore{Store: "gcp", Bucket: bucket, Key: key, GCPCreds: creds})
}

func AzureInput(bucket, key string, creds AzureCreds) JobOpt {
	return inputBlob(BlobStore{Store: "azure", Bucket: bucket, Key: key, AzureCreds: &creds})
}
func AzureOutput(bucket, key string, creds AzureCreds) JobOpt {
	return outputBlob(BlobStore{Store: "azure", Bucket: bucket, Key: key, AzureCreds: &creds})
}

func outputBlob(blob BlobStore) JobOpt {
	return func(jr jobRequest) { jr.OutputBlob = &blob }
}
func inputBlob(blob BlobStore) JobOpt {
	return func(jr jobRequest) { jr.InputBlob = &blob }
}

func sendRequest(service string, reqBody json.RawMessage) (Job, error) {
	url := Settings.URL + "/" + service
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return Job{}, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", Settings.userAgent)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := client.Do(req)
	if err != nil {
		return Job{}, err
	}

	decoder := json.NewDecoder(resp.Body)
	// Job created successfully
	if resp.StatusCode == http.StatusCreated {
		job := Job{}
		if err := decoder.Decode(&job); err != nil {
			return Job{}, err
		}
		return job, nil
	} else {
		createErr := Error{}
		if err := decoder.Decode(&createErr); err != nil {
			return Job{}, err
		}
		return Job{}, createErr
	}
}
