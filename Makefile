build:
		go build

protoc:
		protoc  todo/transport/gRPC/todo.proto --go_out=todo/transport/gRPC --proto_path=/usr/local/include --proto_path=todo/transport/gRPC

swagger:
		swagger generate spec -o ./swagger.json --scan-models && swagger serve -F=swagger swagger.json