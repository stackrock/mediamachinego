package mediamachine

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
)

// SummaryType represent the possible output type of the summary.
type SummaryType = string

const (
	// SummaryTypeGif - represents an output of type `gif`
	SummaryTypeGif SummaryType = "gif"
	// SummaryTypeMp4 - represents an output of type `mp4`
	SummaryTypeMp4 SummaryType = "mp4"
)

/*
SummaryConfig configures the request for a video summary.
The input video location can be specified via the FromUrl or the From method.

By default, the output has the same dimensions as the input video, set Width to desired value to customize.
Height is automatically calculated according to input aspect ratio.
*/
type SummaryConfig struct {
	RemoveAudio bool // Only applicable when Type is set to SummaryTypeMp4, ignored otherwise

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

	Width     uint      // Optional - by default, the output has same width as input video
	Watermark Watermark // Optional

	SuccessURL string // Optional - Expect a POST call when job is successfully finished
	FailureURL string // Optional - Expect a POST call with failure details
}

/*
SummaryGIF enqueues a request to the MediaMachine backend to asynchronously generate a summary
uploaded to an s3-compatible-store for the input video.

The output is a GIF and is uploaded to the location specified in the SummaryConfig.
Errors if the input configuration is invalid.
*/
func (m MediaMachine) SummaryGIF(cfg SummaryConfig) (Job, error) {
	return m.summary(SummaryTypeGif, cfg)
}

/*
SummaryMP4 enqueues a request to the MediaMachine backend to asynchronously generate a summary
uploaded to an s3-compatible-store for the input video.

The output summary video is packaged as an mp4 and is uploaded to the location specified in the SummaryConfig.
Errors if the input configuration is invalid.
*/
func (m MediaMachine) SummaryMP4(cfg SummaryConfig) (Job, error) {
	return m.summary(SummaryTypeMp4, cfg)
}

func (m MediaMachine) summary(summaryType SummaryType, cfg SummaryConfig) (Job, error) {
	if err := validateInputOutput(cfg.InputURL, cfg.OutputURL, cfg.InputCreds, cfg.OutputCreds); err != nil {
		return Job{}, err
	}
	sr := struct {
		APIKey string
		SummaryConfig
	}{
		APIKey:        m.APIKey,
		SummaryConfig: cfg,
	}

	body, err := json.Marshal(sr)
	if err != nil {
		return Job{}, err
	}
	return m.submit("/summary/"+summaryType, bytes.NewBuffer(body))
}

func validateInputOutput(inputURL, outputURL string, inputCreds, outputCreds Creds) error {
	uri, err := url.ParseRequestURI(inputURL)
	if err != nil {
		return err
	}
	switch uri.Scheme {
	case "s3", "azure", "gcp":
		if inputCreds == nil {
			return fmt.Errorf("inputCreds are needed when store is '%s'", uri.Scheme)
		}
	case "http", "https":
		// no-op, pass it through as it is
	default:
		return fmt.Errorf("inputURL has unsupported scheme: '%s'", uri.Scheme)
	}

	uri, err = url.ParseRequestURI(outputURL)
	if err != nil {
		return err
	}
	switch uri.Scheme {
	case "s3", "azure", "gcp":
		if outputCreds == nil {
			return fmt.Errorf("outputCreds are needed when store is '%s'", uri)
		}
	case "http", "https":
		// no-op, pass it through as it is
	default:
		return fmt.Errorf("outputURL has unsupported scheme: '%s'", uri)
	}

	return nil
}
