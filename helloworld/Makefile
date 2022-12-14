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

BIN ?=helloworld
# service name
NAME ?=helloworld

$(info $$BIN=$(BIN))
# $(info $$INTERNAL_PROTO_FILES=$(INTERNAL_PROTO_FILES))
# $(info $$API_PROTO_FILES=$(API_PROTO_FILES))

.PHONY: install info print build config proto swagger tool

install:
	@go get \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		github.com/go-bindata/go-bindata/... \
		github.com/elazarl/go-bindata-assetfs/... 

	@go install \
    	github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    	github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    	google.golang.org/protobuf/cmd/protoc-gen-go \
    	google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		github.com/go-bindata/go-bindata/go-bindata \
		github.com/elazarl/go-bindata-assetfs/go-bindata-assetfs

	@echo "install finished"

info:
	@echo $(GOBIN)
	@echo $(GOPATH)
	@echo $(BIN)
	@echo $(INTERNAL_PROTO_FILES)
	@echo $(API_PROTO_FILES)

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

print:
	echo "version: $(VERSION)"

config:
	@protoc --proto_path=./internal \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./internal \
	       $(INTERNAL_PROTO_FILES)

# --grpc-gateway_out=paths=source_relative:./api
proto:
	@protoc --proto_path=./api \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./api \
 	       --go-grpc_out=paths=source_relative:./api\
	       $(API_PROTO_FILES)

#--openapiv2_opt=disable_service_tags=true

swagger:
#@go-bindata --nocompress -pkg swagger -o pkg/swagger/data.go third_party/swagger-ui/...
	@go-bindata-assetfs -pkg swagger -o pkg/swagger/asset.go third_party/swagger-ui/...
	@protoc --proto_path=./api \
			--proto_path=./third_party \
			--openapiv2_out=logtostderr=true:./api/openapi \
			$(API_PROTO_FILES)

#usage:
# grpcurl.
## 1. grpcurl localhost:8080 list helloworld.api.v1.HelloworldService //??????????????????????????????
## 2. grpcurl localhost:8080 describe helloworld.api.v1.HelloworldService.Login //??????????????????????????????
## 3. grpcurl -d '{"userName": "hello", "password": "world"}' localhost:8080 helloworld.api.v1.HelloworldService/Login
# grpcui.
## grpcui -plaintext ??????app????????????  //????????????ui?????????????????????grpc????????????
## eg: grpcui -plaintext localhost:8080
tool:
	@go get \
		github.com/fullstorydev/grpcurl/... \
		github.com/fullstorydev/grpcui/...

	@go install \
		github.com/fullstorydev/grpcurl/cmd/grpcurl \
		github.com/fullstorydev/grpcui/cmd/grpcui \

	@echo "install grpc tools finished"
