package mediamachine

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

// JobRequest represents the new mediamachine job request.
// Use the convenience methods to invoke api calls.
type JobRequest struct {
	APIKey string `json:"apiKey"`

	// Web-hook urls
	SuccessURL string `json:"successUrl,omitempty"`
	FailureURL string `json:"failureUrl,omitempty"`

	// One of these two should be set
	OutputURL  string     `json:"outputUrl,omitempty"`
	OutputBlob *Blob `json:"outputBlob,omitempty"`

	// One of these two should be set
	InputURL  string     `json:"inputUrl,omitempty"`
	InputBlob *Blob `json:"inputBlob,omitempty"`
}

// WatermarkImage represents a Watermark of type image that users can use to put on their
// output images or videos.
type WatermarkImage struct {
	Path string `json:"path,omitempty"`
	ImageName string `json:"imageName,omitempty"`
	Width uint8 `json:"width,omitempty"`
	Height uint8 `json:"height,omitempty"`
}

// Watermark represent a watermark object to be used in the image/gif/video.
// This watermark can be either a text or an image.
type Watermark struct {
	WatermarkText string `json:"text,omitempty"`
	WatermarkImage *WatermarkImage `json:"image,omitempty"`
	WatermarkFontSize uint8 `json:"fontSize,omitempty"`
	WatermarkColor string `json:"color,omitempty"`
	WatermarkOpacity float64 `json:"opacity,omitempty"`
	WatermarkPosition string `json:"position,omitempty"`// oneOf: topLeft, topRight, bottomLeft, bottomRight
}

// Blob represents a bucket store like S3/Google Bucket/Azure Bucket
// where objects are presented by a combination of bucket + (prefixed) keys
type Blob struct {
	// One of: s3,gcp,azure
	Store  string `json:"store"`
	BucketName string `json:"bucket"`
	KeyPath   string `json:"key"`

	// Any one of these three
	AwsCreds   *AWSCreds       `json:"awsCreds"`
	AzureCreds *AzureCreds     `json:"azureCreds"`
	GCPCreds   json.RawMessage `json:"gcpCreds"`
}

func newBlobWithDefaults() *Blob{
	return &Blob{}
}

// NewS3BlobWithDefaults returns a new Blob instance configured to be used with S3.
// after calling this method you need to provide the credentials.
func NewS3BlobWithDefaults() *Blob {
	blob := newBlobWithDefaults()
	blob.Store = "s3"
	return blob
}

// NewAzureBlobWithDefaults returns a new Blob instance configured to be used with Azure.
// After calling this method you need to provide the credentials.
func NewAzureBlobWithDefaults() *Blob {
	blob := newBlobWithDefaults()
	blob.Store = "azure"
	return blob
}

// NewGCPBlobWithDefaults returns a new Blob instance configured to be used with GCP.
// After calling this method you need to provide the credentials.
func NewGCPBlobWithDefaults() *Blob {
	blob := newBlobWithDefaults()
	blob.Store = "gcp"
	return blob
}

// Bucket returns the Blob instance configured with the bucket to be used on this blob.
func (b *Blob) Bucket(bucketName string) *Blob {
	b.BucketName = bucketName
	return b
}

// Key returns the Blob instance configured with the key to be used on this blob.
func (b *Blob) Key(key string) *Blob {
	b.KeyPath = key
	return b
}

// AWSCredentials returns the Blob instance configured with the AWS Credentials. You need to provide:
//  - accessKeyId : the accessKeyId provided by Amazon
//  - secretAccessKey : the secretAccessKey provided by Amazon
//  - region: the region in which your bucket is located
func (b *Blob) AWSCredentials(accessKeyId string, secretAccessKey string, region string) *Blob {
	creds := &AWSCreds{
		AccessKeyID:     accessKeyId,
		SecretAccessKey: secretAccessKey,
		Region:          region,
	}
	b.AwsCreds = creds
	return b
}


// AzureCredentials returns the Blob instance configured with the Azure credentials. You need to provide:
//  - accountName : accountName provided by Azure
//  - accountKey : accountKey provided by Azure
func (b *Blob) AzureCredentials(accountName string, accountKey string) *Blob {
	creds := &AzureCreds{
		AccountName: accountName,
		AccountKey:  accountKey,
	}
	b.AzureCreds = creds
	return b
}

// GCPCredentials returns the Blob instance configured with the Google Private Cloud credentials. You need to provide:
//  - json : a string representation of the json provided by Google to access their services.
func (b * Blob) GCPCredentials(json json.RawMessage) *Blob {
	b.GCPCreds = json
	return b
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
type JobOpt = func(tr JobRequest)


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
