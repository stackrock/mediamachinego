package mediamachine

import "encoding/json"

// SummaryGIF - create an intelligent summary for the input video (outputs a Gif)
// StackRock will intelligently process the video to create a summary for the input video.
// The input video location can be specified via the InputURL or the InputBlob options.
// Similarly, OutputURL or the OutputBlob options for thumbnail image output location.
// By default, the summary is the same {Width x Height} as the input video - use the options to customize
// To generate a summary as a movie (optionally with sound), use SummaryMP4
func SummaryGIF(options ...JobOpt) (Job, error) {
	jr := jobRequest{}
	for _, opt := range options {
		opt(jr)
	}
	body, err := json.Marshal(jr)
	if err != nil {
		return Job{}, err
	}
	return sendRequest(SvcSummaryGIF, body)
}

// SummaryMP4 - create an intelligent summary for the input video (outputs an mp4 video)
// StackRock will intelligently process the video to create a summary for the input video.
// The input video location can be specified via the InputURL or the InputBlob options.
// Similarly, OutputURL or the OutputBlob options for thumbnail image output location.
// By default, the summary is the same {Width x Height} as the input video - use the options to customize
// Audio track is also preserved by default - callers can customize to remove audio if needed.
func SummaryMP4(options ...JobOpt) (Job, error) {
	jr := jobRequest{}
	for _, opt := range options {
		opt(jr)
	}
	body, err := json.Marshal(jr)
	if err != nil {
		return Job{}, err
	}
	return sendRequest(SvcSummaryMP4, body)
}
