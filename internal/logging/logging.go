package logging

import (
	"io/ioutil"
	"log"
	"os"
)

// Logger a collection of `*log.Logger` with different levels.
type Logger struct {
	Info, Warning, Err, Debug, Trace *log.Logger
}

// NewLogger a factory for Logger struct.
func NewLogger(tf, ef, wf, nf, df int, debug bool) *Logger {
	if debug {
		return &Logger{
			Info:    log.New(os.Stdout, "INFO: ", nf),
			Err:     log.New(os.Stderr, "ERROR: ", ef),
			Warning: log.New(os.Stderr, "WARNING: ", wf),
			Trace:   log.New(os.Stderr, "TRACE: ", tf),
			Debug:   log.New(os.Stderr, "DEBUG: ", df),
		}
	}
	return &Logger{
		Info:    log.New(os.Stdout, "INFO: ", nf),
		Err:     log.New(os.Stderr, "ERROR: ", ef),
		Warning: log.New(ioutil.Discard, "WARNING: ", wf),
		Trace:   log.New(ioutil.Discard, "TRACE: ", tf),
		Debug:   log.New(ioutil.Discard, "DEBUG: ", df),
	}
}
