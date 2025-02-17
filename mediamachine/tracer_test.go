package mediamachine_test

import (
	"fmt"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/stackrock/mediamachinego/colors"
	"github.com/stackrock/mediamachinego/mediamachine"
)

var MEDIAMACHINE_API_KEY = os.Getenv("MEDIAMACHINE_API_KEY")
var BUCKET = os.Getenv("BUCKET")
var INPUT_KEY = os.Getenv("INPUT_KEY")
var OUTPUT_KEY = os.Getenv("OUTPUT_KEY_SUMMARY_GIF")
var AWS_REGION = os.Getenv("AWS_REGION")
var AWS_ACCESS_KEY_ID = os.Getenv("AWS_ACCESS_KEY_ID")
var AWS_SECRET_ACCESS_KEY = os.Getenv("AWS_SECRET_ACCESS_KEY")
var mm = mediamachine.MediaMachine{APIKey: MEDIAMACHINE_API_KEY}

var _ = Describe("tracer", func() {
	// Using S3: input video from s3, output uploaded to s3
	// It is a good security practice to make narrow scoped AWS access keys
	// that only restrict access to a specific bucket (or even specific prefixes and objects if needed).
	creds := mediamachine.CredsAWS{
		AccessKeyID:     AWS_ACCESS_KEY_ID,
		SecretAccessKey: AWS_SECRET_ACCESS_KEY,
		Region:          AWS_REGION,
	}

	/*
	 * Tracer Bullet for a thumbnail-s3-compatible-store job.
	 * We use this job internally at MediaMachine for two reasons:
	 *  1) To keep the SDK in sync with API
	 *  2) To Test our API is running as expected
	 */
	It("tracer - thumbnail-s3-compatible-store", func() {
		job, err := mm.Thumbnail(mediamachine.ThumbnailConfig{
			InputURL:  fmt.Sprintf("s3://%s/%s", BUCKET, INPUT_KEY),
			OutputURL: fmt.Sprintf("s3://%s/%s", BUCKET, OUTPUT_KEY),
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
		})

		Expect(err).To(BeNil())

		fmt.Printf("Job id: %s\n", job.ID)

		_, err = job.FetchStatus()
		Expect(err).To(BeNil())

		checkFn := func() string {
			status, err := job.FetchStatus()
			Expect(err).To(BeNil())
			return status
		}

		Eventually(checkFn, "5m").Should(Equal(mediamachine.JobStatusDone))
	})

	/*
	 * Tracer Bullet for a Summary Gif job.
	 * We use this job internally at MediaMachine for two reasons:
	 *  1) To keep the SDK in sync with API
	 *  2) To Test our API is running as expected
	 */
	It("tracer - Summary Gif", func() {
		job, err := mm.SummaryGIF(mediamachine.SummaryConfig{
			InputURL:  fmt.Sprintf("s3://%s/%s", BUCKET, INPUT_KEY),
			OutputURL: fmt.Sprintf("s3://%s/%s", BUCKET, OUTPUT_KEY),
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
		})

		Expect(err).To(BeNil())

		fmt.Printf("Job id: %s\n", job.ID)

		_, err = job.FetchStatus()
		Expect(err).To(BeNil())

		checkFn := func() string {
			status, err := job.FetchStatus()
			Expect(err).To(BeNil())
			return status
		}

		Eventually(checkFn, "5m").Should(Equal("done"))
	})

	/*
	 * Tracer Bullet for a Summary MP4 job.
	 * We use this job internally at MediaMachine for two reasons:
	 *  1) To keep the SDK in sync with API
	 *  2) To Test our API is running as expected
	 */
	It("tracer - Summary MP4", func() {
		job, err := mm.SummaryMP4(mediamachine.SummaryConfig{
			InputURL:  fmt.Sprintf("s3://%s/%s", BUCKET, INPUT_KEY),
			OutputURL: fmt.Sprintf("s3://%s/%s", BUCKET, OUTPUT_KEY),
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
		})

		Expect(err).To(BeNil())

		fmt.Printf("Job id: %s\n", job.ID)

		_, err = job.FetchStatus()
		Expect(err).To(BeNil())

		checkFn := func() string {
			status, err := job.FetchStatus()
			Expect(err).To(BeNil())
			return status
		}

		Eventually(checkFn, "5m").Should(Equal("done"))
	})

	/*
	 * Tracer Bullet for a Transcode job.
	 * We use this job internally at MediaMachine for two reasons:
	 *  1) To keep the SDK in sync with API
	 *  2) To Test our API is running as expected
	 */
	It("tracer - Transcode", func() {
		job, err := mm.Transcode(mediamachine.TranscodeConfig{
			InputURL:    fmt.Sprintf("s3://%s/%s", BUCKET, INPUT_KEY),
			OutputURL:   fmt.Sprintf("s3://%s/%s", BUCKET, OUTPUT_KEY),
			InputCreds:  creds,
			OutputCreds: creds,
			Width:       500, // Defaults to size of input
			Height:      400,
			Container:   mediamachine.ContainerMP4,
			Encoder:     mediamachine.EncoderH264,
			BitrateKBPS: mediamachine.Bitrate1Mbps,
		})

		Expect(err).To(BeNil())

		fmt.Printf("Job id: %s\n", job.ID)

		_, err = job.FetchStatus()
		Expect(err).To(BeNil())

		checkFn := func() string {
			status, err := job.FetchStatus()
			Expect(err).To(BeNil())
			return status
		}

		Eventually(checkFn, "5m").Should(Equal("done"))
	})
})
