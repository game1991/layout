package log

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitLog(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			assert.Nil(t, err)
		}
	}()
	InitLog(
		WithOutput("file"),
		WithLevelString("info"),
		WithDir("/home/www/logs/test"),
		WithFileName("test.log"),
		WithCallerSkip(1),
		WithRotateConfig(128, 30, 7, false),
		WithFields(map[string]interface{}{"appid": "123456789"}),
	)
}

func TestBuild(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			assert.Nil(t, err)
		}
	}()
	options := Options{
		Output:     "file",
		Level:      "info",
		Dir:        "/home/www/logs/test-build",
		FileName:   "test-build.log",
		AddCaller:  true,
		CallerSkip: 1,
		MaxSize:    128,
		MaxAge:     30,
		MaxBackups: 7,
		Compress:   true,
	}

	options.Build()
}
