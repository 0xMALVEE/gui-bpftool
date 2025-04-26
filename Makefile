APP_NAME = gui-bpftool
BUILD_DIR = build/bin
EXECUTABLE = $(BUILD_DIR)/$(APP_NAME)

.PHONY: go


go:
	go run .
