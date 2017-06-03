.PHONY: all

IMAGE_NAME := gorobot/robologger:dev

DEV_OPTS := --rm -it -v "$$PWD:/go/src/github.com/gorobot/robologger" -v "/var/run/docker.sock:/var/run/docker.sock" --name robolog_dev

default: dev

all: dev

build-dev:
	docker build -t $(IMAGE_NAME) -f Dockerfile .

dev: build-dev
	docker run $(DEV_OPTS) $(IMAGE_NAME) sh
