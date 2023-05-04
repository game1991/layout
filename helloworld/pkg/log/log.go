package log

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Logger ...
type Logger struct {
	sugar       *zap.SugaredLogger
	unsugar     *zap.Logger
	atomicLevel *zap.AtomicLevel
	fields      []zapcore.Field
	rotateLog   *lumberjack.Logger
	writer      zapcore.WriteSyncer
	Options
}

// New 通过Option创建Logger实例
func New(options Options) *Logger {
	// 默认参数
	if options.Output == "" {
		options.Output = "stdout"
	}
	if options.Level == "" {
		options.Level = "debug"
	}
	if options.CallerSkip == 0 {
		options.CallerSkip = 2
	}
	if options.MaxSize == 0 {
		options.MaxSize = 128
	}
	if options.MaxAge == 0 {
		options.MaxAge = 128
	}
	if options.MaxBackups == 0 {
		options.MaxBackups = 7
	}

	// 日志Output是否为stdout或file
	if strings.ToLower(options.Output) != "stdout" && strings.ToLower(options.Output) != "file" {
		panic(OPTION_OUT_ERROR)
	}

	// 日志Output设置
	var writeSyncer zapcore.WriteSyncer
	var rotateLog *lumberjack.Logger
	if strings.ToLower(options.Output) == "stdout" {
		writeSyncer = os.Stdout
	} else {
		rotateLog = &lumberjack.Logger{
			Filename:   options.FilePath(),
			MaxSize:    options.MaxSize,
			MaxAge:     options.MaxAge,
			MaxBackups: options.MaxBackups,
			Compress:   options.Compress,
			LocalTime:  true,
		}

		writeSyncer = zapcore.AddSync(rotateLog)
	}

	// level与encoderConfig设置
	level := zap.NewAtomicLevelAt(parseLevel(options.Level))
	encoderConfig := defaultZapConfig()

	// core设置
	core := options.Core
	if core == nil {
		core = zapcore.NewCore(
			func() zapcore.Encoder {
				if !options.JSONEncode {
					return zapcore.NewConsoleEncoder(encoderConfig)
				}
				return zapcore.NewJSONEncoder(encoderConfig)
			}(),
			writeSyncer,
			level,
		)
	}

	// 默认给warn级别的日志添加堆栈跟踪
	zapOptions := []zap.Option{}
	if options.ShowStacktrace {
		zapOptions = append(zapOptions, zap.AddStacktrace(WarnLevel))
	}

	// caller与初始化fields设置
	if options.AddCaller {
		zapOptions = append(zapOptions, zap.AddCaller(), zap.AddCallerSkip(options.CallerSkip))
	}
	if options.Fields != nil {
		var field []Field
		for k, v := range options.Fields {
			field = append(field, zap.Any(k, v))
		}
		zapOptions = append(zapOptions, zap.Fields(field...))
	}

	logger := zap.New(
		core,
		zapOptions...,
	)

	newLogger := &Logger{
		Options:     options,
		atomicLevel: &level,
		unsugar:     logger,
		sugar:       logger.Sugar(),
		rotateLog:   rotateLog,
		writer:      writeSyncer,
	}

	return newLogger
}

// InitLog 初始化defaultLogger
func InitLog(opts ...Option) {
	var options = defaultOptions
	for _, opt := range opts {
		opt(&options)
	}

	defaultLogger = New(options)
}

// InitLogger 初始化defaultLogger
func InitLogger(options Options) {
	defaultLogger = New(options)
}

// Sugar 返回*zap.SugaredLogger
func (l *Logger) Sugar() *zap.SugaredLogger {
	return l.sugar
}

// Unsugar 返回*zap.Logger
func (l *Logger) Unsugar() *zap.Logger {
	return l.unsugar
}

// Writer 返回zapcore.WriteSyncer
func (l *Logger) Writer() zapcore.WriteSyncer {
	return l.writer
}

// Debugz ...
func (l Logger) Debugz(msg string, fields ...Field) {
	l.unsugar.Debug(msg, fields...)
}

// Debug 打印debug级别的日志
func (l *Logger) Debug(msg string, args ...interface{}) {
	l.sugar.Debugw(msg, args...)
}

// Debugf 格式化打印debug级别的日志
func (l Logger) Debugf(format string, args ...interface{}) {
	l.sugar.Debugf(format, args...)
}

