# Mediamachine

StackRock's Mediamachine SDK for Go.

Mediamachine APIs provide an easy way to create intelligent video Thumbnails and Summaries, as well as for common video transcoding tasks.

This is the client to connect to [Stackrock](https://stackrock.io)'s services for:

- Generate a thumbnail from a video
- Generate a summary from a video
- Transcode a video

## Installation

```
    go get github.com/stackrock/mediamachinego
```

## Usage

Every processor (thumbnail, transcode and summary) returns a Job object that you can use to query
the state of the Job.

Every processor can get their input from any of the following sources:

- URL
- Amazon S3
- Google GCP
- Microsoft Azure buckets

Also, every processor can store the output in any of the following:

- Amazon S3
- Google GCP
- Microsoft Azure buckets.
- URL (We `POST` to that URL when the output is ready).

Additionally, every processor accepts a success and failure endpoint, that we will call with the output of the process
once is done.

### Intelligent Thumbnail creation

To create a new Thumbnail process you can use the ThumbnailJob struct [TODO Add link to documentation in the godoc].

#### Examples

Creates a new thumbnail processor job which consumes data from S3 and put the thumbnail on S3.

```golang
package main

import (
	"github.com/stackrock/mediamachinego"
	"os"
)

var STACKROCK_API_KEY = os.Getenv("STACKROCK_API_KEY")
var BUCKET = os.Getenv("BUCKET")
var INPUT_KEY = os.Getenv("INPUT_KEY")
var OUTPUT_KEY = os.Getenv("OUTPUT_KEY_THUMBNAIL")
var AWS_REGION = os.Getenv("AWS_REGION")
var AWS_ACCESS_KEY_ID = os.Getenv("AWS_ACCESS_KEY_ID")
var AWS_SECRET_ACCESS_KEY = os.Getenv("AWS_SECRET_ACCESS_KEY")

inputFile := NewS3BlobWithDefaults()
                .Bucket(BUCKET)
                .Key(INPUT_KEY)
                .AWSCredentials(AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, AWS_REGION)
outputFile := NewS3BlobWithDefaults()
                .Bucket(BUCKET)
                .Key(OUTPUT_KEY)
                .AWSCredentials(AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, AWS_REGION)

job, err := NewThumbnailJobWithDefaults()
                .ApiKey(STACKROCK_API_KEY)
                .From(inputFile)
                .To(outputFile)
                .Width(150)
                .WatermarkFromText("stackrock.io")
                .Execute()

statut, err = job.Status()
```

Creates a new thumbnail processor job which consumes data from Google and put the thumbnail on Azures, adding a text watermark to the output and reports the success or failure of the process on two different endpoints.

```golang
package main

import (
    "github.com/stackrock/mediamachinego"
    "os"
)

var STACKROCK_API_KEY = os.Getenv("STACKROCK_API_KEY")
var BUCKET = os.Getenv("BUCKET")
var INPUT_KEY = os.Getenv("INPUT_KEY")
var OUTPUT_KEY = os.Getenv("OUTPUT_KEY_THUMBNAIL")
var GCP_JSON = os.Getenv("GCP_JSON")
var AZURE_ACCOUNT_NAME = os.Getenv("AZURE_ACCOUNT_NAME")
var AZURE_ACCOUNT_KEY = os.Getenv("AZURE_ACCOUNT_KEY")

inputFile := NewGCPBlobWithDefaults()
    .Bucket(BUCKET)
    .Key(INPUT_KEY)
    .GCPCredentials(GCP_JSON)
outputFile := NewAzureBlobWithDefaults()
    .Bucket(BUCKET)
    .Key(OUTPUT_KEY)
    .AzureCredentials(AZURE_ACCOUNT_NAME, AZURE_ACCOUNT_KEY)

job, err := NewThumbnailJobWithDefaults()
    .ApiKey(STACKROCK_API_KEY)
    .From(inputFile)
    .To(outputFile)
    .Width(150)
    .Webhooks("stackrock.io/process/success", "stackrock.io/process/failure")
    .WatermarkFromText("stackrock.io")
    .Execute()

statut, err = job.Status()
```

### Intelligent Summary creation

Creates a new Intelligent Summary in GIF or MP4 format.

#### Gif

To create a new Intelligent Summary process that output Gif you can use
SummaryJob with the `Type` set to `SUMMARY_GIF`.

##### Examples
Creates a new summary Gif processor job which consumes data from S3 and put the summary output on S3.

```golang
package main

import (
	"github.com/stackrock/mediamachinego"
	"os"
)

var STACKROCK_API_KEY = os.Getenv("STACKROCK_API_KEY")
var BUCKET = os.Getenv("BUCKET")
var INPUT_KEY = os.Getenv("INPUT_KEY")
var OUTPUT_KEY = os.Getenv("OUTPUT_KEY_SUMMARY_GIF")
var AWS_REGION = os.Getenv("AWS_REGION")
var AWS_ACCESS_KEY_ID = os.Getenv("AWS_ACCESS_KEY_ID")
var AWS_SECRET_ACCESS_KEY = os.Getenv("AWS_SECRET_ACCESS_KEY")

inputFile := NewS3BlobWithDefaults()
                .Bucket(BUCKET)
                .Key(INPUT_KEY)
                .AWSCredentials(AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, AWS_REGION)

outputFile := NewS3BlobWithDefaults()
                .Bucket(BUCKET)
                .Key(OUTPUT_KEY)
                .AWSCredentials(AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, AWS_REGION)

job, err := NewSummaryJobWithDefaults()
            .ApiKey(STACKROCK_API_KEY)
            .Type(SUMMARY_GIF)
            .From(inputFile)
            .To(outputFile)
            .Width(150)
            .WatermarkFromText("stackrock.io")
.Execute()

status, err = job.Status()
```

#### MP4

To create a new Intelligent Summary process that output MP4 you can use
SummaryJob with the `Type` set to `SUMMARY_MP4`.

#### Examples
Creates a new summary MP4 processor job which consumes data from S3 and put the summary output on S3.

```golang
package main

import (
	"github.com/stackrock/mediamachine"
	"os"
)

var STACKROCK_API_KEY = os.Getenv("STACKROCK_API_KEY")
var BUCKET = os.Getenv("BUCKET")
var INPUT_KEY = os.Getenv("INPUT_KEY")
var OUTPUT_KEY = os.Getenv("OUTPUT_KEY_SUMMARY_MP4")
var AWS_REGION = os.Getenv("AWS_REGION")
var AWS_ACCESS_KEY_ID = os.Getenv("AWS_ACCESS_KEY_ID")
var AWS_SECRET_ACCESS_KEY = os.Getenv("AWS_SECRET_ACCESS_KEY")

inputFile := NewS3BlobWithDefaults()
                .Bucket(BUCKET)
                .Key(INPUT_KEY)
                .AWSCredentials(AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, AWS_REGION)
outputFile := NewS3BlobWithDefaults()
                .Bucket(BUCKET)
                .Key(OUTPUT_KEY)
                .AWSCredentials(AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, AWS_REGION)

job, err := NewSummaryJobWithDefaults()
                .ApiKey(STACKROCK_API_KEY)
                .Type(SUMMARY_MP4)
                .From(inputFile)
                .To(outputFile)
                .Width(150)
                .WatermarkFromText("stackrock.io")
                .Execute()

status, err = job.Status()
```

Creates a new summary processor job of type MP4 which consumes 
data from Google and put the summary on Azures, adding a text watermark 
to the output and reports the success or failure of the process on 
two different endpoints.
HERE
```golang
package main
import (
	"github.com/stackrock/mediamachinego"
	"os"
)

var STACKROCK_API_KEY = os.Getenv("STACKROCK_API_KEY")
var BUCKET = os.Getenv("BUCKET")
var INPUT_KEY = os.Getenv("INPUT_KEY")
var OUTPUT_KEY = os.Getenv("OUTPUT_KEY_SUMMARY_MP4")
var GCP_JSON = os.Getenv("GCP_JSON")
var AZURE_ACCOUNT_NAME = os.Getenv("AZURE_ACCOUNT_NAME")
var AZURE_ACCOUNT_KEY = os.Getenv("AZURE_ACCOUNT_KEY")

inputFile := NewGCPBlobWithDefaults()
               .Bucket(BUCKET)
               .Key(INPUT_KEY)
               .GCPCredentials(GCP_JSON)

outputFile := NewAzureBlobWithDefaults()
               .Bucket(BUCKET)
               .Key(OUTPUT_KEY)
               .AzureCredentials(AZURE_ACCOUNT_NAME, AZURE_ACCOUNT_KEY)


job, err := NewSummaryJobWithDefaults()
               .ApiKey(STACKROCK_API_KEY)
               .Type(SUMMARY_MP4)
               .From(inputFile)
               .To(outputFile)
               .Width(150)
               .Webhooks("stackrock.io/process/success", "stackrock.io/process/failure")
               .WatermarkFromText("stackrock.io")
               .Execute()

status, err = job.Status()
```

### Video Transcoding

Creates a new transcode processor job that consumes data 
from a provided input and put the result in the provided output.

A Transcode processor can be configured to return different kind of outputs:

- Supported output containers: `MP4`, `WEBM`.
- Supported output bitrate: `8000kbps`, `4000kbps`, `1000kbps`.
- Supported output encoder: `H265`, `H264`, `VP8`.
- Supported output video definition: `FULL_HD`, `HD`, `SD`.

#### Examples

Creates a new transcode processor job which consumes data from S3 
and put the transcoded video on S3.

```golang
package main

import (
	"github.com/stackrock/mediamachinego"
	"os"
)
var STACKROCK_API_KEY = os.Getenv("STACKROCK_API_KEY")
var BUCKET = os.Getenv("BUCKET")
var INPUT_KEY = os.Getenv("TRANSCODE_INPUT_KEY")
var OUTPUT_KEY = os.Getenv("OUTPUT_KEY_TRANSCODE")
var AWS_REGION = os.Getenv("AWS_REGION")
var AWS_ACCESS_KEY_ID = os.Getenv("AWS_ACCESS_KEY_ID")
var AWS_SECRET_ACCESS_KEY = os.Getenv("AWS_SECRET_ACCESS_KEY")

inputFile := NewS3BlobWithDefaults()
                .Bucket(BUCKET)
                .Key(INPUT_KEY)
                .AWSCredentials(AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, AWS_REGION)
outputFile := NewS3BlobWithDefaults()
                .Bucket(BUCKET)
                .Key(OUTPUT_KEY)
                .AWSCredentials(AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, AWS_REGION)

// This options object is the one that configures the processor.
transcodeOpts := NewTranscodeOptsWithDefaults()
                    .Encoder(ENCODER_H264)
                    .BitrateKbps(BITRATE_FOUR_MEGAKBPS)
                    .Container(CONTAINER_MP4)
                    .VideoSize(VIDEOSIZE_HD)

job, err := NewTranscodeJobWithDefaults()
                .ApiKey(STACKROCK_API_KEY)
                .From(inputFile)
                .To(outputFile)
                .Width(150)
                .WatermarkFromText("stackrock.io")
                .Opts(transcodeOpts)
                .Execute()

status, err = job.Status()
```

Creates a new transcode processor job which consumes data from Google 
and put the transcoded video on Azure, adding a text watermark to the output 
and reports the success or failure of the process on two different endpoints.

```golang
package main

import (
	"github.com/stackrock/mediamachinego"
	"os"
)

var STACKROCK_API_KEY = os.Getenv("STACKROCK_API_KEY")
var BUCKET = os.Getenv("BUCKET")
var INPUT_KEY = os.Getenv("INPUT_KEY")
var OUTPUT_KEY = os.Getenv("OUTPUT_KEY_SUMMARY_MP4")
var GCP_JSON = os.Getenv("GCP_JSON")
var AZURE_ACCOUNT_NAME = os.Getenv("AZURE_ACCOUNT_NAME")
var AZURE_ACCOUNT_KEY = os.Getenv("AZURE_ACCOUNT_KEY")

inputFile := NewGCPBlobWithDefaults()
                .Bucket(BUCKET)
                .Key(INPUT_KEY)
                .GCPCredentials(GCP_JSON)

outputFile := NewAzureBlobWithDefaults()
                .Bucket(BUCKET)
                .Key(OUTPUT_KEY)
                .AzureCredentials(AZURE_ACCOUNT_NAME, AZURE_ACCOUNT_KEY)

// This options object is the one that configures the processor.
transcodeOpts := NewTranscodeOptsWithDefaults()
                    .Encoder(ENCODER_H264)
                    .BitrateKbps(BITRATE_FOUR_MEGAKBPS)
                    .Container(CONTAINER_MP4)
                    .VideoSize(VIDEOSIZE_HD)

job, err := NewTranscodeJobWithDefaults()
                .ApiKey(STACKROCK_API_KEY)
                .From(inputFile)
                .To(outputFile)
                .Width(150)
                .WatermarkFromText("stackrock.io")
                .Webhooks("stackrock.io/process/success", "stackrock.io/process/failure")
                .Opts(transcodeOpts)
                .Execute()

status, err = job.Status()
```

## Contributing