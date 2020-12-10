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
	It("tracer - thumbnail", func() {
		var STACKROCK_API_KEY = os.Getenv("STACKROCK_API_KEY")
		var BUCKET = os.Getenv("BUCKET")
		var INPUT_KEY = os.Getenv("INPUT_KEY")
		var OUTPUT_KEY = os.Getenv("OUTPUT_KEY_THUMBNAIL")
		var AWS_REGION = os.Getenv("AWS_REGION")
		var AWS_ACCESS_KEY_ID = os.Getenv("AWS_ACCESS_KEY_ID")
		var AWS_SECRET_ACCESS_KEY = os.Getenv("AWS_SECRET_ACCESS_KEY")

		inputFile := NewS3BlobWithDefaults().Bucket(BUCKET).Key(INPUT_KEY).AWSCredentials(AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, AWS_REGION)
		outputFile := NewS3BlobWithDefaults().Bucket(BUCKET).Key(OUTPUT_KEY).AWSCredentials(AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, AWS_REGION)

		job, err := NewThumbnailJobWithDefaults().ApiKey(STACKROCK_API_KEY).From(inputFile).To(outputFile).Width(150).WatermarkFromText("stackrock.io").Execute()

		Expect(err).To(BeNil())

		fmt.Printf("Job id: %s\n" , job.Id)

		_, err = job.Status()
		Expect(err).To(BeNil())


		checkFn := func() string {
			status, err := job.Status()
			Expect(err).To(BeNil())
			return status
		}

		Eventually(checkFn, "5m").Should(Equal("done"))
	})

	/*
	 * Tracer Bullet for a Summary Gif job.
	 * We use this job internally at StackRock for two reasons:
	 *  1) To keep the SDK in sync with API
	 *  2) To Test our API is running as expected
	 */
	It("tracer - Summary Gif", func() {
		var STACKROCK_API_KEY = os.Getenv("STACKROCK_API_KEY")
		var BUCKET = os.Getenv("BUCKET")
		var INPUT_KEY = os.Getenv("INPUT_KEY")
		var OUTPUT_KEY = os.Getenv("OUTPUT_KEY_SUMMARY_GIF")
		var AWS_REGION = os.Getenv("AWS_REGION")
		var AWS_ACCESS_KEY_ID = os.Getenv("AWS_ACCESS_KEY_ID")
		var AWS_SECRET_ACCESS_KEY = os.Getenv("AWS_SECRET_ACCESS_KEY")

		inputFile := NewS3BlobWithDefaults().Bucket(BUCKET).Key(INPUT_KEY).AWSCredentials(AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, AWS_REGION)
		outputFile := NewS3BlobWithDefaults().Bucket(BUCKET).Key(OUTPUT_KEY).AWSCredentials(AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, AWS_REGION)

		job, err := NewSummaryJobWithDefaults().ApiKey(STACKROCK_API_KEY).Type(SUMMARY_GIF).From(inputFile).To(outputFile).Width(150).WatermarkFromText("stackrock.io").Execute()

		Expect(err).To(BeNil())

		fmt.Printf("Job id: %s\n" , job.Id)

		_, err = job.Status()
		Expect(err).To(BeNil())

		checkFn := func() string {
			status, err := job.Status()
			Expect(err).To(BeNil())
			return status
		}

		Eventually(checkFn, "5m").Should(Equal("done"))
	})

	/*
	 * Tracer Bullet for a Summary MP4 job.
	 * We use this job internally at StackRock for two reasons:
	 *  1) To keep the SDK in sync with API
	 *  2) To Test our API is running as expected
	 */
	It("tracer - Summary MP4", func() {
		var STACKROCK_API_KEY = os.Getenv("STACKROCK_API_KEY")
		var BUCKET = os.Getenv("BUCKET")
		var INPUT_KEY = os.Getenv("INPUT_KEY")
		var OUTPUT_KEY = os.Getenv("OUTPUT_KEY_SUMMARY_MP4")
		var AWS_REGION = os.Getenv("AWS_REGION")
		var AWS_ACCESS_KEY_ID = os.Getenv("AWS_ACCESS_KEY_ID")
		var AWS_SECRET_ACCESS_KEY = os.Getenv("AWS_SECRET_ACCESS_KEY")

		inputFile := NewS3BlobWithDefaults().Bucket(BUCKET).Key(INPUT_KEY).AWSCredentials(AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, AWS_REGION)
		outputFile := NewS3BlobWithDefaults().Bucket(BUCKET).Key(OUTPUT_KEY).AWSCredentials(AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, AWS_REGION)

		job, err := NewSummaryJobWithDefaults().ApiKey(STACKROCK_API_KEY).Type(SUMMARY_MP4).From(inputFile).To(outputFile).Width(150).WatermarkFromText("stackrock.io").Execute()

		Expect(err).To(BeNil())

		fmt.Printf("Job id: %s\n" , job.Id)

		_, err = job.Status()
		Expect(err).To(BeNil())

		checkFn := func() string {
			status, err := job.Status()
			Expect(err).To(BeNil())
			return status
		}

		Eventually(checkFn, "2m").Should(Equal("done"))
	})

	/*
	 * Tracer Bullet for a Transcode job.
	 * We use this job internally at StackRock for two reasons:
	 *  1) To keep the SDK in sync with API
	 *  2) To Test our API is running as expected
	 */
	It("tracer - Transcode", func() {
		var STACKROCK_API_KEY = os.Getenv("STACKROCK_API_KEY")
		var BUCKET = os.Getenv("BUCKET")
		var INPUT_KEY = os.Getenv("TRANSCODE_INPUT_KEY")
		var OUTPUT_KEY = os.Getenv("OUTPUT_KEY_TRANSCODE")
		var AWS_REGION = os.Getenv("AWS_REGION")
		var AWS_ACCESS_KEY_ID = os.Getenv("AWS_ACCESS_KEY_ID")
		var AWS_SECRET_ACCESS_KEY = os.Getenv("AWS_SECRET_ACCESS_KEY")

		inputFile := NewS3BlobWithDefaults().Bucket(BUCKET).Key(INPUT_KEY).AWSCredentials(AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, AWS_REGION)
		outputFile := NewS3BlobWithDefaults().Bucket(BUCKET).Key(OUTPUT_KEY).AWSCredentials(AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, AWS_REGION)

		transcodeOpts := NewTranscodeOptsWithDefaults().Encoder(ENCODER_H264).BitrateKbps(BITRATE_FOUR_MEGAKBPS).Container(CONTAINER_MP4).VideoSize(VIDEOSIZE_HD)

		job, err := NewTranscodeJobWithDefaults().ApiKey(STACKROCK_API_KEY).From(inputFile).To(outputFile).Width(150).WatermarkFromText("stackrock.io").Opts(transcodeOpts).Execute()

		Expect(err).To(BeNil())

		fmt.Printf("Job id: %s\n" , job.Id)

		_, err = job.Status()
		Expect(err).To(BeNil())

		checkFn := func() string {
			status, err := job.Status()
			Expect(err).To(BeNil())
			return status
		}

		Eventually(checkFn, "10m").Should(Equal("done"))
	})
})

