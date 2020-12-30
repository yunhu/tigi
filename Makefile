#取git 名当服务名
SUB_PACKAGE  := $(subst $(shell git rev-parse --show-toplevel),,$(CURDIR))
PACKAGE_ROOT := $(shell git remote -v | grep '^origin\s.*(fetch)$$' | awk '{print $$2}' | sed -E 's/^.*(\/\/|@)(.*)\.git\/?$$/\2/' | sed 's/:/\//g')
PACKAGE_NAME = $(PACKAGE_ROOT)$(SUB_PACKAGE)

APP      := $(shell basename $(PACKAGE_NAME))
OUTPUT := $(CURDIR)/output/bin/
.DEFAULT: build

build:
	go fmt ./...
	go build  -o "$(OUTPUT)$(APP)" main.go
run:
	nohup $(OUTPUT)$(APP) &
stop:
	ps -ef |grep $(APP) |grep -v grep|awk '{print $$2}'|xargs kill -9


.PHONY: build stop run