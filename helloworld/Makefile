# 
DBG_MAKEFILE ?=
ifeq ($(DBG_MAKEFILE),1)
    $(warning ***** starting Makefile for goal(s) "$(MAKECMDGOALS)")
    $(warning ***** $(shell date))
else
    # If we're not debugging the Makefile, don't echo recipes.
    MAKEFLAGS += -s
endif

# We don't need make's built-in rules.
MAKEFLAGS += --no-builtin-rules
# Be pedantic about undefined variables.
MAKEFLAGS += --warn-undefined-variables

.SUFFIXES:

# This version-strategy uses git tags to set the version string
VERSION ?= $(shell git describe --tags --always --dirty)
TAG ?= $(shell git describe --abbrev=0 --tags)
BRANCH ?= $(shell git symbolic-ref --short -q HEAD)
LASTTIME ?= $(shell git log -1 --format="%ct")
LASTCOMMIT ?= $(shell git describe --always)
GitRepoPATH ?= $(shell git config  --local  --get remote.origin.url)

# Which Go modules mode to use ("mod" or "vendor")
MOD ?= mod

# Satisfy --warn-undefined-variables.
GOFLAGS ?=
HTTP_PROXY ?=
HTTPS_PROXY ?=

# os env
OS := $(if $(GOOS),$(GOOS),$(shell go env GOOS))
ARCH := $(if $(GOARCH),$(GOARCH),$(shell go env GOARCH))
#GOHOSTOS:=$(shell go env GOHOSTOS)
#os	:=$(shell expr substr $(shell uname -s) 1 5)
#ifeq ($(os),MINGW)
ifeq ($(OS), windows)
# GOBIN :=$(subst \,/,$(GOBIN))
# GOPATH :=$(subst \,/,$(GOPATH))
# $(info path alter $$GOBIN=$(GOBIN))
# $(info path alter $$GOPATH=$(GOPATH))

# ifneq ($(GOBIN),)
# GOBIN :=/$(shell echo "$(GOBIN)" | cut -d';' -f1 | sed 's/://g')
# endif
# ifneq ($(GOPATH),)
# GOPATH :=/$(shell echo "$(GOPATH)" | cut -d';' -f1 | sed 's/://g')
# endif
# $(info cut first $$GOBIN=$(GOBIN))
# $(info cut first $$GOPATH=$(GOPATH))

#the `find.exe` is different from `find` in bash/shell.
#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
#Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
# Git_Bash=/$(subst \,/,$(subst cmd\,bin\bash.exe,$(dir $(shell where git | grep cmd))))
# $(info $$dir=$(subst cmd\git.exe,bin\bash.exe,$(shell where git | grep cmd | sed 's/://g')))
$(info $$Git_Bash=/$(subst \,/,$(subst cmd\git.exe,bin\bash.exe,$(shell where git | grep cmd | sed 's/://g'))))
Git_Bash=/$(subst \,/,$(subst cmd\git.exe,bin\bash.exe,$(shell where git | grep cmd | sed 's/://g')))
INTERNAL_PROTO_FILES=$(shell '$(Git_Bash)' -c "find internal -name *.proto")
API_PROTO_FILES=$(shell '$(Git_Bash)' -c "find api -name *.proto")
else
INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
API_PROTO_FILES=$(shell find api -name *.proto)
endif

