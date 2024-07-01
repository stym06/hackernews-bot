BINARY_NAME=hn
BUILD_DIR=$(CURDIR)/bin
ZSHRC_PATH=$(HOME)/.zshrc

build:
	go build -o $(CURDIR)/bin
	echo 'alias hn="$(BUILD_DIR)/$(BINARY_NAME)"' >> $(ZSHRC_PATH)

