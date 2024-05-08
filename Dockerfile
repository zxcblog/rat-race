FROM golang:1.22.3-alpine3.18 AS builder

# ENV 设置环境变量
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct

COPY . /go/src/github.com/zxcblog/rat-race

RUN cd /go/src/github.com/zxcblog/rat-race && go build .

FROM alpine:3.18

# RUN 设置代理镜像
RUN echo -e http://mirrors.ustc.edu.cn/alpine/v3.13/main/ > /etc/apk/repositories

# RUN 设置 Asia/Shanghai 时区
RUN apk --no-cache add tzdata  && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone

# COPY 源路径 目标路径 从镜像中 COPY
COPY --from=builder /go/src/github.com/zxcblog/rat-race/rat-race /opt/

# EXPOSE 设置端口映射
EXPOSE 9000/tcp

# WORKDIR 设置工作目录
WORKDIR /opt

# CMD 设置启动命令
CMD ["./rat-race"]