FROM --platform=$TARGETPLATFORM scratch as runner

MAINTAINER willgao <will-gao@hotmail.com>

ARG TARGETOS
ARG TARGETARCH

COPY ./build/markdown-blog-${TARGETOS}-${TARGETARCH}/markdown-blog /data/app/

EXPOSE 5006

ENTRYPOINT ["/data/app/markdown-blog", "web"]