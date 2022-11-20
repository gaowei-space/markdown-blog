FROM golang:alpine as builder

WORKDIR /Users/wei/Space/www/markdown-blog
ENV GOPROXY https://goproxy.cn,direct

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
  apk add --no-cache tzdata make
COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 make build

FROM scratch as runner
COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=builder /Users/wei/Space/www/markdown-blog/bin/markdown-blog /data/app/

EXPOSE 5006

ENTRYPOINT ["/data/app/markdown-blog", "web"]