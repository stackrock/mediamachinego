def main(ctx):
    return [
        pipeline("push"),
        pipeline("cron")
    ]

def pipeline(kind):
    return {
        "kind": "pipeline",
        "type": "docker",
        "name": "tracerBullets %s" % kind,
        "platform": {
            "arch": "arm64"
        },
        "concurrency": {
           "limit": 1
        },
        "trigger": trigger(kind),
        "steps": [
            {
                "name": "tracerBullets",
                "pull": "if-not-exists",
                "image": "arm64v8/golang:1.15-alpine",
                "environment": {
                    "STACKROCK_API_KEY": {
                        "from_secret": "STACKROCK_API_KEY"
                    },
                    "BUCKET": {
                        "from_secret": "BUCKET"
                    },
                    "AWS_REGION": {
                        "from_secret": "AWS_REGION"
                    },
                    "AWS_ACCESS_KEY_ID": {
                        "from_secret": "AWS_ACCESS_KEY_ID"
                    },
                    "AWS_SECRET_ACCESS_KEY": {
                        "from_secret": "AWS_SECRET_ACCESS_KEY"
                    },
                    "INPUT_KEY": {
                        "from_secret": "INPUT_KEY"
                    },
					"TRANSCODE_INPUT_KEY": {
						"from_secret": "TRANSCODE_INPUT_KEY"
					},
					"OUTPUT_KEY_THUMBNAIL": {
						"from_secret": "THUMBNAIL_OUTPUT_KEY"
					},
					"OUTPUT_KEY_SUMMARY_GIF": {
						"from_secret": "SUMMARY_GIF_OUTPUT_KEY"
					},
					"OUTPUT_KEY_SUMMARY_MP4": {
						"from_secret": "SUMMARY_MP4_OUTPUT_KEY"
					},
					"OUTPUT_KEY_TRANSCODE":{
						"from_secret": "TRANSCODE_OUTPUT_KEY"
					}
                },
                "commands": [
                	"apk add gcc musl-dev",
                    "go get github.com/onsi/ginkgo/ginkgo",
                    "go get github.com/onsi/gomega/...",
                    'ginkgo -p -focus "tracer"'
                ]
            },
            {
                "name": "slack",
                "image": "plugins/slack",
                "pull": "if-not-exists",
                "settings": {
                    "webhook": {
                        "from_secret": "SLACK_WEBHOOK"
                    },
                    "channel": {
                        "from_secret": "SLACK_CHANNEL"
                    }
                },
                "when": {
                    "status": ["success", "failure"]
                },
                "depends_on": [
                    "tracerBullets",
                ]
            }
        ]
    }

def trigger(kind):
    if kind == "cron":
        return {
           "branch": ["master"],
           "event": ["cron"],
           "cron": ["hourly"]
        }
    else:
        return {
           "branch": ["master"],
        }
