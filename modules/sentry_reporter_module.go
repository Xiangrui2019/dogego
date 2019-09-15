package modules

import (
	"os"

	"github.com/getsentry/raven-go"
)

var SentryClient *raven.Client

func InitSentryClient() {
	SentryClient, _ = raven.New(os.Getenv("SENTRY_DSN"))
}
