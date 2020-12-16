package mediamachine

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

// ThumbnailJob - create an intelligent thumbnail for the input video.
// StackRock will intelligently search the video to find the best thumbnail and
// automatically try to skip empty/blank frames to provide the most appropriate thumbnail.
// The input video location can be specified via the FromUrl or the From method.
// Similarly, ToUrl or the To methods for thumbnail image output location.
// By default, the thumbnail is the same {Width x Height} as the input video - use the Width method to customize
type ThumbnailJob struct {
	JobRequest
	WidthInt uint16 `json:"width,omitempty"`
	ThumbnailWatermark *Watermark `json:"watermark,omitemtpy"`
}

// NewThumbnailJobWithDefaults returns a new instance of ThumbnailJob configured with default values
func NewThumbnailJobWithDefaults() *ThumbnailJob {
	tj := &ThumbnailJob{}
	return tj
}

// ApiKey returns a the ThumbnailJob instance configured with the provided apiKey.
func (tj *ThumbnailJob) ApiKey(apiKey string) *ThumbnailJob {
	tj.APIKey = apiKey
	return tj
}

// From returns the ThumbnailJob instance configured with the Blob provided as an input file.
// A Blob represents a file located on either Amazon S3, Google GCP or Azure File.
func (tj *ThumbnailJob) From(input *Blob) *ThumbnailJob {
	tj.InputBlob = input
	return tj
}

// FromUrl returns the ThumbnailJob instance configured with the url provided as an input file.
func (tj *ThumbnailJob) FromUrl(inputUrl string) *ThumbnailJob {
	tj.InputURL = inputUrl
	return tj
}

// To returns the ThumbnailJob instance configured with the Blob provided as an output file.
// A Blob represents a file located on either Amazon S3, Google GCP or Azure File.
func (tj *ThumbnailJob) To(output *Blob) *ThumbnailJob {
	tj.OutputBlob = output
	return tj
}

// ToUrl returns the ThumbnailJob instance configured with the url provided as an input file.
func (tj *ThumbnailJob) ToUrl(outputUrl string) *ThumbnailJob {
	tj.OutputURL = outputUrl
	return tj
}

// Webhooks returns the ThumbnailJob instance configured with the success and failure
// Webhooks configured. Success and failure are string representing urls that MediaMachine
// are going to call in case of success or failure.

func (tj *ThumbnailJob) Webhooks(success, failure string) *ThumbnailJob {
	tj.SuccessURL = success
	tj.FailureURL = failure
	return tj
}

// Width returns the ThumbnailJob instance configured with the output width of the image.
func (tj *ThumbnailJob) Width(width uint16) *ThumbnailJob {
	tj.WidthInt = width
	return tj
}

// Watermark returns the ThumbnailJob instance configured with the Watermark to be used on the
// output image.
func (tj *ThumbnailJob) Watermark(watermark *Watermark) *ThumbnailJob {
	tj.ThumbnailWatermark = watermark
	return tj
}

// WatermarkFromText returns the ThumbnailJob instance configured with a Watermark text to be
// used on the output image. This Text will have white color, font size of 12px, located on the bottom-left
// corner of the image, and an opacity of 80%.
func (tj *ThumbnailJob) WatermarkFromText(text string) *ThumbnailJob {
	w := &Watermark{
		WatermarkText:     text,
		WatermarkFontSize: 12,
		WatermarkColor:    "white",
		WatermarkOpacity:  0.8,
		WatermarkPosition: PositionBottomLeft,
	}

	tj.ThumbnailWatermark = w
	return tj
}

// Execute execute the thumbnail job, it returns the job and an error.
// In case of an error, job will be nil and error will contain the cause of the error.
func (tj *ThumbnailJob) Execute() (*Job, error) {
	body, err := json.Marshal(tj)
	url := Settings.URL + "/thumbnail"
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