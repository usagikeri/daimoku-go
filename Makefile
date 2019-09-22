GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get -u
BINARY_NAME=daimoku

all: deps build
build:
		$(GOBUILD) -o $(BINARY_NAME) -v cmd/main.go
test:
		$(GOTEST) -v ./...
clean:
		$(GOCLEAN)
		rm -f $(BINARY_NAME)
deps:
		$(GOGET) github.com/go-yaml/yaml
compile:
	latexmk *.tex
	latexmk *.tex -c
