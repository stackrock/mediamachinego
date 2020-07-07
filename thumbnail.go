package mediamachine

import (
	"encoding/json"
)

// Thumbnail - create an intelligent thumbnail for the input video.
// StackRock will intelligently search the video to find the best thumbnail and
// automatically try to skip empty/blank frames to provide the most appropriate thumbnail.
// The input video location can be specified via the InputURL or the InputBlob options.
// Similarly, OutputURL or the OutputBlob options for thumbnail image output location.
// By default, the thumbnail is the same {Width x Height} as the input video - use the options to customize
func Thumbnail(options ...JobOpt) (Job, error) {
	jr := jobRequest{}
	for _, opt := range options {
		opt(jr)
	}
	body, err := json.Marshal(jr)
	if err != nil {
		return Job{}, err
	}
	return sendRequest(SvcThumbnail, body)
}
