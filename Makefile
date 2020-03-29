GOCMD=go
GOBUILD=$(GOCMD) build
GOFMT=$(GOCMD) fmt
GOLIST=$(GOCMD) list
GOVET=$(GOCMD) vet

.PHONY: all
all: fmt vet lint build

.PHONY: build
build:
	$(GOBUILD) -o whatsapp-cli -v

.PHONY: fmt
fmt: 
	$(GOFMT) ./...

.PHONY: lint
fmt: 
	$(GOLIST) ./... | grep -v /vendor/ | xargs -L1 golint

.PHONY: vet
vet: 
	$(GOVET) ./...

.PHONY: clean
clean:
	rm whatsapp-cli