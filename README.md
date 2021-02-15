# Mediamachine

Go Mediamachine SDK for [MediaMachine](https://mediamachine.io)

#### Mediamachine APIs provide an easy way to create intelligent video Thumbnails and Summaries, as well as for common video transcoding tasks.

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

MediaMachine can work with input (and can store the output) videos from any of the following sources:

- URL (File Servers: For output, MediaMachine will `POST` to that URL)
- Amazon S3
- Google GCP
- Microsoft Azure buckets

Additionally, every processor accepts a success and failure endpoint, that we will call with the output of the process
once is done.

### Intelligent Thumbnail creation

MediaMachine SDK can generate beautiful Thumbnails for your videos. [Check the usage examples](examples/thumbnail)

### Intelligent Summary creation

MediaMachine SDK can generate an automatic Summary for your videos in GIF or MP4
format. [Check the usage examples](examples/summary)

### Video Transcoding

MediaMachine SDK can transcode your videos between different formats. [Check the usage examples](examples/transcode)

A Transcode processor can be configured to return different kind of outputs:

- Supported output containers: `MP4`, `WEBM`.
- Supported output encoders: `H265`, `H264`, `VP8`, `VP9`.
- Supported output bitrates: `1000kbps`, `2000kbps`, `4000kbps`.

## Contributing
