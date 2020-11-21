package mediamachine

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type ThumbnailJob struct {
	jobRequest
	WidthInt uint16 `json:"width,omitempty"`
	Watermark *Watermark `json:"watermark,omitemtpy"`
}
// Thumbnail - create an intelligent thumbnail for the input video.
// StackRock will intelligently search the video to find the best thumbnail and
// automatically try to skip empty/blank frames to provide the most appropriate thumbnail.
// The input video location can be specified via the InputURL or the InputBlob options.
// Similarly, OutputURL or the OutputBlob options for thumbnail image output location.
// By default, the thumbnail is the same {Width x Height} as the input video - use the options to customize
func NewThumbnailJobWithDefaults() *ThumbnailJob {
	tj := &ThumbnailJob{}
	return tj
}

func (tj *ThumbnailJob) ApiKey(apiKey string) *ThumbnailJob {
	tj.APIKey = apiKey
	return tj
}

func (tj *ThumbnailJob) From(input *Blob) *ThumbnailJob {
	tj.InputBlob = input
	return tj
}

func (tj *ThumbnailJob) FromUrl(inputUrl string) *ThumbnailJob {
	tj.InputURL = inputUrl
	return tj
}


func (tj *ThumbnailJob) To(output *Blob) *ThumbnailJob {
	tj.OutputBlob = output
	return tj
}

func (tj *ThumbnailJob) ToUrl(outputUrl string) *ThumbnailJob {
	tj.OutputURL = outputUrl
	return tj
}

func (tj *ThumbnailJob) Webhooks(success, failure string) *ThumbnailJob {
	tj.SuccessURL = success
	tj.FailureURL = failure
	return tj
}

func (tj *ThumbnailJob) Width(width uint16) *ThumbnailJob {
	tj.WidthInt = width
	return tj
}

func (tj *ThumbnailJob) WatermarkFromText(text string) *ThumbnailJob {
	w := &Watermark{
		WatermarkText:     text,
		WatermarkFontSize: 12,
		WatermarkColor:    "white",
		WatermarkOpacity:  0.8,
		WatermarkPosition: PositionBottomLeft,
	}

	tj.Watermark = w
	return tj
}

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