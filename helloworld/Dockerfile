
#syntax=docker/dockerfile:1
##
## Build
##
#FROM golang:1.19-bullseye AS build

ARG IMAGE
FROM ${IMAGE} AS builder
WORKDIR /src
ADD . .
ARG NAME
RUN sed -i "s@http://deb.debian.org@http://mirrors.aliyun.com@g" /etc/apt/sources.list \
	&& rm -Rf /var/lib/apt/lists/* \
	&& apt-get update \
	&& apt-get install net-tools -y \
	&& echo go env \
	&& go env -w GOPROXY=https://goproxy.cn,direct \
	&& go env -w GOPRIVATE=git.xq5.com \
	&& make build NAME=${NAME} \
	&& mkdir -p /src/data


##
## Deploy
##
FROM debian:bullseye-slim
#FROM repository.zonst.com/devops/alpine:base_npc1
ARG PROJECT_NAME
ARG NAME
#设置环境变量（程序运行时使用）
ENV PROJECT_NAME=${PROJECT_NAME}
WORKDIR /data/projects/{PROJECT_NAME}

COPY --from=builder /src/${NAME}          /data/projects/${PROJECT_NAME}/app
COPY --from=builder /src/configs          /data/projects/${PROJECT_NAME}/configs
COPY --from=builder /src/data             /data/projects/${PROJECT_NAME}/data

ENTRYPOINT [ "/data/projects/${PROJECT_NAME}/app" ]
CMD [ "/bin/bash" ]