#project name
PROJECT_NAME :=$(shell echo $(shell pwd)|awk -F '/' '{ print $$NF }')
#image
IMAGE ?=golang:1.19-bullseye
# bin name
BIN ?=helloworld
# service name
NAME ?=helloworld
# pb generate file
API_PB_FILES ?=./api/proto/*/
# config pb generate file
CONF_PB_FILES ?=./internal/conf/

$(info $$BIN=$(BIN))
# $(info $$INTERNAL_PROTO_FILES=$(INTERNAL_PROTO_FILES))
# $(info $$API_PROTO_FILES=$(API_PROTO_FILES))

# docker buildkit
DOCKER_BUILDKIT:=$(if $(DOCKER_BUILDKIT),$(DOCKER_BUILDKIT),DOCKER_BUILDKIT=1)
# image base
DOCKER_BASE:=$(if $(DOCKER_BASE),$(DOCKER_BASE),local/app)
# image tag
DOCKER_TAG:=$(if $(DOCKER_TAG),$(DOCKER_TAG),$(shell echo $(ARCH)|awk '{ sub(/\//,"-"); print $$0 }'))

.PHONY: install info print build config proto swagger tool run docker

install:
	@go get \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		github.com/favadi/protoc-go-inject-tag \
		github.com/DATA-DOG/go-sqlmock \
		github.com/alicebob/miniredis/v2 \
		github.com/go-bindata/go-bindata/... \
		github.com/go-bindata/go-bindata/go-bindata \
		github.com/elazarl/go-bindata-assetfs/... \

	@go install \
    	github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    	github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    	google.golang.org/protobuf/cmd/protoc-gen-go \
    	google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		github.com/go-bindata/go-bindata/go-bindata \
		github.com/elazarl/go-bindata-assetfs/go-bindata-assetfs \
		github.com/favadi/protoc-go-inject-tag \
		github.com/DATA-DOG/go-sqlmock \
		github.com/alicebob/miniredis/v2 \

	@echo "install finished"

info:
	@echo $(PROJECT_NAME)
	@echo $(GOBIN)
	@echo $(GOPATH)
	@echo $(BIN)
	@echo $(INTERNAL_PROTO_FILES)
	@echo $(CONF_PB_FILES)
	@echo $(API_PROTO_FILES)
	@echo $(API_PB_FILES)

build:
	chmod -R a+rx ./build/scripts
	/bin/sh -c  "             \
        BIN=$(BIN)            \
        ARCH=$(ARCH)          \
        OS=$(OS)              \
        VERSION=$(VERSION)    \
		NAME=$(NAME)          \
		BRANCH=$(BRANCH)      \
		TAG=$(TAG)            \
		LASTTIME=$(LASTTIME)    \
		LASTCOMMIT=$(LASTCOMMIT)    \
		GitRepoPATH=$(GitRepoPATH)          \
        MOD=$(MOD)            \
        GOFLAGS=$(GOFLAGS)    \
        ./build/scripts/build.sh \
	"
run: swagger build
	${BIN} serve -d ./configs

print:
	echo "version: $(VERSION)"

config:
	@protoc --proto_path=./internal \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./internal \
	       $(INTERNAL_PROTO_FILES)
	@protoc-go-inject-tag  -input=${CONF_PB_FILES}*.pb.go
# --grpc-gateway_out=paths=source_relative:./api
proto:
	@protoc --proto_path=./api \
	        --proto_path=./third_party \
 	        --go_out=paths=source_relative:./api \
 	        --go-grpc_out=paths=source_relative:./api\
	        $(API_PROTO_FILES)
	@protoc-go-inject-tag  -input=${API_PB_FILES}*.pb.go
#--openapiv2_opt=disable_service_tags=true

swagger:
# 使用插件生成swagger.json文件
	@protoc --proto_path=./api \
			--proto_path=./third_party \
			--openapiv2_out=logtostderr=true:./api/swagger \
			--experimental_allow_proto3_optional \
			$(API_PROTO_FILES)
# 
# @go-bindata --nocompress -pkg swagger -o pkg/swagger/data.go third_party/swagger-ui/...

#swagger.json
	@go-bindata-assetfs -pkg swagger -o pkg/swagger/asset.go api/swagger/proto/v1/...
#swagger-ui
	@go-bindata-assetfs -pkg swaggerUI -o pkg/swagger-ui/asset.go third_party/swagger-ui/...

# 
# @go-bindata --nocompress -pkg swagger -o pkg/swagger/data.go third_party/swagger-ui/...

#swagger.json
	@go-bindata-assetfs -pkg swagger -o pkg/swagger/asset.go api/swagger/proto/v1/...
#swagger-ui
	@go-bindata-assetfs -pkg swaggerUI -o pkg/swagger-ui/asset.go third_party/swagger-ui/...

#usage:
# grpcurl.
## 1. grpcurl localhost:8080 list helloworld.api.v1.HelloworldService //展示当前服务可用接口
## 2. grpcurl localhost:8080 describe helloworld.api.v1.HelloworldService.Login //显示当前接口详情内容
## 3. grpcurl -d '{"userName": "hello", "password": "world"}' localhost:8080 helloworld.api.v1.HelloworldService/Login
# grpcui.
## grpcui -plaintext 您的app服务地址  //开启一个ui可视化界面操作grpc接口调用
## eg: grpcui -plaintext localhost:8080
tool:
	@go get \
		github.com/fullstorydev/grpcurl/... \
		github.com/fullstorydev/grpcui/...

	@go install \
		github.com/fullstorydev/grpcurl/cmd/grpcurl \
		github.com/fullstorydev/grpcui/cmd/grpcui \

	@echo "install grpc tools finished"

docker:
	$(DOCKER_BUILDKIT) docker build \
	--no-cache \
	--build-arg IMAGE=$(IMAGE) \
	--build-arg NAME=$(NAME) \
	--build-arg PROJECT_NAME=$(PROJECT_NAME) \
	-t $(DOCKER_BASE):$(DOCKER_TAG) \
	-f Dockerfile . || exit 1
