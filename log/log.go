package log

import (
	"os"
	"github.com/rs/zerolog"
)

// Logger is the global logger instance.
var Logger zerolog.Logger

func init() {
	Logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}

// SetLogLevel sets the global log level.
func SetLogLevel(level zerolog.Level) {
	zerolog.SetGlobalLevel(level)
}
