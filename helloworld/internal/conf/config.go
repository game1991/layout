package conf

import (
	"errors"

	"helloworld/internal/pkg/env"
	"helloworld/internal/pkg/store"
	"helloworld/pkg/log"
	"os"

	"github.com/ismdeep/args"
	"github.com/spf13/viper"
)

// RunMode 运行模式
type RunMode string

// 运行模式
const (
	PROD  RunMode = "prod"
	PRE   RunMode = "pre"
	DEV   RunMode = "dev"
	DEBUG RunMode = "debug"
)

// Value 对应的值
func (rm RunMode) Value() int {
	switch rm {
	case DEV:
		return 1
	case PRE:
		return 2
	case PROD:
		return 4
	case DEBUG:
		return 8
	default:
		return 0
	}
}

// Has 存在运行模式，参数runmode是集合，表示这些集合中是否包含rm
func (rm RunMode) Has(runmode ...RunMode) bool {
	if rm.Value() == 0 {
		return false
	}
	var flag int
	for _, item := range runmode {
		flag |= item.Value()
	}
	if flag&rm.Value() == rm.Value() {
		return true
	}
	return false

}

// env name
const (
	FileKey = "ZONST_CONFIG_FILE"
	PathKey = "ZONST_CONFIG_PATH"
)

var appConfig *APPConfig

// InitConfig init config
func InitConfig() (err error) {
	viper.SetConfigType("toml")
	viper.AddConfigPath("./configs/")

	// 指定配置文件夹路径
	if args.Exists("-d") {
		viper.AddConfigPath(args.GetValue("-d"))
	}

	// config file define by env
	configFile := os.Getenv(FileKey)
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName("test")
	}
	// config path define by env
	configPath := os.Getenv(PathKey)
	if configPath != "" {
		viper.AddConfigPath(configPath)
	}

	if err = viper.ReadInConfig(); err != nil {
		return err
	}

	// OVERWRITE LOCAL CONFIG
	if env.IsDevelopment() {
		v := viper.New()
		v.SetConfigName("local")
		v.SetConfigType("toml")
		v.AddConfigPath("./configs/")
		if args.Exists("-d") {
			v.AddConfigPath(args.GetValue("-d"))
		}

		// config path define by env
		configPath := os.Getenv(PathKey)
		if configPath != "" {
			v.AddConfigPath(configPath)
		}

		if err = v.ReadInConfig(); err == nil {
			if err = viper.MergeConfigMap(v.AllSettings()); err != nil {
				panic(err)
			}
		}
	}
	// read app config
	appConfig, err = readAPPConfig()
	if err != nil {
		return
	}

	return nil
}

// GetConfig 获取配置文件信息
func GetConfig() *APPConfig {
	if appConfig == nil {
		panic("appConfig is nil")
	}
	return appConfig
}

// Log log
func Log() (conf *log.Options, err error) {
	conf = &log.Options{}
	err = viper.UnmarshalKey("log", conf)
	return
}

// GetServer get server conf info
func GetServer() (conf *Server, err error) {
	conf = &Server{}
	err = viper.UnmarshalKey("Server", conf)
	if conf == nil {
		err = errors.New("Server 配置为空")
	}
	return
}

// MySQL ...
func MySQL() (conf *store.Config, err error) {
	err = viper.UnmarshalKey("MySQL", &conf)
	if conf == nil {
		err = errors.New("MySQL 配置为空")
	}
	return
}

func readAPPConfig() (conf *APPConfig, err error) {
	if appConfig == nil {
		conf = &APPConfig{}
		err = viper.UnmarshalKey("app", conf)
		// if len(conf.Referer) == 0 {
		// 	conf.Referer = []string{"*.zonst.com", "*.zonst.com"}
		// }
		if conf == nil {
			err = errors.New("app 配置为空")
		}
	}
	return
}

// String get single string value
func String(key string) string {
	return viper.GetString(key)
}

// Strings get string list
func Strings(key string) []string {
	return viper.GetStringSlice(key)
}

// Int get int value
func Int(key string) int {
	return viper.GetInt(key)
}

// Int64 get int64 value
func Int64(key string) int64 {
	return viper.GetInt64(key)
}

// Bool get bool value
func Bool(key string) bool {
	return viper.GetBool(key)
}

// Float get float value
func Float(key string) float64 {
	return viper.GetFloat64(key)
}

// UnmarshalKey unmarshal key
func UnmarshalKey(prefix string, obj interface{}) error {
	return viper.UnmarshalKey(prefix, obj)
}
