package log

import (
	"errors"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Level 日志等级
const (
	DebugLevel = zap.DebugLevel
	InfoLevel  = zap.InfoLevel
	WarnLevel  = zap.WarnLevel
	ErrorLevel = zap.ErrorLevel
	PanicLevel = zap.PanicLevel
	FatalLevel = zap.FatalLevel
)

type (
	Field = zap.Field
	Level = zapcore.Level
)

var (
	Skip       = zap.Skip
	Binary     = zap.Binary
	Bool       = zap.Bool
	ByteString = zap.ByteString
	Complex128 = zap.Complex128
	Complex64  = zap.Complex64
	Float64    = zap.Float64
	Float32    = zap.Float32
	Int        = zap.Int
	Int32      = zap.Int32
	Int64      = zap.Int64
	Int16      = zap.Int16
	Int8       = zap.Int8
	String     = zap.String
	Uint       = zap.Uint
	Uint64     = zap.Uint64
	Uint32     = zap.Uint32
	Uint16     = zap.Uint16
	Uint8      = zap.Uint8
	Uintptr    = zap.Uintptr
	Reflect    = zap.Reflect
	Namespace  = zap.Namespace
	Stringer   = zap.Stringer
	Time       = zap.Time
	Stack      = zap.Stack
	Duration   = zap.Duration
	Object     = zap.Object
	Any        = zap.Any
)

// 默认的日志文件名称
const DEFAULT_LOG_FILE_NAME = "default.log"

// 异常
var (
	OPTION_OUT_ERROR = errors.New(`Output参数错误，必须为"stdout"或"file"`)
)
