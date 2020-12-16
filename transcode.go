package mediamachine

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
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
	// ENCODER_H264 is the configuration for a `h264` encoder.
	ENCODER_H264 TranscodeEncoder = "h264"
	// ENCODER_H265 is the configuration for a `h265` encoder.
	ENCODER_H265 TranscodeEncoder = "h265"
	// ENCODER_VP8 is the configuration for a `vp8` encoder.
	ENCODER_VP8 TranscodeEncoder = "vp8"

	// BITRATE_EIGHT_MEGAKBPS is the configuration for a `8000 kbps` bitrate.
	BITRATE_EIGHT_MEGAKBPS TranscodeBitrate = "8000"
	// BITRATE_FOUR_MEGAKBPS is the configuration for a `4000 kbps` bitrate.
	BITRATE_FOUR_MEGAKBPS TranscodeBitrate = "4000"
	// BITRATE_ONE_MEGAKBPS is the configuration for a `1000 kbps` bitrate.
	BITRATE_ONE_MEGAKBPS TranscodeBitrate = "1000"

	// CONTAINER_MP4 is the configuration for a `mp4` video container.
	CONTAINER_MP4 TranscodeContainer = "mp4"
	// CONTAINER_WEBM is the configuration for a `webm` video container.
	CONTAINER_WEBM TranscodeContainer = "webm"

	// VIDEOSIZE_FULL_HD is the configuration for a `1080` video output.
	VIDEOSIZE_FULL_HD = "1080"
	// VIDEOSIZE_HD is the configuration for a `720` video output.
	VIDEOSIZE_HD = "720"
	// VIDEOSIZE_SD is the configuration for a `480` video output.
	VIDEOSIZE_SD = "480"
)

// TranscodeOpts contains the configuration data of the output video of a transcode process.
type TranscodeOpts struct {
	TranscodeEncoder string `json:"encoder,omitempty"`
	TranscodeBitrateKbps string `json:"bitrateKbps,omitempty"`
	TranscodeContainer string `json:"container,omitempty"`
	TranscodeVideoSize string `json:"videoSize,omitempty"`
}


// NewTranscodeOptsWithDefaults returns a new TranscodeOpts struct configured with the
// following default values:
//   - Encoder: H264
//   - Bitrate: 4000Kbps
//   - Container: MP4
//   - VideoSize: 720p
func NewTranscodeOptsWithDefaults() *TranscodeOpts {
	to := &TranscodeOpts{
		TranscodeEncoder:     ENCODER_H264,
		TranscodeBitrateKbps: BITRATE_FOUR_MEGAKBPS,
		TranscodeContainer:   CONTAINER_MP4,
		TranscodeVideoSize:   VIDEOSIZE_HD,
	}

	return to
}

// Encoder configures the encoder used on the output video.
// Valid options are:
//   - ENCODER_H264
//   - ENCODER_H265
//   - ENCODER_VP8
func (to *TranscodeOpts) Encoder(encoder TranscodeEncoder) *TranscodeOpts {
	to.TranscodeEncoder = encoder
	return to
}

// BitrateKbps configures the output bitrate.
// Valid options are:
//  - BITRATE_EIGHT_MEGAKBPS
//  - BITRATE_FOUR_MEGAKBPS
//  - BITRATE_ONE_MEGAKBPS
func (to *TranscodeOpts) BitrateKbps(bitrate TranscodeBitrate) *TranscodeOpts {
	to.TranscodeBitrateKbps = bitrate
	return to
}

// Container configures the container format of the output.
// Valid options are:
//  - CONTAINER_MP4
//  - CONTAINER_WEBM
func (to *TranscodeOpts) Container(container TranscodeContainer) *TranscodeOpts {
	to.TranscodeContainer = container
	return to
}

// VideoSize configures the output video resolution.
// Valid options are:
//  - VIDEOSIZE_FULL_HD
//  - VIDEOSIZE_HD
//  - VIDEOSIZE_SD
func (to *TranscodeOpts) VideoSize(videoSize TranscodeVideoSize) *TranscodeOpts {
	to.TranscodeVideoSize = videoSize
	return to
}

