package gojsonq

import (
	"errors"
)

// option describes type for providing configuration options to JSONQ
type option struct {
	debug   bool
	logger  Logger
	decoder Decoder
}

// OptionFunc represents a contract for option func, it basically set options to jsonq instance options
type OptionFunc func(*JSONQ) error

// SetDebug enable/disable info log
func SetDebug(debug bool) OptionFunc {
	return func(j *JSONQ) error {
		j.option.debug = debug
		if debug {
			j.log("Debug mode on")
		}
		return nil
	}
}

// SetLogger set a user provided logger
func SetLogger(logger Logger) OptionFunc {
	return func(j *JSONQ) error {
		if logger == nil {
			return errors.New("logger can not be nil")
		}
		j.log("Enable custom logger")
		j.option.logger = logger
		j.log("Enable custom logger complete")
		return nil
	}
}

// SetDecoder take a custom decoder to decode JSON
func SetDecoder(u Decoder) OptionFunc {
	return func(j *JSONQ) error {
		if u == nil {
			return errors.New("decoder can not be nil")
		}
		j.log("Setup custom decoder")
		j.option.decoder = u
		j.log("Setup custom decoder complete")
		return nil
	}
}
