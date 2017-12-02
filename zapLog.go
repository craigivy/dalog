package dalog

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
)

var onceSync sync.Once
var zapLogInst *zap.Logger

func zapInstance() *zap.Logger {
	onceSync.Do(func() {
		logger, _ := zap.NewProduction()
		defer logger.Sync()
		zapLogInst = logger
	})
	return zapLogInst
}

type zapLog struct {
	contexts []Context
}

func (zl zapLog) Infof(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	zapInstance().Info(msg, convert(zl.contexts)...)
}

func (zl zapLog) Warnf(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	zapInstance().Warn(msg, convert(zl.contexts)...)
}

func (zl zapLog) Errorf(format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	zapInstance().Error(msg, convert(zl.contexts)...)
}

func convert(contexts []Context) []zapcore.Field {
	var fields = make([]zapcore.Field, len(contexts))
	for i, context := range contexts {
		fields[i] = zap.String(context.Key, context.Value)
	}
	return fields
}
