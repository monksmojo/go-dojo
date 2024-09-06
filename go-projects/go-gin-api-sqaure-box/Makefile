BUILD_DIR := bin
BINARY_NAME := matrix

#  make file command for linux os

spin-up-linux: clean-linux swag-init test build-linux run-linux ## Run all tests, then build and run

clean-linux: ## Clean up, i.e. remove build artifacts in linux
	@if [ -d "$(BUILD_DIR)" ]; then rm -rf "$(BUILD_DIR)"; fi
	@go mod tidy

build-linux:
	@set GO_ENABLED=0 GOARCH=amd64 GOOS=linux & go build -o "./$(BUILD_DIR)/$(BINARY_NAME)-linux" "./cmd/app/main.go"

run-linux: 
	./"$(BUILD_DIR)/$(BINARY_NAME)-linux"


# make file command for window os

spin-up-windows: clean-windows swag-init test test-windows build-windows run-windows ## Run all tests, then build and run

clean-windows: ## Clean up, i.e. remove build artifacts in windows
	@if exist ".\$(BUILD_DIR)\" rd /s /q ".\$(BUILD_DIR)\"
	@go mod tidy

build-windows: 
	@set CGO_ENABLED=0 GOARCH=amd64 GOOS=windows & go build -o "$(BUILD_DIR)\$(BINARY_NAME)-windows.exe" ".\cmd\app\main.go"

run-windows: 
	.\"$(BUILD_DIR)\$(BINARY_NAME)-windows.exe"


# make file command for macos
spin-up-macos: clean-macos swag-init test build-macos run-macos ## Run all tests, then build and run for macOS
build-macos: 
	@CGO_ENABLED=0 GOARCH=amd64 GOOS=darwin go build -o "$(BUILD_DIR)/$(BINARY_NAME)-macos" "./cmd/app/main.go"
clean-macos: 
	@rm -rf "$(BUILD_DIR)" || true 
	@go mod tidy
run-macos:
	"./$(BUILD_DIR)/$(BINARY_NAME)-macos"



.PHONY: test
test:
	go test ./... -v

.PHONY: swag-init
swag-init:
	swag init --parseDependency --parseInternal -g cmd/app/main.go 
