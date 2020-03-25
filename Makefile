build:
	go build -o hermes_server src/main.go

run:
	go run -race src/main.go

clean:
	go clean

build_idl:
	protoc --go_out=plugins=grpc:. ./idl/hermes.proto
