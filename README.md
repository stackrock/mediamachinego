# Mediamachine

Go Mediamachine SDK for [MediaMachine](https://www.mediamachine.io)

[![Go Reference](https://pkg.go.dev/badge/github.com/stackrock/mediamachinego.svg)](https://pkg.go.dev/github.com/stackrock/mediamachinego)
[![Go Report Card](https://goreportcard.com/badge/github.com/stackrock/mediamachinego)](https://goreportcard.com/report/github.com/stackrock/mediamachinego)

#### Mediamachine APIs provide an easy way to create intelligent video Thumbnails and Summaries, as well as for common video transcode tasks.

You can use this SDK for:

- Generating a thumbnail from videos
- Generating a summary from videos
- Transcode videos into different formats

## Installation

```
go get github.com/stackrock/mediamachinego
```

## Usage

Initialize the MediaMachine SDK with an API obtained from your [MediaMachine Account](https://www.mediamachine.io)

```golang
import "github.com/stackrock/mediamachinego/mediamachine"
mm := mediamachine.MediaMachine{APIKey: apiKey}
```

MediaMachine works with various video storage sources:

- URL (File Servers: For output, MediaMachine will `POST` to that URL)
- Amazon S3
- Google GCP
- Microsoft Azure buckets

The input/output sources can also be different! You can leverage MediaMachine API to also move video assets
between different storage locations as part of the processing steps.

Additionally, the API calls also accept success and failure endpoints - MediaMachine asynchronously calls them when the video is processed.

### Intelligent Thumbnail creation

MediaMachine SDK can generate beautiful Thumbnails for your videos by scanning the video for good quality frames. [Sample code](examples/thumbnail)

### Intelligent Summary creation

MediaMachine SDK can generate an automatic Summary for your videos in GIF or MP4 format. [Sample code](examples/summary)

### Video Transcoding

MediaMachine SDK can transcode your videos between different formats. [Sample code](examples/transcode)

A Transcode processor can be configured to return different kind of outputs:

- Supported output containers: `MP4`, `WEBM`.
- Supported output encoders: `H265`, `H264`, `VP8`, `VP9`.
- Supported output bitrates: `1000kbps`, `2000kbps`, `4000kbps`.

## Contributing

We welcome feedback and PRs and appreciate efforts to help us improve.

Our general guidance is:

- Check the issues/PR list to see if someone might have already answered your concern
- If you want to suggest changes, feel free to open an issue first and give details about the errors you might be seeing or suggested improvements
- If you're opening issues/PRs please try to give as much detailed information as possible so that we can help you right away

