BIN:=bin
PROTOC:=protoc

.PHONY: host plugin

all: greeter_host greeter_plugin

greeter_host greeter_plugin:
	mkdir -p $(BIN)
	go build -v -o $(BIN)/$@ cmd/$@/main.go

run:
	$(BIN)/greeter_host $(BIN)/greeter_plugin

generate:
	$(PROTOC) --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative cmd/pluginproto/plugin.proto cmd/hostproto/host.proto



