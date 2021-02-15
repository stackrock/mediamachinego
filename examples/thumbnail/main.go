/*
Example usage of MediaMachine SDK for extracting intelligent Thumbnails.

This example shows how to use s3 for storage, however, using Azure blob store or GCP buckets is also supported:
Simply change the scheme in InputURL/OutputURL to "azure://example-bucket..." or "gcp://example-bucket..." according
to your requirements.
*/
package main

import (
	"github.com/stackrock/mediamachinego/colors"
	"github.com/stackrock/mediamachinego/mediamachine"
	"log"
	"time"
)

// Use your MediaMachine API Key to initialize the SDK
const apiKey = "my-mediamachine-key"

func main() {
	mm := mediamachine.MediaMachine{APIKey: apiKey}

	// Using S3: input video from s3, output uploaded to s3
	// It is a good security practice to make narrow scoped AWS access keys
	// that only restrict access to a specific bucket (or even specific prefixes and objects if needed).
	creds := mediamachine.CredsAWS{
		AccessKeyID:     "my-aws-access-key-id",
		SecretAccessKey: "my-aws-secret-access-key",
		Region:          "my-aws-region",
	}

	s3Job, err := mm.Thumbnail(mediamachine.ThumbnailConfig{
		InputURL:  "s3://example-bucket/my-awesome-video.mp4",
		OutputURL: "s3://example-bucket/my-awesome-video-thumbnail.jpg",
		// account for example or to a different bucket if you generate keys specific to bucket etc.
		// Note: You can use a different set of creds for input and output if you want to upload to a totally different
		InputCreds:  creds,
		OutputCreds: creds,
		Width:       500, // Defaults to size of input
		Watermark: mediamachine.WatermarkText{
			Text:      "My Awesome Company",
			FontSize:  10,
			FontColor: colors.Brown, // See docs for other color options
			Opacity:   0.5,          // Should be between [0,1]
			Position:  mediamachine.PositionBottomLeft,
		},
		// You can opt to get notified via webhooks here or periodically check job status depending on your preferred setup
		SuccessURL: "https://example.com/mediamachine/jobdone",
		FailureURL: "https://example.com/mediamachine/jobfailed",
	})

	// Handle any errors returned during job creation
	if err != nil {
		log.Panicf("failed to create a thumbnail job: %+v", err)
	}

	// If your video assets are served via a file server, you can directly use their urls. You can also mix-and-match
	// using file server urls and bucket stores.
	// Example for videos served via a file server:
	fileServerJob, err := mm.Thumbnail(mediamachine.ThumbnailConfig{
		InputURL:  "https://example.com/my-awesome-video.mp4",
		OutputURL: "https://example.com/my-awesome-video-thumbnail.jpg",
		Width:     500, // Defaults to size of input
		Watermark: mediamachine.WatermarkText{
			Text:      "My Awesome Company",
			FontSize:  10,
			FontColor: colors.Brown, // See docs for other color options
			Opacity:   0.5,          // Should be between [0,1]
			Position:  mediamachine.PositionBottomLeft,
		},
		// You can opt to get notified via webhooks here or periodically check job status depending on your preferred setup
		SuccessURL: "https://example.com/mediamachine/jobdone",
		FailureURL: "https://example.com/mediamachine/jobfailed",
	})

	// Handle any errors returned during job creation
	if err != nil {
		log.Panicf("failed to create a thumbnail job: %+v", err)
	}

	// Example function for waiting for job completion
	waitForJob := func(job mediamachine.Job, done chan struct{}) {
		for range time.NewTicker(time.Second * 60).C {
			status, err := job.FetchStatus()
			if err != nil {
				log.Printf("failed to fetch status for job: %s", job.ID)
				done <- struct{}{}
				return
			}

			switch status {
			case mediamachine.JobStatusDone:
				log.Printf("thumbnail is ready! JobId: %s", job.ID)
				done <- struct{}{}
				return
			case mediamachine.JobStatusErrored:
				log.Printf("thumbnail creation failed :( JobId: %s", job.ID)
				done <- struct{}{}
				return
			}
		}
	}

	jobsDone := make(chan struct{}, 2)
	go waitForJob(s3Job, jobsDone)
	go waitForJob(fileServerJob, jobsDone)

	// Wait for both jobs to finish
	<-jobsDone
	<-jobsDone

	log.Printf("All done!")
}
