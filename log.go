package dalog

import (
	"os"
	"strings"
)

// Context to associate to a logger
type Context struct {
	Key   string
	Value string
}

// WithID creates an ID context
func WithID(id string) Context {
	return Context{Key: "ID", Value: id}
}

// WithHostname creates and hostname context
func WithHostname() Context {
	if name, err := os.Hostname(); err == nil {
		return Context{Key: "Hostname", Value: name}
	}
	return Context{}
}

// Log is the logger
type Log interface {
	Info(a ...interface{})
	Warn(a ...interface{})
	Error(a ...interface{})
	Debug(a ...interface{})
	Infof(format string, a ...interface{})
	Warnf(format string, a ...interface{})
	Errorf(format string, a ...interface{})
	Debugf(format string, a ...interface{})
	WithContext(contexts ...Context) Log
}

// NoContext creates a logger without any configured context.
func NoContext() Log {
	debug := false
	if "TRUE" == strings.ToUpper(os.Getenv("DALOG_DEBUG")) {
		debug = true
	}

	if "ZAP" == os.Getenv("DALOG_LOGGER") {
		return zapLog{contexts: []Context{}, debugMode: debug}
	}
	return goLog{contexts: []Context{}, debugMode: debug}
}

// WithContext creates a logger with a context
func WithContext(contexts ...Context) Log {
	logger := NoContext()
	logger.WithContext(contexts...)
	return logger
}
