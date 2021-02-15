/*
Example usage of MediaMachine SDK for creating video Summaries.

This example shows how to use s3 for storage, however, using Azure blob store or GCP buckets is also supported:
Simply change the scheme in InputURL/OutputURL to "azure://example-bucket..." or "gcp://example-bucket..." according
to your requirements.
*/
package main

import (
	mediamachine "github.com/stackrock/mediamachinego"
	"github.com/stackrock/mediamachinego/colors"
	"log"
	"time"
)

// Use your StackRock API Key to initialize the MediaMachine SDK
const apiKey = "stackrock-web-test-ad85cef0-3414-11eb-80ae-415459e1a243"

//const apiKey = "stackrock_test_me"

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

	// If your video assets are served via a file server, you can directly use their urls. You can also mix-and-match
	// using file server urls and bucket stores.
	// See Thumbnail example for working with videos served via a file server.
	s3GIFSummaryJob, err := mm.Summary(mediamachine.SummaryConfig{
		Type:      mediamachine.SummaryTypeGif,
		InputURL:  "s3://example-bucket/my-awesome-video.mp4",
		OutputURL: "s3://example-bucket/my-awesome-video-summary.jpg",
		// account for example or to a different bucket if you generate keys specific to bucket etc.
		// Note: You can use a different set of creds for input and output if you want to upload to a totally different
		InputCreds:  creds,
		OutputCreds: creds,
		Watermark: mediamachine.WatermarkText{ // Optional
			Text:     "My Awesome Company",
			FontSize: 10,
			Color:    colors.Brown, // See docs for other color options
			Opacity:  0.5,          // Should be between [0,1]
			Position: mediamachine.PositionBottomLeft,
		},
		// You can opt to get notified via webhooks here or periodically check job status depending on your preferred setup
		SuccessURL: "https://example.com/mediamachine/jobdone",
		FailureURL: "https://example.com/mediamachine/jobfailed",
	})

	// Handle any errors returned during job creation
	if err != nil {
		log.Panicf("failed to create a summary job: %+v", err)
	}

	// Example function for waiting for job completion
	waitForJob := func(job mediamachine.Job, done chan struct{}) {
		defer close(done)

		for range time.NewTicker(time.Second * 60).C {
			status, err := job.FetchStatus()
			if err != nil {
				log.Printf("failed to fetch status for job: %s", job.ID)
				return
			}

			switch status {
			case mediamachine.JobStatusDone:
				log.Printf("summary is ready! JobId: %s", job.ID)
				return
			case mediamachine.JobStatusErrored:
				log.Printf("summary creation failed :( JobId: %s", job.ID)
				return
			}
		}
	}

	jobsDone := make(chan struct{})
	go waitForJob(s3GIFSummaryJob, jobsDone)

	// Wait for job to finish
	<-jobsDone

	log.Printf("All done!")
}