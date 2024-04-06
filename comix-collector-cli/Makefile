GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_PATH=cmd/xkcd
MAIN_FILE=main.go
BINARY_NAME=xkcdparse

build:
	$(GOCMD) mod tidy
	$(GOBUILD) -o $(BINARY_NAME) $(BINARY_PATH)/$(MAIN_FILE)

clean:
	rm -f $(BINARY_NAME)

install:
	$(GOBUILD) -o $(GOPATH)/bin/$(BINARY_NAME) -v $(BINARY_PATH)/$(MAIN_FILE)

uninstall:
	rm -f $(GOPATH)/bin/$(BINARY_NAME)

default: build
