BIN:=bin
PROTOC:=../../grpc/protoc3/bin/protoc

.PHONY: host plugin

all: greeter_host greeter_plugin

greeter_host greeter_plugin:
	mkdir -p $(BIN)
	go build -v -o $(BIN)/$@ cmd/$@/main.go

run:
	$(BIN)/greeter_host $(BIN)/greeter_plugin


