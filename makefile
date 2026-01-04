BINARY_NAME=qrgen.exe
SOURCE=./cmd/main.go
DIST_DIR=./dist

all: build

build:
	@if not exist "$(DIST_DIR)" mkdir "$(DIST_DIR)"
	go build -o $(DIST_DIR)/$(BINARY_NAME) $(SOURCE)

clean:
	@if exist "$(DIST_DIR)" rmdir /s /q "$(DIST_DIR)"

.PHONY: all build clean