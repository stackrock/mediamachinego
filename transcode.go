package mediamachine

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type TranscodeEncoder = string
type TranscodeBitrate = string
type TranscodeContainer = string
type TranscodeVideoSize = string

const (
	ENCODER_H264 TranscodeEncoder = "h264"
	ENCODER_H265 TranscodeEncoder = "h265"
	ENCODER_VP8 TranscodeEncoder = "vp8"

	BITRATE_EIGHT_MEGAKBPS TranscodeBitrate = "8000"
	BITRATE_FOUR_MEGAKBPS TranscodeBitrate = "4000"
	BITRATE_ONE_MEGAKBPS TranscodeBitrate = "1000"

	CONTAINER_MP4 TranscodeContainer = "mp4"
	CONTAINER_WEBM TranscodeContainer = "webm"

	VIDEOSIZE_FULL_HD = "1080"
	VIDEOSIZE_HD = "720"
	VIDEOSIZE_SD = "480"
)

type TranscodeOpts struct {
	TranscodeEncoder string `json:"encoder,omitempty"`
	TranscodeBitrateKbps string `json:"bitrateKbps,omitempty"`
	TranscodeContainer string `json:"container,omitempty"`
	TranscodeVideoSize string `json:"videoSize,omitempty"`
}

func NewTranscodeOptsWithDefaults() *TranscodeOpts {
	to := &TranscodeOpts{
		TranscodeEncoder:     ENCODER_H264,
		TranscodeBitrateKbps: BITRATE_FOUR_MEGAKBPS,
		TranscodeContainer:   CONTAINER_MP4,
		TranscodeVideoSize:   VIDEOSIZE_HD,
	}

	return to
}

func (to *TranscodeOpts) Encoder(encoder TranscodeEncoder) *TranscodeOpts {
	to.TranscodeEncoder = encoder
	return to
}

func (to *TranscodeOpts) BitrateKbps(bitrate TranscodeBitrate) *TranscodeOpts {
	to.TranscodeBitrateKbps = bitrate
	return to
}

func (to *TranscodeOpts) Container(container TranscodeContainer) *TranscodeOpts {
	to.TranscodeContainer = container
	return to
}

func (to *TranscodeOpts) VideoSize(videoSize TranscodeVideoSize) *TranscodeOpts {
	to.TranscodeVideoSize = videoSize
	return to
}

type TranscodeJob struct {
	jobRequest
	WidthInt uint16 `json:"width,omitempty"`
	TranscodeWatermark *Watermark `json:"watermark,omitempty"`
	TranscodeOpts *TranscodeOpts `json:"transcode,omitemtpy"`
}

func NewTranscodeJobWithDefaults() *TranscodeJob {
	opts := NewTranscodeOptsWithDefaults()

	tj := &TranscodeJob{
		TranscodeOpts: opts,
	}

	return tj
}

func (tj *TranscodeJob) ApiKey(apiKey string) *TranscodeJob {
	tj.APIKey = apiKey
	return tj
}

func (tj *TranscodeJob) From(input *Blob) *TranscodeJob {
	tj.InputBlob = input
	return tj
}

func (tj *TranscodeJob) FromUrl(inputUrl string) *TranscodeJob {
	tj.InputURL = inputUrl
	return tj
}


func (tj *TranscodeJob) To(output *Blob) *TranscodeJob {
	tj.OutputBlob = output
	return tj
}

func (tj *TranscodeJob) ToUrl(outputUrl string) *TranscodeJob {
	tj.OutputURL = outputUrl
	return tj
}

func (tj *TranscodeJob) Webhooks(success, failure string) *TranscodeJob {
	tj.SuccessURL = success
	tj.FailureURL = failure
	return tj
}

func (tj *TranscodeJob) Width(width uint16) *TranscodeJob {
	tj.WidthInt = width
	return tj
}

func (tj *TranscodeJob) Watermark(watermark *Watermark) *TranscodeJob {
	tj.TranscodeWatermark = watermark
	return tj
}

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

func (tj *TranscodeJob) Opts(opts *TranscodeOpts) *TranscodeJob {
	tj.TranscodeOpts = opts
	return tj
}

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