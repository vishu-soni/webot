package sentry

import (
	"time"

	log "github.com/sirupsen/logrus"

	"webot/config"

	"github.com/getsentry/sentry-go"
)

// Init initializes sentry SDK
func Init() error {
	cfg := config.Sentry()
	if cfg.Enabled {
		err := sentry.Init(sentry.ClientOptions{
			Dsn:         cfg.DSN,
			Environment: config.Env().Env,
			SampleRate:  1.0, // Process all events
		})
		log.Info("sentry: sentry enabled")
		return err
	}
	log.Info("sentry: sentry disabled")
	return nil
}

// Flush waits until the underlying Transport sends any buffered events to the Sentry server, blocking for at most the given timeout. It returns false if the timeout was reached. In that case, some events may not have been sent.
// Flush should be called before terminating the program to avoid unintentionally dropping events.
func Flush(timeout time.Duration) bool {
	return sentry.Flush(timeout)
}
