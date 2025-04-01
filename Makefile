APP_NAME = gui-bpftool
BUILD_DIR = build/bin
EXECUTABLE = $(BUILD_DIR)/$(APP_NAME)

.PHONY: all build run clean

all: build run

dev: build run

go:
	go build && sudo ./gui-bpftool

build:
	wails build -devtools -tags debug

run: build
	sudo $(EXECUTABLE)

clean:
	rm -rf $(BUILD_DIR)
