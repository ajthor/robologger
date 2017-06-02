FROM golang:1.8-alpine

WORKDIR /go/src/github.com/gorobot/robologger

RUN apk add --no-cache \
    gcc \
    git \
    glide \
    musl-dev \
    nano \
    openssl \
    make

CMD [ "sh" ]
