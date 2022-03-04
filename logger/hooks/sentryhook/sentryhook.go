package sentryhook

import (
	"net/http"

	"github.com/getsentry/sentry-go"
	log "github.com/sirupsen/logrus"
)

var (
	severityMap = map[log.Level]sentry.Level{
		log.TraceLevel: sentry.LevelDebug,
		log.DebugLevel: sentry.LevelDebug,
		log.InfoLevel:  sentry.LevelInfo,
		log.WarnLevel:  sentry.LevelWarning,
		log.ErrorLevel: sentry.LevelError,
		log.FatalLevel: sentry.LevelFatal,
		log.PanicLevel: sentry.LevelFatal,
	}
)

// Hook to report events to sentry
type Hook struct {
	levels []log.Level
}

// New creates hook to be added to an instance of logger.
func New(levels []log.Level) *Hook {
	return &Hook{
		levels: levels,
	}
}

// Fire takes the entry that the hook is fired for.
// This method never returns an error
func (hook *Hook) Fire(entry *log.Entry) error {
	// Clone current hub from global namespace
	hub := sentry.CurrentHub().Clone()
	// Get this hub's client and scope
	client, scope := hub.Client(), hub.Scope()
	if client == nil || scope == nil {
		return nil
	}
	// Enrich scope with context
	if ctx := entry.Context; ctx != nil {
		// Check if http.Request exists in context
		if r := ctx.Value(sentry.RequestContextKey); r != nil {
			if r, ok := r.(*http.Request); ok {
				// Add it to current scope
				scope.SetRequest(r)
			}
		}
	}
	// Create a new event for this fired entry
	event := sentry.NewEvent()
	// Enrich the event
	event.Level = severityMap[entry.Level]
	// Title of the event
	event.Message = entry.Message
	// Current thead's stacktrace
	event.Threads = []sentry.Thread{{
		Stacktrace: sentry.NewStacktrace(),
		Crashed:    false,
		Current:    true,
	}}
	event.Extra = entry.Data
	// Enrich event with data fields in logrus.Entry
	hint := &sentry.EventHint{}
	hint.Data = entry.Data
	// Enrich exception if error field exists in logrus.Entry
	if err, ok := entry.Data[log.ErrorKey]; ok {
		if err, ok := err.(error); ok {
			hint.OriginalException = err
		}
	}
	// Send event to sentry
	client.CaptureEvent(event, hint, scope)
	return nil
}

// Levels returns a slice of `log.Levels` the hook is fired for.
func (hook *Hook) Levels() []log.Level {
	return hook.levels
}
