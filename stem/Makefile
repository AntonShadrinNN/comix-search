# Определение переменных
GOCMD=go
GOBUILD=$(GOCMD) build
BINARY_NAME=stem

build:
	$(GOCMD) mod tidy
	$(GOBUILD) -o $(BINARY_NAME) ./...

clean:
	rm -f $(BINARY_NAME)

install:
	$(GOBUILD) -o $(GOPATH)/bin/$(BINARY_NAME) -v

uninstall:
	rm -f $(GOPATH)/bin/$(BINARY_NAME)

default: build
