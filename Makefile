APP_NAME = gui-bpftool
BUILD_DIR = build/bin
EXECUTABLE = $(BUILD_DIR)/$(APP_NAME)

.PHONY: all build run clean

all: build run

build:
	wails build

run: build
	sudo $(EXECUTABLE)

clean:
	rm -rf $(BUILD_DIR)
