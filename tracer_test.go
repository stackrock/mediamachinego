package mediamachine

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
)

var _ = Describe("tracer", func() {
	/*
	 * Tracer Bullet for a thumbnail job.
	 * We use this job internally at StackRock for two reasons:
	 *  1) To keep the SDK in sync with API
	 *  2) To Test our API is running as expected
	 */
	var _ = Describe("tracer - thumbnail", func() {
		var STACKROCK_API_KEY = os.Getenv("STACKROCK_API_KEY")
		var BUCKET = os.Getenv("BUCKET")
		var INPUT_KEY = os.Getenv("INPUT_KEY")
		var OUTPUT_KEY = os.Getenv("OUTPUT_KEY")
		var AWS_REGION = os.Getenv("AWS_REGION")
		var AWS_ACCESS_KEY_ID = os.Getenv("AWS_ACCESS_KEY_ID")
		var AWS_SECRET_ACCESS_KEY = os.Getenv("AWS_SECRET_ACCESS_KEY")

		inputFile := NewS3BlobWithDefaults().Bucket(BUCKET).Key(INPUT_KEY).AWSCredentials(AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, AWS_REGION)
		outputFile := NewS3BlobWithDefaults().Bucket(BUCKET).Key(OUTPUT_KEY).AWSCredentials(AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, AWS_REGION)

		job, err := NewThumbnailJobWithDefaults().ApiKey(STACKROCK_API_KEY).From(inputFile).To(outputFile).Width(150).WatermarkFromText("stackrock.io").Execute()

		fmt.Println("---- after creating job ----")
		fmt.Printf("%+v\n", job)
		fmt.Printf("%+v\n", err)
		fmt.Println("--------------------------------")

		Expect(err).To(BeNil())

		fmt.Printf("Job id: %s\n" , job.Id)
		status, err := job.Status()
		Expect(err).To(BeNil())
		fmt.Println(status)
		Eventually(job.Status()).Should(Equal("done"))
	})

	//var _ = Describe("summary", func() {
	//
	//})
	//
	//var _ = Describe("transcode", func() {
	//
	//})
})

