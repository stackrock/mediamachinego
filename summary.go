package mediamachine

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

// SummaryGIF - create an intelligent summary for the input video (outputs a Gif)
// StackRock will intelligently process the video to create a summary for the input video.
// The input video location can be specified via the InputURL or the InputBlob options.
// Similarly, OutputURL or the OutputBlob options for thumbnail image output location.
// By default, the summary is the same {Width x Height} as the input video - use the options to customize
// To generate a summary as a movie (optionally with sound), use SummaryMP4

//}

// SummaryMP4 - create an intelligent summary for the input video (outputs an mp4 video)
// StackRock will intelligently process the video to create a summary for the input video.
// The input video location can be specified via the InputURL or the InputBlob options.
// Similarly, OutputURL or the OutputBlob options for thumbnail image output location.
// By default, the summary is the same {Width x Height} as the input video - use the options to customize
// Audio track is also preserved by default - callers can customize to remove audio if needed.
type SummaryType = string

const (
	SUMMARY_GIF SummaryType = "gif"
	SUMMARY_MP4 SummaryType = "mp4"
)

type SummaryJob struct {
	jobRequest
	WidthInt uint16 `json:"width,omitempty"`
	Watermark *Watermark `json:"watermark,omitempty"`
	SummaryType SummaryType `json:"-"`
}

func NewSummaryJobWithDefaults() *SummaryJob {
	sj := &SummaryJob{
		SummaryType: SUMMARY_GIF,
	}
	return sj
}

func (sj *SummaryJob) ApiKey(apiKey string) *SummaryJob {
	sj.APIKey = apiKey
	return sj
}

func (sj *SummaryJob) From(input *Blob) *SummaryJob {
	sj.InputBlob = input
	return sj
}

func (sj *SummaryJob) FromUrl(inputUrl string) *SummaryJob {
	sj.InputURL = inputUrl
	return sj
}


func (sj *SummaryJob) To(output *Blob) *SummaryJob {
	sj.OutputBlob = output
	return sj
}

func (sj *SummaryJob) ToUrl(outputUrl string) *SummaryJob {
	sj.OutputURL = outputUrl
	return sj
}

func (sj *SummaryJob) Webhooks(success, failure string) *SummaryJob {
	sj.SuccessURL = success
	sj.FailureURL = failure
	return sj
}

func (sj *SummaryJob) Width(width uint16) *SummaryJob {
	sj.WidthInt = width
	return sj
}

func (sj *SummaryJob) WatermarkFromText(text string) *SummaryJob {
	w := &Watermark{
		WatermarkText:     text,
		WatermarkFontSize: 12,
		WatermarkColor:    "white",
		WatermarkOpacity:  0.8,
		WatermarkPosition: PositionBottomLeft,
	}

	sj.Watermark = w
	return sj
}

func (sj *SummaryJob) Type(summaryType SummaryType) *SummaryJob {
	sj.SummaryType = summaryType
	return sj
}

func (sj *SummaryJob) Execute() (*Job, error) {
	body, err := json.Marshal(sj)
	url := Settings.URL + "/summary/" + sj.SummaryType
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