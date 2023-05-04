package log

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"go.uber.org/zap/zapcore"
)

// Option ...
type Option func(*Options)

// Options ...
type Options struct {
	Output         string                 // 日志输出方式："stdout" 或者 "file"
	JSONEncode     bool                   // 是否开启JSON格式
	Dir            string                 // 日志文件路径
	FileName       string                 // 日志文件名称
	Level          string                 // 日志等级
	Fields         map[string]interface{} // 初始化fields
	AddCaller      bool                   // 是否打印Caller信息
	CallerSkip     int                    // Caller跳过数
	ShowStacktrace bool                   // Warn和Error级别日志是否显示堆栈信息
	MaxSize        int                    // 日志文件最大size(MB)
	MaxAge         int                    // 旧日志文件最大保存时间（天）
	MaxBackups     int                    // 旧日志文件保留最大数量
	Compress       bool                   // 是否开启压缩日志文件
	Core           zapcore.Core           // 若Core!=nil, zap.New时使用该参数
}

// Build 通过Option构建出Logger
func (options Options) Build() *Logger {
	return New(options)
}

// defaultZapConfig 默认配置
func defaultZapConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     timeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

// defaultOptions 默认Options
var defaultOptions = Options{
	Output:         "stdout",
	Level:          "debug",
	MaxSize:        128, // 单位：MB
	MaxAge:         30,  // 单位：天
	MaxBackups:     7,
	Compress:       false,
	CallerSkip:     2,
	AddCaller:      true,
	ShowStacktrace: false,
	JSONEncode:     true,
}

// WithOutput 设置日志输出方式
func WithOutput(out string) Option {
	return func(o *Options) {
		o.Output = out
	}
}

// WithJSONEncode 设置日志格式
func WithJSONEncode(isJSON bool) Option {
	return func(o *Options) {
		o.JSONEncode = isJSON
	}
}

// WithDir 设置日志文件路径
func WithDir(dir string) Option {
	return func(o *Options) {
		dir = strings.TrimSuffix(dir, "/")
		o.Dir = dir
	}
}

// WithFileName 设置日志文件名称
func WithFileName(fileName string) Option {
	return func(o *Options) {
		o.FileName = fileName
	}
}

// WithLevelString 设置日志等级,参数为string类型
func WithLevelString(level string) Option {
	return func(o *Options) {
		o.Level = level
	}
}

// WithAddCaller 设置打印caller信息
func WithAddCaller(addCaller bool) Option {
	return func(o *Options) {
		o.AddCaller = addCaller
	}
}

// WithShowStacktrace 设置Warn和Error级别日志是否显示堆栈信息
func WithShowStacktrace(showStacktrace bool) Option {
	return func(o *Options) {
		o.ShowStacktrace = showStacktrace
	}
}

// WithCallerSkip 设置caller跳过数
func WithCallerSkip(skip int) Option {
	return func(o *Options) {
		o.CallerSkip = skip
	}
}

// WithRotateConfig 设置轮转配置
func WithRotateConfig(maxSize, maxAge, maxBackups int, compress bool) Option {
	return func(o *Options) {
		o.MaxSize = maxSize
		o.MaxAge = maxAge
		o.MaxBackups = maxBackups
		o.Compress = compress
	}
}

// WithFields 设置初始化fields
func WithFields(fields map[string]interface{}) Option {
	return func(o *Options) {
		o.Fields = fields
	}
}

// WithCore 设置Core
func WithCore(core zapcore.Core) Option {
	return func(o *Options) {
		o.Core = core
	}
}

// parseLevel 日志等级映射
func parseLevel(lvl string) (lv Level) {
	level := strings.ToUpper(strings.TrimSpace(lvl))
	switch level {
	case "DEBUG":
		lv = DebugLevel
	case "INFO":
		lv = InfoLevel
	case "WARN":
		lv = WarnLevel
	case "ERROR":
		lv = ErrorLevel
	case "PANIC":
		lv = PanicLevel
	case "FATAL":
		lv = FatalLevel
	default:
		lv = DebugLevel
	}

	return
}

// timeEncoder 设置时间格式
func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

// FilePath 获取日志文件全路径
func (options Options) FilePath() string {
	dir := options.Dir
	// 若Dir为空字符串，则从基础环境变量库中取值("/home/www/logs/"+AppName)
	// 注意：开发环境使用go run main.go运行时，AppName为main
	if dir == "" {
		dir = DirOfLog()
	}

	fileName := options.FileName
	// 若FileName为空字符串，默认日志文件名为default.log
	if fileName == "" {
		fileName = DEFAULT_LOG_FILE_NAME
	}

	path := fmt.Sprintf("%s/%s", dir, fileName)

	return path
}

// DirOfLog ...
func DirOfLog() string {
	// LogDir gets application log directory.
	logDir := os.Getenv("_APP_LOG_DIR")
	if logDir == "" {
		logDir = "/home/www/logs/"
	}
	return fmt.Sprintf("%s%s/", logDir, AppName())
}

// AppName gets application name.
func AppName() string {
	appName := os.Getenv("_APPNAME")
	if appName == "" {
		appName = filepath.Base(os.Args[0])
	}
	return appName
}
