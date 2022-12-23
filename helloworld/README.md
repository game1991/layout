# Helloworld Layout

## 项目结构说明

### /api

API 定义的目录，如果我们采用的是 grpc 那这里面一般放的就是 proto 文件，除此之外也有可能是 openapi/swagger 定义文件，以及他们生成的文件。

### /cmd

我们一般采用 `/cmd/[appname]/server.go`的形式进行组织

- 首先 cmd 目录下一般是项目的主干目录
- 这个目录下的文件不应该有太多的代码，不应该包含业务逻辑
- `server.go`当中主要做的事情就是负责程序的生命周期，服务所需资源的依赖注入等

main.go 会在/cmd目录下，使用 `cobra.Command`作为命令启动服务。

#### /api

这个是对外暴露的 api 的服务，可以是 `http`, 也可以是 `grpc`

#### /cron

定时任务

#### /genorm

数据库对应的结构 `code generator`。使用的是 `gorm`的 `gen`工具

#### /job

这个用于处理来自 message 的流式任务

#### /scripts/xxx

一次性执行的脚本，有时候会有一些脚本任务

### /configs

配置文件模板或默认配置。一般作为本地开发调试使用。线上环境可使用配置中心。

### /internal

internal 目录下的包，不允许被其他项目中进行导入，这是在 Go 1.4 当中引入的 feature，会在编译时执行

- 所以我们一般会把项目文件夹放置到 internal 当中，例如 /internal/app
- 如果是可以被其他项目导入的包我们一般会放到 pkg 目录下
- 如果是我们项目内部进行共享的包，而不期望外部共享，我们可以放到 /internal/pkg  当中
- 注意 internal 目录的限制并不局限于顶级目录，在任何目录当中都是生效的

#### /conf

内部使用的config的结构定义

#### /controller

针对路由层接口的handler的实现，这一层是对于前端的接口请求做一些简单的参数校验和返回数据处理的，不做太多关于业务的处理，是比较轻的一层。

#### /middlerware

如果使用的是gin框架，结合middlerware层进行一些功能补充

#### /pkg

这一层是内部相互访问调用的公共包

#### /repository

存储层，针对数据操作的的接口定义和实现

#### ~~/router~~

~~这一层是建议通过统一的路由入口，可以方便开发者快速找到对应的接口实现~~

#### /service

这一层是基于controller层传入的请求，进行的业务类逻辑处理，组合各种服务的调用之类的

### /pkg

一般而言，我们在 pkg 目录下放置可以被外部程序安全导入的包。

- pkg 目录下的包一般会按照功能进行区分，例如 /pkg/cache 、 /pkg/conf  等
- 如果你的目录结构比较简单，内容也比较少，其实也可以不使用 pkg  目录，直接把上面的这些包放在最上层即可
- 一般而言我们应用程序 app 在最外层会包含很多文件，例如 .gitlab-ci.yml  Makefile  .gitignore  等等，这种时候顶层目录会很多并且会有点杂乱，建议还是放到 /pkg  目录比较好

### /test

额外的外部测试应用程序和测试数据。

## 参考文献

- [Gin + Grpc同时使用（监听同一端口）_Bearki-CN的博客-CSDN博客_gin grpc](https://blog.csdn.net/weixin_45985984/article/details/124071909)
- [三、grpc-gateway 应用 - 4. grpc-gateway gRPC+gRPC Gateway 能不能不用证书？ - 《Golang Gin 实践》 - 技术池(jishuchi.com)](https://www.jishuchi.com/read/gin-practice/3809)
- [项目结构 | Kratos (go-kratos.dev)](https://go-kratos.dev/docs/intro/layout)
