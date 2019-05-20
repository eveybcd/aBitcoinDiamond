

CLI := build/aBcd-cli
SRC_CLI := github.com/aBitcoinDiamond/cmd/btcctl
APP := build/aBcdd


build:
	@go build $(BUILD_FLAGS) -v -i -o $(APP)
	@go build -v -i -o $(CLI) $(SRC_CLI)

build_ci: depends ## Build the binary file for CI
	@go build -v -i -o $(CLI) $(SRC_CLI)
	@go build $(BUILD_FLAGS) -v -o $(APP)

