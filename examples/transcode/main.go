/*
Example usage of MediaMachine SDK for video transcoding.

This example shows how to use s3 for storage, however, using Azure blob store or GCP buckets is also supported:
Simply change the scheme in InputURL/OutputURL to "azure://example-bucket..." or "gcp://example-bucket..." according
to your requirements.
*/
package main

import (
	"github.com/stackrock/mediamachinego/mediamachine"
	"log"
	"time"
)

// Use your StackRock API Key to initialize the MediaMachine SDK
const apiKey = "my-stackrock-key"

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
	s3GIFTranscodeJob, err := mm.Transcode(mediamachine.TranscodeConfig{
		InputURL:  "s3://example-bucket/my-awesome-video.mp4",
		OutputURL: "s3://example-bucket/my-awesome-video-transcode.jpg",
		// account for example or to a different bucket if you generate keys specific to bucket etc.
		// Note: You can use a different set of creds for input and output if you want to upload to a totally different
		InputCreds:  creds,
		OutputCreds: creds,

		// Make sure the encoder and container are compatible with each other.
		Container: mediamachine.ContainerMP4,
		Encoder:   mediamachine.EncoderH264,

		// You can opt to get notified via webhooks here or periodically check job status depending on your preferred setup
		SuccessURL: "https://example.com/mediamachine/jobdone",
		FailureURL: "https://example.com/mediamachine/jobfailed",
	})

	// Handle any errors returned during job creation
	if err != nil {
		log.Panicf("failed to create a transcode job: %+v", err)
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
				log.Printf("transcode is ready! JobId: %s", job.ID)
				return
			case mediamachine.JobStatusErrored:
				log.Printf("transcode creation failed :( JobId: %s", job.ID)
				return
			}
		}
	}

	jobsDone := make(chan struct{})
	go waitForJob(s3GIFTranscodeJob, jobsDone)

	// Wait for job to finish
	<-jobsDone

	log.Printf("All done!")
}