// Infoz ...
func (l *Logger) Infoz(msg string, fields ...Field) {
	l.unsugar.Info(msg, fields...)
}

// Info 打印info级别的日志
func (l *Logger) Info(msg string, args ...interface{}) {
	l.sugar.Infow(msg, args...)
}

// Infof 格式化打印info级别的日志
func (l Logger) Infof(format string, args ...interface{}) {
	l.sugar.Infof(format, args...)
}

// Warnz ...
func (l *Logger) Warnz(msg string, fields ...Field) {
	l.unsugar.Warn(msg, fields...)
}

// Warn 打印warn级别的日志
func (l Logger) Warn(msg string, args ...interface{}) {
	l.sugar.Warnw(msg, args...)
}

// Infof 格式化打印warn级别的日志
func (l Logger) Warnf(format string, args ...interface{}) {
	l.sugar.Warnf(format, args...)
}

// Errorz ...
func (l *Logger) Errorz(msg string, fields ...Field) {
	l.unsugar.Error(msg, fields...)
}

// Error 打印error级别的日志
func (l Logger) Error(msg string, args ...interface{}) {
	l.sugar.Errorw(msg, args...)
}

// Errorf 格式化打印error级别的日志
func (l Logger) Errorf(format string, args ...interface{}) {
	l.sugar.Errorf(format, args...)
}

// Panicz ...
func (l *Logger) Panicz(msg string, fields ...Field) {
	l.unsugar.Panic(msg, fields...)
}

// Panic 打印epanic级别的日志，并发生panic
func (l Logger) Panic(msg string, args ...interface{}) {
	l.sugar.Panicw(msg, args...)
}

// Panicf 格式化打印panic级别的日志，并发生panic
func (l Logger) Panicf(format string, args ...interface{}) {
	l.sugar.Panicf(format, args...)
}

// Fatalz ...
func (l *Logger) Fatalz(msg string, fields ...Field) {
	l.unsugar.Fatal(msg, fields...)
}

// Fatal 打印falal级别的日志，并调用 os.Exit(1) 退出
func (l Logger) Fatal(msg string, args ...interface{}) {
	l.sugar.Fatalw(msg, args...)
}

// Fatalf 格式化打印falal级别的日志，并调用 os.Exit(1) 退出
func (l Logger) Fatalf(format string, args ...interface{}) {
	l.sugar.Fatalf(format, args...)
}

// Level 返回日志当前输出等级
func (l Logger) Level() Level {
	return l.atomicLevel.Level()
}

// SetLevel 设置日志输出等级
func (l *Logger) SetLevel(level Level) *Logger {
	l.atomicLevel.SetLevel(level)
	return l
}

// SetDefaultFields 设置默认字段
func (l *Logger) SetDefaultFields(fields ...Field) *Logger {
	l.fields = append(l.fields, fields...)
	return l
}

func (l *Logger) clone() *Logger {
	cloned := *l
	return &cloned
}

// WithFields 设置初始化fileds
func (l *Logger) WithFields(fields ...Field) *Logger {
	cloned := l.clone()
	var args = make([]interface{}, len(fields))
	for ind, f := range fields {
		args[ind] = f
	}
	cloned.sugar = cloned.sugar.With(args...)
	cloned.unsugar = cloned.unsugar.With(fields...)
	return cloned
}

// SetDefaultDir 设置默认输出目录
func (l *Logger) SetDefaultDir(dir string) *Logger {
	l.Options.Dir = dir
	return l
}

// Sync flushes buffer data into disk
func (l *Logger) Sync() error {
	if err := l.sugar.Sync(); err != nil {
		return err
	}
	if err := l.unsugar.Sync(); err != nil {
		return err
	}
	return nil
}

// Close flushes buffer and closes logger fd
func (l *Logger) Close() error {
	if err := l.Sync(); err != nil {
		return err
	}

	if l.rotateLog != nil {
		return l.rotateLog.Close()
	}

	return nil
}

// FieldErr 打印error级别的日志
func (l Logger) FieldErr(err error) Field {
	return zap.Error(err)
}

// WithOptions设置Option
func (l *Logger) WithOptions(opts ...zap.Option) *Logger {
	cloned := l.clone()
	cloned.unsugar = cloned.unsugar.WithOptions(opts...)
	cloned.sugar = cloned.unsugar.Sugar()
	return cloned
}
