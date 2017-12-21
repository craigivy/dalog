package dalog

import (
	"fmt"
	"log"
)

type goLog struct {
	contexts  []Context
	debugMode bool
}

func (golog goLog) Debug(a ...interface{}) {
	if !golog.debugMode {
		return
	}
	msg := fmt.Sprint(a...)
	msg = appendContexts(msg, golog.contexts)
	msg = prependLevel("DEBUG", msg)
	log.Println(msg)
}

func (golog goLog) Info(a ...interface{}) {
	msg := fmt.Sprint(a...)
	msg = appendContexts(msg, golog.contexts)
	msg = prependLevel("INFO", msg)
	log.Println(msg)
}

func (golog goLog) Warn(a ...interface{}) {
	msg := fmt.Sprint(a...)
	msg = appendContexts(msg, golog.contexts)
	msg = prependLevel("WARN", msg)
	log.Println(msg)
}

func (golog goLog) Error(a ...interface{}) {
	msg := fmt.Sprint(a...)
	msg = appendContexts(msg, golog.contexts)
	msg = prependLevel("ERROR", msg)
	log.Println(msg)
}

func (golog goLog) Debugf(format string, a ...interface{}) {
	if !golog.debugMode {
		return
	}
	msg := fmt.Sprintf(format, a...)
	msg = appendContexts(msg, golog.contexts)
	msg = prependLevel("DEBUG", msg)
	log.Println(msg)
}

func (golog goLog) Infof(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	msg = appendContexts(msg, golog.contexts)
	msg = prependLevel("INFO", msg)
	log.Println(msg)
}

func (golog goLog) Warnf(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	msg = appendContexts(msg, golog.contexts)
	msg = prependLevel("WARN", msg)
	log.Println(msg)
}

func (golog goLog) Errorf(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	msg = appendContexts(msg, golog.contexts)
	msg = prependLevel("ERROR", msg)
	log.Println(msg)
}

// Return a new logger with the combined contexts of the old logger and the
// provided contexts.
func (golog goLog) WithContext(contexts ...Context) Log {
	golog.contexts = append(golog.contexts, contexts...)
	return golog
}

func appendContexts(msg string, contexts []Context) string {
	for _, context := range contexts {
		msg = fmt.Sprintf("%s, %s=%s", msg, context.Key, context.Value)
	}
	return msg
}

func prependLevel(level string, msg string) string {
	return fmt.Sprintf("%s %s", level, msg)
}
