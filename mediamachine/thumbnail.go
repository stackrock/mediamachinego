package mediamachine

import (
	"bytes"
	"encoding/json"
)

/* ThumbnailConfig configures the request for intelligent thumbnail-s3-compatible-store generation for the input video

The input video location can be specified via the FromUrl or the From method.
By default, the thumbnail-s3-compatible-store has the same dimensions as the input video, set the Width field to desired value to customize.

If Width is set, Height is calculated automatically to maintain aspect ratio.
*/
type ThumbnailConfig struct {
	// Structured as {http|https|s3|azure|gcp}://{bucket-name}/{prefix-if-any}/{object-name}
	// Examples: s3://bucket/prefix/input.mp4, https://example.com/files/input.mp4
	InputURL  string
	OutputURL string

	// Provide credentials to S3/Azure/GCP for input/output locations
	// Can be nil if using http(s) input/output urls - make sure url endpoints are accessible
	// Note: You can use a different set of creds for input and output if you want to upload to a totally different
	// account for example or to a different bucket if you generate keys specific to bucket etc. or reuse the same
	// Creds object. See examples folder for usage.
	InputCreds  Creds
	OutputCreds Creds

	Width     uint      // Optional - by default, the thumbnail-s3-compatible-store has same width as input video
	Watermark Watermark // Optional

	SuccessURL string // Optional - Expect a POST call when job is successfully finished
	FailureURL string // Optional - Expect a POST call with failure details
}

/*
Thumbnail enqueues a request to the MediaMachine backend to asynchronously generate a thumbnail-s3-compatible-store for the input video.

The output image is uploaded to the location specified in the ThumbnailConfig.
Errors if the input configuration is invalid.
*/
func (m MediaMachine) Thumbnail(cfg ThumbnailConfig) (Job, error) {
	if err := validateInputOutput(cfg.InputURL, cfg.OutputURL, cfg.InputCreds, cfg.OutputCreds); err != nil {
		return Job{}, err
	}

	tr := struct {
		APIKey string
		ThumbnailConfig
	}{
		APIKey:          m.APIKey,
		ThumbnailConfig: cfg,
	}

	body, err := json.Marshal(tr)
	if err != nil {
		return Job{}, err
	}
	return m.submit("/thumbnail", bytes.NewBuffer(body))
}
