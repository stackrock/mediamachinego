package mediamachine

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

// SummaryType represent the possible output type of the summary.
type SummaryType = string

const (
	// Represents an output of type `gif`.
	SUMMARY_GIF SummaryType = "gif"
	// Represents an output of type `mp4`.
	SUMMARY_MP4 SummaryType = "mp4"
)

// SummaryJob creates an intelligent preview for the input video. The output can be either a GIF (SUMMARY_GIF)
// or a MP4 (SUMMARY_MP4).
// StackRock will intelligently process the video to create a summary for the input video.
// The input video location can be specified via the FromUrl or the From methods.
// Similarly, ToUrl or the To methods for output location.
// By default, the summary is the same {Width x Height} as the input video - use the Width method to customize.
// In the case of a MP4 output, Audio track is also preserved by default - callers can customize tho remove audio if
// needed with the RemoveAudio method.
type SummaryJob struct {
	JobRequest
	WidthInt uint16 `json:"width,omitempty"`
	SummaryWatermark *Watermark `json:"watermark,omitempty"`
	SummaryType SummaryType `json:"-"`
	SummaryRemoveAudio bool `json:"removeAudio,omitempty"`
}

// NewSummaryJobWithDefaults returns a new instance of SummaryJob configured with default values.
func NewSummaryJobWithDefaults() *SummaryJob {
	sj := &SummaryJob{
		SummaryType: SUMMARY_GIF,
	}
	return sj
}

// ApiKey returns the SummaryJob instance configured with the api key provided.
func (sj *SummaryJob) ApiKey(apiKey string) *SummaryJob {
	sj.APIKey = apiKey
	return sj
}


// From returns the SummaryJob instance configured with the Blob provided as an input file.
// A Blob represents a file located on either Amazon S3, Google GCP or Azure File.
func (sj *SummaryJob) From(input *Blob) *SummaryJob {
	sj.InputBlob = input
	return sj
}

// FromUrl returns the SummaryJob instance configured with the url provided as an input file.
func (sj *SummaryJob) FromUrl(inputUrl string) *SummaryJob {
	sj.InputURL = inputUrl
	return sj
}

// To returns the SummaryJob instance configured with the Blob provided as an output file.
// A Blob represents a file located on either Amazon S3, Google GCP or Azure File.
func (sj *SummaryJob) To(output *Blob) *SummaryJob {
	sj.OutputBlob = output
	return sj
}

// ToUrl returns the SummaryJob instance configured with the url provided as an input file.
func (sj *SummaryJob) ToUrl(outputUrl string) *SummaryJob {
	sj.OutputURL = outputUrl
	return sj
}

// Webhooks returns the SummaryJob instance configured with the success and failure
// Webhooks configured. Success and failure are string representing urls that MediaMachine
// are going to call in case of success or failure.
func (sj *SummaryJob) Webhooks(success, failure string) *SummaryJob {
	sj.SuccessURL = success
	sj.FailureURL = failure
	return sj
}

// Width returns the SummaryJob instance configured with the output width of the image.
func (sj *SummaryJob) Width(width uint16) *SummaryJob {
	sj.WidthInt = width
	return sj
}

// Watermark returns the SummaryJob instance configured with the Watermark to be used on the
// output gif/video.
func (sj *SummaryJob) Watermark(watermark *Watermark) *SummaryJob {
	sj.SummaryWatermark = watermark
	return sj
}

// WatermarkFromText returns the SummaryJob instance configured with a Watermark text to be
// used on the output gif/video. This Text will have white color, font size of 12px, located on the bottom-left
// corner of the gif/video, and an opacity of 80%.
func (sj *SummaryJob) WatermarkFromText(text string) *SummaryJob {
	w := &Watermark{
		WatermarkText:     text,
		WatermarkFontSize: 12,
		WatermarkColor:    "white",
		WatermarkOpacity:  0.8,
		WatermarkPosition: PositionBottomLeft,
	}

	sj.SummaryWatermark = w
	return sj
}

// Type returns the SummaryJob instance configured with the type of summary.
// Valid values are:
//  - SUMMARY_GIF output will be a gif.
//  - SUMMARY_MP4 output will be a mp4 video.
func (sj *SummaryJob) Type(summaryType SummaryType) *SummaryJob {
	sj.SummaryType = summaryType
	return sj
}

// RemoveAudio returns the SummaryJob instance configured to strip the video from the output
// video (will do nothing if the type of the SummaryJob is SUMMARY_GIF).
func (sj *SummaryJob) RemoveAudio(removeAudio bool) *SummaryJob {
	sj.SummaryRemoveAudio = removeAudio
	return sj
}

// Execute execute the summary job, it returns the job and an error.
// In case of an error, job will be nil and error will contain the cause of the error.
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