// TranscodeJob - transcode an input video from virtually any format to mp4/webm.
// The input video location can be specified via the FromUrl or the From method.
// Similarly, ToUrl or the To methods for transcode output video location.
// By default, the transcoded video is the same {Width x Height} as the input video - use the Width method to customize
type TranscodeJob struct {
	JobRequest
	WidthInt uint16 `json:"width,omitempty"`
	TranscodeWatermark *Watermark `json:"watermark,omitempty"`
	TranscodeOpts *TranscodeOpts `json:"transcode,omitemtpy"`
}

// NewTranscodeJobWithDefaults returns a configured TranscodeJob. The returned instance is configured
// with the following defaults:
//   - Encoder: H264
//   - Bitrate: 4000Kbps
//   - Container: MP4
//   - VideoSize: 720p
func NewTranscodeJobWithDefaults() *TranscodeJob {
	opts := NewTranscodeOptsWithDefaults()

	tj := &TranscodeJob{
		TranscodeOpts: opts,
	}

	return tj
}


// ApiKey returns the TranscodeJob instance configured with the provided apiKey.
func (tj *TranscodeJob) ApiKey(apiKey string) *TranscodeJob {
	tj.APIKey = apiKey
	return tj
}

// From returns the TranscodeJob instance configured with the Blob provided as an input file.
// A Blob represents a file located on either Amazon S3, Google GCP or Azure File.
func (tj *TranscodeJob) From(input *Blob) *TranscodeJob {
	tj.InputBlob = input
	return tj
}

// FromUrl returns the TranscodeJob instance configured with the url provided as an input file.
func (tj *TranscodeJob) FromUrl(inputUrl string) *TranscodeJob {
	tj.InputURL = inputUrl
	return tj
}

// To returns the TranscodeJob instance configured with the Blob provided as an output file.
// A Blob represents a file located on either Amazon S3, Google GCP or Azure File.
func (tj *TranscodeJob) To(output *Blob) *TranscodeJob {
	tj.OutputBlob = output
	return tj
}

// ToUrl returns the TranscodeJob instance configured with the url provided as an input file.
func (tj *TranscodeJob) ToUrl(outputUrl string) *TranscodeJob {
	tj.OutputURL = outputUrl
	return tj
}

// Webhooks returns the TranscodeJob instance configured with the success and failure
// Webhooks configured. Success and failure are string representing urls that MediaMachine
// are going to call in case of success or failure.
func (tj *TranscodeJob) Webhooks(success, failure string) *TranscodeJob {
	tj.SuccessURL = success
	tj.FailureURL = failure
	return tj
}

// Width returns the TranscodeJob instance configured with the output width of the video.
func (tj *TranscodeJob) Width(width uint16) *TranscodeJob {
	tj.WidthInt = width
	return tj
}

// Watermark returns the TranscodeJob instance configured with the Watermark to be used on the
// output video.
func (tj *TranscodeJob) Watermark(watermark *Watermark) *TranscodeJob {
	tj.TranscodeWatermark = watermark
	return tj
}

// WatermarkFromText returns the TranscodeJob instance configured with a Watermark text to be
// used on the output video. This Text will have a white color, font size of 12px, located on the bottom-left
// corner of the video, and an opacity of 80%.
func (tj *TranscodeJob) WatermarkFromText(text string) *TranscodeJob {
	w := &Watermark{
		WatermarkText:     text,
		WatermarkFontSize: 12,
		WatermarkColor:    "white",
		WatermarkOpacity:  0.8,
		WatermarkPosition: PositionBottomLeft,
	}

	tj.TranscodeWatermark = w
	return tj
}

// Opts returns the TranscodeJob instance configured tiwh the TranscodeOpts provided.
func (tj *TranscodeJob) Opts(opts *TranscodeOpts) *TranscodeJob {
	tj.TranscodeOpts = opts
	return tj
}

// Execute execute the transcode job, it returns the job and an error.
// In case of an error, job will be nil and error will contain the cause of the error.
func (tj *TranscodeJob) Execute() (*Job, error) {
	body, err := json.Marshal(tj)
	url := Settings.URL + "/transcode"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", Settings.userAgent)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	decoder := json.NewDecoder(resp.Body)

	// Job created successfully
	if resp.StatusCode == http.StatusOK {
		job := &Job{}
		if err := decoder.Decode(&job); err != nil {
			return nil, err
		}
		return job, nil
	} else {
		createErr := Error{}
		if err := decoder.Decode(&createErr); err != nil {
			return nil, err
		}
		return nil, createErr
	}
}