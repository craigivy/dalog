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
		logger, _ := cfg.Build(zap.AddCaller(), zap.AddCallerSkip(2))

		//		logger, _ := zap.NewProduction(zap.AddCaller(), zap.AddCallerSkip(2))
		defer logger.Sync()
		zapLogInst = logger

	})
	return zapLogInst
}

type zapLog struct {
	contexts  []Context
	debugMode bool
}

func (zl zapLog) Debugf(format string, a ...interface{}) {
	if !zl.debugMode {
		return
	}
	msg := fmt.Sprintf(format, a...)
	zapInstance(zl.debugMode).Debug(msg, convert(zl.contexts)...)
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

func convert(contexts []Context) []zapcore.Field {
	var fields = make([]zapcore.Field, len(contexts))
	for i, context := range contexts {
		fields[i] = zap.String(context.Key, context.Value)
	}
	return fields
}
