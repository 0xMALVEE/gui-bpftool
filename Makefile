APP_NAME = tui-bpftool
BUILD_DIR = build/bin
EXECUTABLE = $(BUILD_DIR)/$(APP_NAME)

.PHONY: go


go:
	go build . && sudo ./tui-bpftool
