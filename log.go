package dalog

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"strings"
)

// Context to associate to a logger
type Context struct {
	Key   string
	Value string
}

const (
	debugContext = "debugContext"
)

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

// WithDebugContext is a specialized context to allow debug lines to be related and annotated
func WithDebugContext(context string) Context {
	return Context{debugContext, context}
}

// Log is the logger
type Log interface {
	Info(a ...interface{})
	Warn(a ...interface{})
	Error(err error)
	Debug(a ...interface{})
	Infof(format string, a ...interface{})
	Warnf(format string, a ...interface{})
	Debugf(format string, a ...interface{})
	WithContext(contexts ...Context) Log
}

// NoContext creates a logger without any configured context.
func NoContext() Log {
	debug := false
	if "TRUE" == strings.ToUpper(os.Getenv("DALOG_DEBUG")) {
		debug = true
	}

	includeStack := false
	if "TRUE" == strings.ToUpper(os.Getenv("DALOG_STACK")) {
		includeStack = true
	}

	if "ZAP" == os.Getenv("DALOG_LOGGER") {
		return zapLog{contexts: []Context{}, debugMode: debug, includeStack: includeStack}
	}
	return goLog{contexts: []Context{}, debugMode: debug, includeStack: includeStack}
}

// WithContext creates a logger with a context
func WithContext(contexts ...Context) Log {
	logger := NoContext()
	logger = logger.WithContext(contexts...)
	return logger
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func stackString(err error) (string, bool) {

	str, ok := err.(stackTracer)
	if ok {
		st := str.StackTrace()
		return stringify(st), ok
	}

	cause, ok := errors.Cause(err).(stackTracer)
	if !ok {
		return "", ok
	}
	st := cause.StackTrace()
	return fmt.Sprintf("%+v", st[:]), ok
}

func stringify(stackTrace errors.StackTrace) string {
	return fmt.Sprintf("%+v", stackTrace[:])
}
