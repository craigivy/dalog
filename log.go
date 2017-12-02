package dalog

import (
	"os"
)

// Context to associate to a logger
type Context struct {
	Key   string
	Value string
}

// WithID creates and ID context
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
	Infof(format string, a ...interface{})
	Warnf(format string, a ...interface{})
	Errorf(format string, a ...interface{})
}

// WithContext creates a logger with a context
func WithContext(contexts ...Context) Log {
	if "ZAP" == os.Getenv("DALOG") {
		return zapLog{contexts: contexts}

	}
	return goLog{contexts: contexts}
}
