package log

import (
	"fmt"

	"github.com/micro/go-micro/v2/logger"
	"go.uber.org/zap"
)

type zaplog struct {
	zap *zap.Logger
}

func (l *zaplog) Init(...logger.Option) error {
	return nil
}

func (l *zaplog) Fields(fields map[string]interface{}) logger.Logger {
	nfields := make(map[string]interface{})
	for k, v := range fields {
		nfields[k] = v
	}

	//data := make([]interface{}, 0, len(nfields))
	data := make([]Field, 0, len(nfields))
	for k, v := range fields {
		data = append(data, Any(k, v))
	}

	zl := &zaplog{
		zap: l.zap.With(data...),
	}

	return zl
}

func (l *zaplog) Error(err error) logger.Logger {
	return l.Fields(map[string]interface{}{"error": err})
}

func (l *zaplog) Log(level logger.Level, args ...interface{}) {
	lvl := loggerToZapLevel(level)
	msg := fmt.Sprint(args...)
	switch lvl {
	case DebugLevel:
		l.zap.Debug(msg)
	case InfoLevel:
		l.zap.Info(msg)
	case WarnLevel:
		l.zap.Warn(msg)
	case ErrorLevel:
		l.zap.Error(msg)
	case FatalLevel:
		l.zap.Fatal(msg)
	}
}

func (l *zaplog) Logf(level logger.Level, format string, args ...interface{}) {
	lvl := loggerToZapLevel(level)
	msg := fmt.Sprintf(format, args...)
	switch lvl {
	case DebugLevel:
		l.zap.Debug(msg)
	case InfoLevel:
		l.zap.Info(msg)
	case WarnLevel:
		l.zap.Warn(msg)
	case ErrorLevel:
		l.zap.Error(msg)
	case FatalLevel:
		l.zap.Fatal(msg)
	}
}

func (l *zaplog) String() string {
	return "zap"
}

// Options 获取参数
func (l *zaplog) Options() logger.Options {
	return logger.Options{}
}

// NewLogger ...
func NewMicroLogger(logger ...*zap.Logger) *zaplog {
	l := &zaplog{}
	if len(logger) > 0 {
		l.zap = logger[0]
	} else {
		if defaultLogger.unsugar == nil {
			panic("default zap.Logger instance is nil")
		}
		l.zap = defaultLogger.unsugar
	}

	return l
}

// loggerToZapLevel 将micro中logger的level转化成zap的level类型
func loggerToZapLevel(level logger.Level) Level {
	switch level {
	case logger.TraceLevel, logger.DebugLevel:
		return DebugLevel
	case logger.InfoLevel:
		return InfoLevel
	case logger.WarnLevel:
		return WarnLevel
	case logger.ErrorLevel:
		return ErrorLevel
	case logger.FatalLevel:
		return FatalLevel
	default:
		return InfoLevel
	}
}
