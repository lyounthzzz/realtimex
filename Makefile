GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
API_PROTO_FILES=$(shell find api -name *.proto)
BUF_INSTALLED := $(shell command -v buf 2> /dev/null)

ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	INTERNAL_PROTO_FILES=$(shell $(Git_Bash) -c "find internal -name *.proto")
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
else
	INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
	API_PROTO_FILES=$(shell find api -name *.proto)
endif

.PHONY: errors
# generate errors code
errors:
	protoc --proto_path=. \
               --proto_path=./third_party \
               --go_out=paths=source_relative:. \
               --go-errors_out=paths=source_relative:. \
               $(API_PROTO_FILES)

.PHONY: validate
# generate validate code
validate:
	protoc --proto_path=. \
               --proto_path=./third_party \
               --go_out=paths=source_relative:. \
               --validate_out=paths=source_relative,lang=go:. \
               $(API_PROTO_FILES)

.PHONY: config
# generate internal proto
config:
	protoc --proto_path=. \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:. \
	       $(INTERNAL_PROTO_FILES)

.PHONY: api
# generate api proto
api:
	protoc --proto_path=. \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:. \
 	       --go-http_out=paths=source_relative:. \
 	       --go-grpc_out=paths=source_relative:. \
		   --openapiv2_out=. \
	       $(API_PROTO_FILES)

.PHONY: build
# build
build:
	mkdir -p bin/ && GOPROXY="https://goproxy.cn,direct" GOPRIVATE="gitlab.yc345.tv/*" go build -ldflags '-w -s -extldflags "-static" -X main.Version=$(VERSION)' -tags musl -o ./bin/ ./...

.PHONY: generate
# generate
generate:
	go generate ./...

.PHONY: lint
# golangci-lint
lint:
	go mod tidy
	GOGC=20 golangci-lint run ./...

.PHONY: swagger
# proto -> swagger.json
swagger:
	protoc --proto_path=. \
        --proto_path=./third_party \
        --openapiv2_out . \
        --openapiv2_opt logtostderr=true \
        --openapiv2_opt json_names_for_fields=false \
        $(API_PROTO_FILES)

.PHONY: wire
# wire
wire:
	wire ./...

.PHONY: gosec
# gosec 代码安全检查
gosec:
	GO111MODULE=on go install github.com/securego/gosec/v2/cmd/gosec@latest
	gosec -quiet -exclude=G104,G108,G403,G501,G502 -exclude-dir=api,third_party,sql,test ./...

.PHONY: checkVersion
# 定义变量; 要检查更新的包
PACKAGES := gitlab.yc345.tv/backend/utils/v2 gitlab.yc345.tv/backend/go-logger
# checkVersion 检查包版本并是否更新
checkVersion:
	@for package in $(PACKAGES); do \
		result="$$(go list -m -u $$package)"; \
		if [ -z "$${result##*[[]*}" ]; then \
			content="$$(echo "$$result" | grep -oE '\[([^]]+)\]' | sed 's/[][]//g')"; \
			echo "\033[1;33m$$package 已存在最新版本: $$content\033[0m"; \
			echo "\033[31mTips: 升级包有可能出现依赖包不兼容的情况 \033[0m"; \
			read -p "是否更新到最新版本 (y/N): " choice && \
			if [ "$$choice" = "y" ]; then \
				go get $$package && \
				upresult="$$(go list -m -u $$package)"; \
				if [ -z "$${upresult##*[[]*}" ]; then \
					echo "\033[1;31m$$package 更新失败!\033[0m"; \
				fi; \
				go mod tidy; \
			fi; \
		fi; \
	done

.PHONY: forceVersion
# forceVersion 强制检查包版本并更新 仅用于ci测试阶段
forceVersion:
	@for package in $(PACKAGES); do \
		result="$$(go list -m -u $$package)"; \
		if [ -z "$${result##*[[]*}" ]; then \
			content="$$(echo "$$result" | grep -oE '\[([^]]+)\]' | sed 's/[][]//g')"; \
			echo "\033[1;33m$$package 已存在最新版本: $$content\033[0m"; \
			echo "\033[31mTips: 升级包有可能出现依赖包不兼容的情况 \033[0m"; \
			go get $$package && \
			upresult="$$(go list -m -u $$package)"; \
			if [ -z "$${upresult##*[[]*}" ]; then \
				echo "\033[1;31m$$package 更新失败!\033[0m"; \
			fi; \
			go mod tidy; \
		fi; \
	done

.PHONY: buf
# buf 格式化 proto
buf:
	@if [ -n "$(BUF_INSTALLED)" ]; then \
        cd ./api  && \
        buf format -w && \
        echo "proto format finish"; \
    else \
        echo "please installation buf: https://buf.build/docs/installation"; \
    fi

.PHONY: gotest
# gotest
gotest:
	GO_ENV=development go test -race -cover -coverprofile=test-coverage.txt ./...

.PHONY: all
# generate all
all:
	make api;
	make errors;
	make validate;
	make config;
	go mod tidy;
	make checkVersion;
#	make generate;
	make swagger;
	make wire;
	make build;

deps:
	pip3 install  pre-commit

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
