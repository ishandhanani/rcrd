# Makefile for rcrd

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
BINARY_NAME=rcrd
INSTALL_PATH=/usr/local/bin

# Build target
build:
	$(GOBUILD) -o $(BINARY_NAME) -v

# Install target
install: build
	mv $(BINARY_NAME) $(INSTALL_PATH)

# Clean target
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

# Run target
run: build
	./$(BINARY_NAME)

.PHONY: build install clean run
