package dalog

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
)

var onceSync sync.Once
var zapLogInst *zap.Logger

func zapInstance(debugMode bool) *zap.Logger {
	onceSync.Do(func() {
		cfg := zap.NewProductionConfig()
		if debugMode {
			cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		}
		cfg.Development = false
		cfg.DisableCaller = true
		cfg.DisableStacktrace = true
		logger, _ := cfg.Build()
		//		logger, _ := cfg.Build(zap.AddCaller(), zap.AddCallerSkip(2))

		//		logger, _ := zap.NewProduction(zap.AddCaller(), zap.AddCallerSkip(2))
		defer logger.Sync()
		zapLogInst = logger

	})
	return zapLogInst
}

type zapLog struct {
	contexts     []Context
	debugMode    bool
	includeStack bool
	debugContext string
}

func (zl zapLog) Debug(a ...interface{}) {
	if !zl.debugMode {
		return
	}
	msg := fmt.Sprint(a...)
	zl.debug(msg)
}

func (zl zapLog) Debugf(format string, a ...interface{}) {
	if !zl.debugMode {
		return
	}
	msg := fmt.Sprintf(format, a...)
	zl.debug(msg)
}

func (zl zapLog) debug(msg string) {

	if debugContext == "" {
		zapInstance(zl.debugMode).Debug(msg, convert(zl.contexts)...)
	}

	// add the debug context to contexts for the debug line
	debugContext := Context{Key: debugContext, Value: zl.debugContext}
	contexts := append(zl.contexts, debugContext)
	zapInstance(zl.debugMode).Debug(msg, convert(contexts)...)
}

func (zl zapLog) Info(a ...interface{}) {
	msg := fmt.Sprint(a...)
	zapInstance(zl.debugMode).Info(msg, convert(zl.contexts)...)
}

func (zl zapLog) Warn(a ...interface{}) {
	msg := fmt.Sprint(a...)
	zapInstance(zl.debugMode).Warn(msg, convert(zl.contexts)...)
}

func (zl zapLog) Error(err error) {
	fields := convert(zl.contexts)
	stackString, stackExists := stackString(err)
	if zl.includeStack && stackExists {
		field := zap.String("stack", fmt.Sprintf("%s", stackString))
		fields = append(fields, field)
	}

	zapInstance(zl.debugMode).Error(err.Error(), fields...)
}

func (zl zapLog) Infof(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	zapInstance(zl.debugMode).Info(msg, convert(zl.contexts)...)
}

func (zl zapLog) Warnf(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	zapInstance(zl.debugMode).Warn(msg, convert(zl.contexts)...)
}

func (zl zapLog) Errorf(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	zapInstance(zl.debugMode).Error(msg, convert(zl.contexts)...)
}

// Return a new logger with the combined contexts of the old logger and the
// provided contexts.
func (zl zapLog) WithContext(contexts ...Context) Log {

	// check for the specialized debug context
	for i, context := range contexts {
		if context.Key == debugContext {
			zl.debugContext = context.Value
			contexts = append(contexts[:i], contexts[i+1:]...)
			break
		}
	}
	zl.contexts = append(zl.contexts, contexts...)
	return zl
}

func convert(contexts []Context) []zapcore.Field {
	var fields = make([]zapcore.Field, len(contexts))
	for i, context := range contexts {
		fields[i] = zap.String(context.Key, context.Value)
	}
	return fields
}
