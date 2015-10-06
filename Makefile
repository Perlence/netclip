DOCKER = docker
PACKAGE = github.com/Perlence/netclip
GOOS ?= windows
GOARCH ?= amd64
ifeq "$(GOOS)" "windows"
	ext = .exe
endif

build:
	$(DOCKER) run \
		--rm \
		-v "$(PWD)/Godeps/_workspace/src":/app/src \
		-v "$(PWD)":/app/src/$(PACKAGE) \
		-w /app/src/$(PACKAGE) \
		-e GOPATH=/app \
		-e GOOS=$(GOOS) \
		-e GOARCH=$(GOARCH) \
		golang:1.5 \
		go build -v -o $(notdir $(PACKAGE))-$(GOOS)-$(GOARCH)$(ext)

.PHONY: build
