package gojsonq

import (
	"log"
)

// Logger provides logger contract to log operations.
type Logger interface {
	Printf(format string, v ...interface{})
}

// DefaultLogger provide default logger to write logs in stdout
type DefaultLogger struct {
}

// Printf wirte logs to stdout
func (*DefaultLogger) Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}
