# golang1.9 or latest

CHAIN33=github.com/33cn/chain33
CHAIN33_PATH=vendor/${CHAIN33}
plugin=github.com/33cn/plugin


PKG_LIST_VET := `go list ./... | grep -v "vendor" | grep -v plugin/dapp/evm/executor/vm/common/crypto/bn256`
PKG_LIST_INEFFASSIGN= `go list -f {{.Dir}} ./... | grep -v "vendor"`
.PHONY: default build

default: build

all: vendor build

build:
	go build -v -i -o build/aBcd
	go build -v -i -o build/aBcdm-cli github.com/aBitcoinDiamond/cmd/btcctl
	go build -v -i -o build/aBcds-cli github.com/aBitcoinDiamond/slave/cli
	@cp slaveChain.toml build


autotest: ## build autotest binary
	@cd build/autotest && bash ./build.sh && cd ../../
	@if [ -n "$(dapp)" ]; then \
		rm -rf build/autotest/local \
		&& cp -r $(CHAIN33_PATH)/build/autotest/local $(CHAIN33_PATH)/build/autotest/*.sh build/autotest/ \
		&& cd build/autotest && bash ./copy-autotest.sh local && cd local && bash ./local-autotest.sh $(dapp) && cd ../../../; fi



clean:
	@rm -rf build
	@go clean
