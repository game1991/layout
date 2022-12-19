package log

import (
	"go.uber.org/zap/zapcore"
)

var defaultLogger *Logger

// DefaultLogger 获取defaultLogger实例
func DefaultLogger() *Logger {
	return defaultLogger
}

// Debugz 使用默认的unsugar(*zap.Logger)打印debug级别的日志
func Debugz(msg string, fields ...Field) {
	defaultLogger.Debugz(msg, fields...)
}

// Debug 使用默认的sugar(*zap.SugaredLogger)打印debug级别的日志
func Debug(msg string, args ...interface{}) {
	defaultLogger.Debug(msg, args...)
}

// Debugf 使用默认的sugar(*zap.SugaredLogger)格式化打印debug级别的日志
func Debugf(format string, args ...interface{}) {
	defaultLogger.Debugf(format, args...)
}

// Infoz 使用默认的unsugar(*zap.Logger)打印info级别的日志
func Infoz(msg string, fields ...Field) {
	defaultLogger.Infoz(msg, fields...)
}

// Info 使用默认的sugar(*zap.SugaredLogger)打印info级别的日志
func Info(msg string, args ...interface{}) {
	defaultLogger.Info(msg, args...)
}

// Infof 使用默认的sugar(*zap.SugaredLogger)格式化打印info级别的日志
func Infof(format string, args ...interface{}) {
	defaultLogger.Infof(format, args...)
}

// Warnz 使用默认的unsugar(*zap.Logger)打印warn级别的日志
func Warnz(msg string, fields ...Field) {
	defaultLogger.Warnz(msg, fields...)
}

// Warn 使用默认的sugar(*zap.SugaredLogger)打印warn级别的日志
func Warn(msg string, args ...interface{}) {
	defaultLogger.Warn(msg, args...)
}

// Warnf 使用默认的sugar(*zap.SugaredLogger)格式化打印warn级别的日志
func Warnf(format string, args ...interface{}) {
	defaultLogger.Warnf(format, args...)
}

// Errorz 使用默认的unsugar(*zap.Logger)打印error级别的日志
func Errorz(msg string, fields ...Field) {
	defaultLogger.Errorz(msg, fields...)
}

// Error 使用默认的sugar(*zap.SugaredLogger)打印error级别的日志
func Error(msg string, args ...interface{}) {
	defaultLogger.Error(msg, args...)
}

// Errorf 使用默认的sugar(*zap.SugaredLogger)格式化打印error级别的日志
func Errorf(format string, args ...interface{}) {
	defaultLogger.Errorf(format, args...)
}

// Panicz 使用默认的unsugar(*zap.Logger)打印panic级别的日志,并发生panic
func Panicz(msg string, fields ...Field) {
	defaultLogger.Panicz(msg, fields...)
}

// Panic 使用默认的sugar(*zap.SugaredLogger)打印panic级别的日志,并发生panic
func Panic(msg string, args ...interface{}) {
	defaultLogger.Panic(msg, args...)
}

// Panicf 使用默认的sugar(*zap.SugaredLogger)格式化打印panic级别的日志,并发生panic
func Panicf(format string, args ...interface{}) {
	defaultLogger.Panicf(format, args...)
}

// Fatalz 使用默认的unsugar(*zap.Logger)打印fatal级别的日志,并调用 os.Exit(1) 退出
func Fatalz(msg string, fields ...Field) {
	defaultLogger.Fatalz(msg, fields...)
}

// Fatal 使用默认的sugar(*zap.SugaredLogger)打印fatal级别的日志,并调用 os.Exit(1) 退出
func Fatal(msg string, args ...interface{}) {
	defaultLogger.Fatal(msg, args...)
}

// Fatalf 使用默认的sugar(*zap.SugaredLogger)格式化打印fatal级别的日志,并调用 os.Exit(1) 退出
func Fatalf(format string, args ...interface{}) {
	defaultLogger.Fatalf(format, args...)
}

// FieldErr 打印error日志
func FieldErr(err error) Field {
	return defaultLogger.FieldErr(err)
}

// SetDefaultFields 设置默认的fields
func SetDefaultFields(fields ...zapcore.Field) *Logger {
	return defaultLogger.SetDefaultFields(fields...)
}

// SetDefaultDir 设置默认的日志文件路径
func SetDefaultDir(dir string) *Logger {
	return defaultLogger.SetDefaultDir(dir)
}

// SetLevel 设置日志等级
func SetLevel(level Level) *Logger {
	return defaultLogger.SetLevel(level)
}

// SetLevelString 设置日志等级,
func SetLevelString(level string) *Logger {
	lv := parseLevel(level)
	return defaultLogger.SetLevel(lv)
}
