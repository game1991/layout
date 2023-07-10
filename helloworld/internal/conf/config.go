package conf

import (
	"errors"

	"os"

	"github.com/game1991/layout/helloworld/internal/pkg/env"
	"github.com/game1991/layout/helloworld/internal/pkg/store"
	"github.com/game1991/layout/helloworld/pkg/log"

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
	return flag&rm.Value() == rm.Value()
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
	err = viper.UnmarshalKey("Server", &conf)
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

// Session conf ...
func SessionConf() (conf *Session, err error) {
	err = viper.UnmarshalKey("Session", &conf)
	if conf == nil {
		err = errors.New("Session 配置为空")
	}
	if conf.SessionName == nil {
		err = errors.New("SessionName 配置为空")
	}
	return
}

func (s *Session) SessionSecret() [][]byte {
	result := make([][]byte, len(s.Secret))
	for i, item := range s.Secret {
		if i%2 == 1 {
			bts := []byte(item)
			switch len(bts) {
			case 16, 24, 32:
				result[i] = bts
			default:
				panic("session Secret 设置错误")
			}
		} else {
			result[i] = []byte(item)
		}
	}
	return result
}

func (s *Session) SessionNames() (result []string) {
	for _, v := range s.SessionName {
		result = append(result, v)
	}
	return
}

// 获取session名称 入参填写配置文件中设置的key键
func (s *Session) GetSessionNameFromKey(key string) string {
	if v, ok := s.SessionName[key]; ok {
		return v
	}
	return ""
}

func Horus() (cfg *Hours, err error) {
	if err := viper.UnmarshalKey("horus", &cfg); err != nil {
		panic("unmarshal horus config failed")
	}
	return cfg, nil
}

func readAPPConfig() (conf *APPConfig, err error) {
	if appConfig == nil {
		err = viper.UnmarshalKey("app", &conf)
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
