package mediamachine

import (
	"bytes"
	"encoding/json"
)

// TranscodeEncoder is the type representing the type of encoder that can be used for
// a transcode job.
type TranscodeEncoder = string

//TranscodeBitrate is the type representing the bitrate to be used for a transcode job.
type TranscodeBitrate = string

// TranscodeContainer is the type representing the container of the output video.
type TranscodeContainer = string

// TranscodeVideoSize is the type representing the output video size.
type TranscodeVideoSize = string

const (
	// EncoderH264 is the configuration for a `h264` encoder.
	EncoderH264 TranscodeEncoder = "h264"
	// EncoderH265 is the configuration for a `h265` encoder.
	EncoderH265 TranscodeEncoder = "h265"
	// EncoderVp8 is the configuration for a `vp8` encoder.
	EncoderVp8 TranscodeEncoder = "vp8"
	// EncoderVp9 is the configuration for a `vp9` encoder.
	EncoderVp9 TranscodeEncoder = "vp9"

	// Bitrate4Mbps is the configuration for a `4000 kbps` bitrate.
	Bitrate4Mbps TranscodeBitrate = "4000"
	// Bitrate2Mbps is the configuration for a `2000 kbps` bitrate.
	Bitrate2Mbps TranscodeBitrate = "2000"
	// Bitrate1Mbps is the configuration for a `1000 kbps` bitrate.
	Bitrate1Mbps TranscodeBitrate = "1000"

	// ContainerMP4 is the configuration for a `mp4` video container.
	ContainerMP4 TranscodeContainer = "mp4"
	// ContainerWebm is the configuration for a `webm` video container.
	ContainerWebm TranscodeContainer = "webm"
)

/*
TranscodeConfig configures the request for a video transcode operation.
The input video location can be specified via the FromUrl or the From method.

By default, the output has the same dimensions as the input video.
Set Width to desired value to customize.
Height can also be specified, however it is automatically calculated according to input aspect ratio if not specified.
*/
type TranscodeConfig struct {
	Container   TranscodeContainer // required
	Encoder     TranscodeEncoder   // required
	BitrateKBPS TranscodeBitrate   // required

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

	Height uint // Optional - by default, the output has same height as input video
	Width  uint // Optional - by default, the output has same width as input video

	SuccessURL string // Optional - Expect a POST call when job is successfully finished
	FailureURL string // Optional - Expect a POST call with failure details
}

/*
Transcode enqueues a request to the MediaMachine backend to asynchronously transcode the input video.

The output video is uploaded to the location specified in the TranscodeConfig.
Errors if the input configuration is invalid.
*/
func (m MediaMachine) Transcode(cfg TranscodeConfig) (Job, error) {
	if err := validateInputOutput(cfg.InputURL, cfg.OutputURL, cfg.InputCreds, cfg.OutputCreds); err != nil {
		return Job{}, err
	}

	tr := struct {
		APIKey string
		TranscodeConfig
	}{
		APIKey:          m.APIKey,
		TranscodeConfig: cfg,
	}

	body, err := json.Marshal(tr)
	if err != nil {
		return Job{}, err
	}
	return m.submit("/transcode", bytes.NewBuffer(body))
}
