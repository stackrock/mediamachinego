# Mediamachinego

Mediamachine SDK for Go.

Mediamachine APIs provide an easy way to create intelligent video Thumbnails and Summaries, as well as for common video transcoding tasks.

## Usage

### Intelligent Thumbnail creation
```go
// Using an S3 bucket URL for video input/output
awsCreds := mediamachine.AWSCreds{AccessKeyID: "testkey", SecretAccessKey: "testsecret", Region: "us-east-1"} 
job, err := mediamachine.Thumbnail(
    mediamachine.APIKey("SECRET_API_KEY"),
    mediamachine.Webhooks("http://mytestserver/success", "http://mytestserver/failure"),
    mediamachine.S3Input("myBucket", "/foo/bar/myVideo.mp4", awsCreds),
    mediamachine.S3Output("myBucket", "/foo/bar/myThumbnail.jpg", awsCreds),
)
```
(`AzureInput` and `GCPInput` and their output variants can also be used)


```go
// Using a custom URL for video input/output 
job, err := mediamachine.Thumbnail(
    mediamachine.APIKey("SECRET_API_KEY"),
    mediamachine.Webhooks("http://mytestserver/success", "http://mytestserver/failure"),
    mediamachine.InputURL("http://mytestserver/myvideo.mp4"),
    mediamachine.OutputURL("http://mytestserver/myvideo.jpg"),
)
```

### Intelligent Summary creation
```go
// Using an S3 bucket URL for video input/output
awsCreds := mediamachine.AWSCreds{AccessKeyID: "testkey", SecretAccessKey: "testsecret", Region: "us-east-1"} 
job, err := mediamachine.SummaryGIF(
    mediamachine.APIKey("SECRET_API_KEY"),
    mediamachine.Webhooks("http://mytestserver/success", "http://mytestserver/failure"),
    mediamachine.S3Input("myBucket", "/foo/bar/myVideo.mp4", awsCreds),
    mediamachine.S3Output("myBucket", "/foo/bar/mySummary.gif", awsCreds),
)

// Using an S3 bucket URL for video input/output
awsCreds := mediamachine.AWSCreds{AccessKeyID: "testkey", SecretAccessKey: "testsecret", Region: "us-east-1"} 
job, err := mediamachine.SummaryMP4(
    mediamachine.APIKey("SECRET_API_KEY"),
    mediamachine.Webhooks("http://mytestserver/success", "http://mytestserver/failure"),
    mediamachine.S3Input("myBucket", "/foo/bar/myVideo.mp4", awsCreds),
    mediamachine.S3Output("myBucket", "/foo/bar/mySummary.mp4", awsCreds),
)
```
(`AzureInput` and `GCPInput` and their output variants can also be used)


```go
// Using a custom URL for video input/output 
job, err := mediamachine.SummaryGIF(
    mediamachine.APIKey("SECRET_API_KEY"),
    mediamachine.Webhooks("http://mytestserver/success", "http://mytestserver/failure"),
    mediamachine.InputURL("http://mytestserver/myvideo.mp4"),
    mediamachine.OutputURL("http://mytestserver/myvideo.gif"),
)
```
