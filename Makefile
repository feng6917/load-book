GOPATH:=$(shell go env GOPATH)
API_PROTO_FILES=$(shell find api -name "*.proto")
VERSION=$(shell git describe --tags --always)

COMPILE_TARGET="./"

.PHONY: init
# init env
init:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.5.0
	go install github.com/envoyproxy/protoc-gen-validate@latest

.PHONY: grpc
# generate grpc code
grpc:
	protoc --proto_path=. \
		--proto_path=./third_party \
		--go_out=paths=source_relative:$(COMPILE_TARGET) \
		--go-grpc_out=paths=source_relative:$(COMPILE_TARGET) \
		$(API_PROTO_FILES)

.PHONY: http
# generate http code
http:
	protoc --proto_path=. \
		--proto_path=./third_party \
		--go_out=paths=source_relative:$(COMPILE_TARGET)  \
		--go-http_out=paths=source_relative:$(COMPILE_TARGET)  \
		$(API_PROTO_FILES)

.PHONY: validate
# generate validate code
validate:
	protoc --proto_path=. \
           --proto_path=./third_party \
           --go_out=paths=source_relative:$(COMPILE_TARGET)  \
           --validate_out=paths=source_relative,lang=go:$(COMPILE_TARGET) \
           $(API_PROTO_FILES)

.PHONY: config
# generate config proto
config:
	protoc --proto_path=. \
 	       --go_out=paths=source_relative:$(COMPILE_TARGET) ./internal/conf/conf.proto

.PHONY: swagger
# generate swagger file
swagger:
	protoc --proto_path=. \
		--proto_path=./third_party \
		--openapiv2_out ./swagger \
		--openapiv2_opt logtostderr=true \
		--openapiv2_opt json_names_for_fields=false \
		$(API_PROTO_FILES)

#.PHONY: build
# build
#build:
#	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

.PHONY: generate
# generate
generate:
	go mod tidy
	go get github.com/google/wire/cmd/wire@latest
	go generate ./...

.PHONY: all
# generate all
all:
	make grpc;
	make http;
	make validate;
	make config;
	make swagger;

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
