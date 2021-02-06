package mediamachine_test
//
//import (
//	"encoding/json"
//	"net/http"
//	"time"
//
//	. "github.com/onsi/ginkgo"
//	. "github.com/onsi/gomega"
//	"github.com/onsi/gomega/ghttp"
//
//	mediamachine "github.com/stackrock/mediamachinego"
//)
//
//var _ = Describe("Sdk", func() {
//	var server = ghttp.NewServer()
//	mediamachine.Settings.URL = server.URL()
//	awsCreds := mediamachine.CredsAWS{
//		AccessKeyID:     "testkey",
//		SecretAccessKey: "testsecret",
//		Region:          "us-east-1",
//	}
//	azureCreds := mediamachine.CredsAzure{
//		AccountName: "test-account",
//		AccountKey:  "test-key",
//	}
//	gcpCreds := json.RawMessage(`{"google-creds-file-contents":"contents"}`)
//
//	It("sends a simple thumbnail-s3-compatible-store request", func() {
//		server.AppendHandlers(ghttp.RespondWithJSONEncoded(http.StatusCreated, mediamachine.Job{
//			Id:        "test-job",
//			CreatedAt: time.Now(),
//		}))
//
//		job, err := mediamachine.Thumbnail(
//			mediamachine.APIKey("SECRET_API_KEY"),
//			mediamachine.Webhooks("http://mytestserver/success", "http://mytestserver/failure"),
//			mediamachine.InputURL("http://mytestserver/myvideo.mp4"),
//			mediamachine.OutputURL("http://mytestserver/myvideo.jpg"),
//		)
//		Expect(err).ToNot(HaveOccurred())
//		Expect(job.FetchStatus).To(Equal("queued"))
//		Expect(job.Id).To(Equal("test-job"))
//	})
//
//	It("sends a thumbnail-s3-compatible-store using blob stores", func() {
//		server.AppendHandlers(
//			ghttp.RespondWithJSONEncoded(http.StatusCreated, mediamachine.Job{
//				Id:        "test-job",
//				CreatedAt: time.Now(),
//			}),
//			ghttp.RespondWithJSONEncoded(http.StatusCreated, mediamachine.Job{
//				Id:        "test-job",
//				CreatedAt: time.Now(),
//			}),
//			ghttp.RespondWithJSONEncoded(http.StatusCreated, mediamachine.Job{
//				Id:        "test-job",
//				CreatedAt: time.Now(),
//			}),
//		)
//
//		job, err := mediamachine.Thumbnail(
//			mediamachine.APIKey("SECRET_API_KEY"),
//			mediamachine.Webhooks("http://mytestserver/success", "http://mytestserver/failure"),
//			mediamachine.S3Input("testbucket", "/foo/bar/testvideo.mp4", awsCreds),
//			mediamachine.S3Output("testbucket", "/foo/bar/testvideo.jpg", awsCreds),
//		)
//		Expect(err).ToNot(HaveOccurred())
//		Expect(job.FetchStatus).To(Equal("queued"))
//		Expect(job.Id).To(Equal("test-job"))
//
//		job, err = mediamachine.Thumbnail(
//			mediamachine.APIKey("SECRET_API_KEY"),
//			mediamachine.Webhooks("http://mytestserver/success", "http://mytestserver/failure"),
//			mediamachine.AzureInput("testbucket", "/foo/bar/testvideo.mp4", azureCreds),
//			mediamachine.AzureOutput("testbucket", "/foo/bar/testvideo.jpg", azureCreds),
//		)
//		Expect(err).ToNot(HaveOccurred())
//		Expect(job.FetchStatus).To(Equal("queued"))
//		Expect(job.Id).To(Equal("test-job"))
//
//		job, err = mediamachine.Thumbnail(
//			mediamachine.APIKey("SECRET_API_KEY"),
//			mediamachine.Webhooks("http://mytestserver/success", "http://mytestserver/failure"),
//			mediamachine.GCPInput("testbucket", "/foo/bar/testvideo.mp4", gcpCreds),
//			mediamachine.GCPOutput("testbucket", "/foo/bar/testvideo.jpg", gcpCreds),
//		)
//		Expect(err).ToNot(HaveOccurred())
//		Expect(job.FetchStatus).To(Equal("queued"))
//		Expect(job.Id).To(Equal("test-job"))
//	})
//
//	It("sends a gif summary request", func() {
//		server.AppendHandlers(ghttp.RespondWithJSONEncoded(http.StatusCreated, mediamachine.Job{
//			Id:        "test-job",
//			CreatedAt: time.Now(),
//		}))
//
//		job, err := mediamachine.SummaryGIF(
//			mediamachine.APIKey("SECRET_API_KEY"),
//			mediamachine.Webhooks("http://mytestserver/success", "http://mytestserver/failure"),
//			mediamachine.InputURL("http://mytestserver/myvideo.mp4"),
//			mediamachine.OutputURL("http://mytestserver/myvideo.jpg"),
//		)
//		Expect(err).ToNot(HaveOccurred())
//		Expect(job.FetchStatus).To(Equal("queued"))
//		Expect(job.Id).To(Equal("test-job"))
//	})
//
//	It("sends an mp4 summary request using blob stores", func() {
//		server.AppendHandlers(
//			ghttp.RespondWithJSONEncoded(http.StatusCreated, mediamachine.Job{
//				Id:        "test-job",
//				CreatedAt: time.Now(),
//			}),
//			ghttp.RespondWithJSONEncoded(http.StatusCreated, mediamachine.Job{
//				Id:        "test-job",
//				CreatedAt: time.Now(),
//			}),
//			ghttp.RespondWithJSONEncoded(http.StatusCreated, mediamachine.Job{
//				Id:        "test-job",
//				CreatedAt: time.Now(),
//			}),
//		)
//
//		job, err := mediamachine.SummaryMP4(
//			mediamachine.APIKey("SECRET_API_KEY"),
//			mediamachine.Webhooks("http://mytestserver/success", "http://mytestserver/failure"),
//			mediamachine.S3Input("testbucket", "/foo/bar/testvideo.mp4", awsCreds),
//			mediamachine.S3Output("testbucket", "/foo/bar/testvideo.jpg", awsCreds),
//		)
//		Expect(err).ToNot(HaveOccurred())
//		Expect(job.FetchStatus).To(Equal("queued"))
//		Expect(job.Id).To(Equal("test-job"))
//
//		job, err = mediamachine.SummaryMP4(
//			mediamachine.APIKey("SECRET_API_KEY"),
//			mediamachine.Webhooks("http://mytestserver/success", "http://mytestserver/failure"),
//			mediamachine.AzureInput("testbucket", "/foo/bar/testvideo.mp4", azureCreds),
//			mediamachine.AzureOutput("testbucket", "/foo/bar/testvideo.jpg", azureCreds),
//		)
//		Expect(err).ToNot(HaveOccurred())
//		Expect(job.FetchStatus).To(Equal("queued"))
//		Expect(job.Id).To(Equal("test-job"))
//
//		job, err = mediamachine.SummaryMP4(
//			mediamachine.APIKey("SECRET_API_KEY"),
//			mediamachine.Webhooks("http://mytestserver/success", "http://mytestserver/failure"),
//			mediamachine.GCPInput("testbucket", "/foo/bar/testvideo.mp4", gcpCreds),
//			mediamachine.GCPOutput("testbucket", "/foo/bar/testvideo.jpg", gcpCreds),
//		)
//		Expect(err).ToNot(HaveOccurred())
//		Expect(job.FetchStatus).To(Equal("queued"))
//		Expect(job.Id).To(Equal("test-job"))
//	})
//})
