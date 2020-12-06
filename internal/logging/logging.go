package logging

import (
	"io/ioutil"
	"log"
	"os"
	"sync"
)

// Logger a collection of `*log.Logger` with different levels.
type Logger struct {
	sync.Mutex
	Info, Warning, Err, Debug, Trace *log.Logger
}

var defaultLogger *Logger
var once sync.Once

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

// DefaultLogger returns a singleton logger to be used across the app
func DefaultLogger() *Logger {
	once.Do(func() {
		defaultLogger = NewLogger(0, 0, 0, 0, 0, false)
	})
	return defaultLogger
}

// ConfigureLogger set the flags and ouptut stream of the logger.
func (l *Logger) ConfigureLogger(tf, ef, wf, nf, df int, debug bool) {
	l.Lock()
	defer l.Unlock()
	l.Info.SetFlags(nf)
	l.Err.SetFlags(ef)
	l.Warning.SetFlags(wf)
	l.Trace.SetFlags(tf)
	l.Debug.SetFlags(df)
	if debug {
		l.Warning.SetOutput(os.Stderr)
		l.Trace.SetOutput(os.Stderr)
		l.Debug.SetOutput(os.Stderr)
	} else {
		l.Warning.SetOutput(ioutil.Discard)
		l.Trace.SetOutput(ioutil.Discard)
		l.Debug.SetOutput(ioutil.Discard)
	}
}
