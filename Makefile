# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
BINARY_NAME=lscontent
BINARY_UNIX=$(BINARY_NAME)_unix

# Installation directory
INSTALL_DIR=/usr/local/bin

all: deps build

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

deps:
	$(GOGET) github.com/sabhiram/go-gitignore
	$(GOGET) github.com/atotto/clipboard

install: build
	mkdir -p $(INSTALL_DIR)
	cp $(BINARY_NAME) $(INSTALL_DIR)

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

uninstall:
	rm -f $(INSTALL_DIR)/$(BINARY_NAME)

# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v

.PHONY: all build deps install clean uninstall build-linux
