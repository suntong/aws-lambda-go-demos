# Makefile

# name of the bootstrap binary
BOOTSTRAP_BINARY = bootstrap

# Go compiler
GO = go

# all .go files in the current directory
GO_FILES = $(wildcard *.go)

all: aws-deploy.zip bootstrap

# bootstrap depends on all .go files
bootstrap: $(GO_FILES)
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 $(GO) build -tags lambda.norpc -ldflags="-s -w" -o $(BOOTSTRAP_BINARY) .

aws-deploy.zip: $(BOOTSTRAP_BINARY)
	zip aws-deploy.zip bootstrap

# clean up the binaries
clean:
	rm -f aws-deploy.zip $(BOOTSTRAP_BINARY)

# Phony targets to avoid conflicts with files of the same name
.PHONY: clean
