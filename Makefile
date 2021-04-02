VERSION := $(shell git describe --tags)
BUILD := $(shell git rev-parse --short HEAD)
PROJECTNAME := $("butler-burton")

# Go related variables.
GOFILES := $(wildcard *.go)

# Redirect error output to a file, so we can show it in development mode.
STDERR=/tmp/.$(PROJECTNAME)-stderr.txt

# Use linker flags to provide version
LDFLAGS=-ldflags "-X=main.Version=$(VERSION)"

# Make is verbose in Linux. Make it silent.
MAKEFLAGS += --silent

## install: install app, runs 'go install' internally
install:
	go install $(LDFLAGS)

## build: build binary, runs 'go build' internally
build:
	go build $(LDFLAGS) -o $(PROJECTNAME) 

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
