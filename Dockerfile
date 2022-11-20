FROM golang:alpine as builder

RUN mkdir -p /app
COPY . /app

ENV GOROOT /usr/local/go
ENV PATH $PATH:$GOROOT/bin
ENV GOPROXY https://goproxy.cn,direct
ENV GO111MODULE on

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
  apk add --no-cache tzdata
RUN cd /app && \
  CGO_ENABLED=0 go build -o markdown-blog ./

FROM scratch as runner

MAINTAINER willgao <will-gao@hotmail.com>

COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=builder /app/markdown-blog /data/app/

EXPOSE 5006

ENTRYPOINT ["/data/app/markdown-blog", "web"]