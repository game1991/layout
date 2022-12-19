# Log

log库对zap库 ([github.com/uber-go/zap](http://github.com/uber-go/zap "github.com/uber-go/zap")) 进行了简要的封装，支持输出json格式的日志，日志分割，动态调整日志级别，实现了go-micro/logger接口。

## 安装

------

```go
go get github.com/game1991/service-lib-log
```

## 快速开始

------

```go
package main

import (
	"github.com/game1991/service-lib-log"
)

func main() {
	//log.SetLevelString("info")  //设置日志输出等级，若不设置默认输出debug级别以上的日志
	log.Info("fetch URL success", "url", "www.chinauos.com")    //zap.SugaredLogger 写法
	log.Infof("fetch URL:%s success", "www.chinauos.com")
	log.Errorz("failed to fetch URL",                           //zap.Logger 写法
		log.String("url", "www.chinauos.com"),
		log.String("err", "url format is not correct"),
	)
}
```

### 日志示例

------



```json
{
	"level":"error",
	"time":"2020-07-02 14:53:27",
	"caller":"test/main.go:11",
	"msg":"failed to fetch URL",
	"url":"www.chinauos.com",
	"err":"url format is not correct"
}
```

- level：日志等级。
- time：日志打印时间。
- caller：调用log库的代码所在处。
- msg：日志描述。
- err：具体错误信息。

#### 初始化

------

##### 函数式选项模式

```go
package main

import (
	"github.com/game1991/service-lib-log"
)

func main() {
	log.InitLog(
		log.WithOutput("file"),
		log.WithJSONEncode(true),
		log.WithLevelString("info"),
		log.WithDir("/home/www/logs"),
		log.WithFileName("test.log"),
		log.WithAddCaller(true),
		log.WithCallerSkip(2),
		log.WithRotateConfig(128, 30, 7, false),
		log.WithFields(map[string]interface{}{"appid": "123456789"}),
	)
	log.Info("fetch URL success", "url", "www.chinauos.com")
}

```

- WithOutput(output string)：设置日志输出方式，可取参数值为stdout(标准输出)和file，若为stdout，则与文件相关的设置均无用。

- WithLevelString(level string)：设置日志输出等级，可取参数值为debug、info、warn、error、panic、fatal。

- WithJSONEncode(isJSON bool)：设置是否JSON输出，默认值为true。

- log.WithDir(dir string)：设置日志文件路径，默认从基础环境变量中_APP_LOG_DIR取值，如果未配置环境变量，则默认地址（"/home/www/logs" + appName），本地开发时需给该目录授权：

  ```bash
  sudo mkdir -p /home/www/logs
  sudo chmod 777 /home/www/logs
  ```

- WithFileName(filename string)：设置日志文件名，默认值为default.log。

- WithAddCaller(caller bool)：是否打印caller，默认值为true。

- WithCallerSkip(skip int)：caller跳过数，默认值为2。

- WithRotateConfig(maxSize, maxAge, maxBackups int, compress bool)：设置日志rotate信息。

  - maxsize：日志文件最大 size（MB），默认值为128。

  - maxAge：旧日志文件最大保存时间（天），默认值为30。

  - maxBackups：旧日志文件保留最大数量，默认值为7。

  - compress：是否需要压缩滚动日志，默认值为false。

- WithFields(fields map[string]interface{})：设置日志默认字段与其值。

##### 结构体参数

```go
package main

import (
	"github.com/game1991/service-lib-log"
	"config"
)

func main() {
	type MyfConfig struct {
		Log log.Options // log配置
	}
	var cfg MyfConfig
	err := config.Load("conf.toml", &cfg)
	if nil != err {
		panic(err)
	}

	log.InitLogger(cfg.Log)
	log.Info("fetch URL success", "url", "www.chinauos.com")
}

```

###### 配置文件格式

```toml
[Log]
	output = "file"
	level = "info"
	jsonEncode = true
	dir = "/home/www/logs/example"
	fileName = "test.log"
	addCaller = true
	callerSkip = 2
	maxSize = 128
	maxAge = 30
	maxBackups = 7
	compress = true
	[Log.fields]
		appName = "example"
```

#### 在go-micro中使用

------

go-micro（<https://github.com/micro/go-micro>) 中默认使用的是自己的logger日志库，它默认打印的是非json格式的日志，这里将其接口基于zap日志库进行了实现，从而打印json日志，具体用法如下：

```go
package main

import (
	"pkg.deepin.com/service/lib/log"
	"github.com/micro/go-micro/v2/logger"
)

func main() {
	logger.DefaultLogger = log.NewMicroLogger()
	logger.Fields(map[string]interface{}{"key3": "val4"}).Log(logger.DebugLevel, "test_msg")
}
```

#### 动态调整日志级别

------

```bash
kill -10 pid      //切换日志级别至debug
kill -12 pid      //切换日志级别至error
```

#### 日志格式建议

------

日志主要记录在何时，何模块，发生什么，参数是什么。格式建议如下：

```go
log.Error("user login failed", "theme", "user", "username", “lsw”, "err", err.Error())
```

- 第一个字段填写日志描述。
- 第二三个字段填写主题，表示当前模块。其中theme是固定字段，代表key。
- 第四五个字段填写需要记录的参数。便于复现问题时使用，分别代表key和value，若参数有多个请依次填写。
- 最后两个字段记录具体错误信息。其中err是固定字段。

#### 日志级别定义

------

日志级别由低到高：

- debug：该信息能够提供给开发人员，帮助其定位系统运行的路径以及产生问题的场景和数据。
- info：在正常情况下需要被记录的重要信息，例如：系统初始化成功，服务启动或者停止以及成功的处理了重要的业务。查看日志中的Info信息，能够看到应用提供服务的主要状态变更，但是也不要记录过多的Info信息。
- warn：记录会出现潜在错误的情形。例如大量时延过大等。
- error：记录虽然发生错误事件，但仍然不影响系统的继续运行。
- panic：记录严重的错误事件，希望可以用recover对错误进行处理。
- fatal：记录严重的错误事件，它将会导致应用程序的退出。
