#取git 仓库地址的服务名 认为是该服务的名称
PACKAGE  := $(subst $(shell git rev-parse --show-toplevel),,$(CURDIR))
ROOT := $(shell git remote -v | grep '^origin\s.*(fetch)$$' | awk '{print $$2}' | sed -E 's/^.*(\/\/|@)(.*)\.git\/?$$/\2/' | sed 's/:/\//g')
PACKAGENAME = $(ROOT)$(PACKAGE)

APP      := $(shell basename $(PACKAGENAME))
OUTPUT := $(CURDIR)/output/bin/
.DEFAULT: build
build:
	go fmt ./...
	go build  -o "$(OUTPUT)$(APP)" main.go
run:
	nohup $(OUTPUT)$(APP) &
stop:
	ps -ef |grep $(APP) |grep -v grep|awk '{print $$2}'|xargs kill -USR2


.PHONY: build stop run