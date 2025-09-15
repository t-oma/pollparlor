// Package logger provides a logger for the application
package logger

import (
	"io"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

// Options represents the logger options
type Options struct {
	Level  string // "debug","info","warn","error"
	Format string // "json" або "console"
	Out    io.Writer
}

// New creates a new logger
func New(opts Options) zerolog.Logger {
	if opts.Out == nil {
		opts.Out = os.Stdout
	}
	level := parseLevel(opts.Level)

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(level)

	out := opts.Out
	if strings.ToLower(opts.Format) == "console" {
		out = zerolog.ConsoleWriter{Out: opts.Out, TimeFormat: time.RFC3339}
	}

	return zerolog.New(out).With().Caller().Timestamp().Logger()
}

func parseLevel(s string) zerolog.Level {
	switch strings.ToLower(s) {
	case "debug":
		return zerolog.DebugLevel
	case "warn", "warning":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	default:
		return zerolog.InfoLevel
	}
}
