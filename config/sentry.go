package config

// SentryConfig holds configuration for sentry
type SentryConfig struct {
	DSN     string
	Enabled bool
}