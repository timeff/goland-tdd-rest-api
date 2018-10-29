all: protoc swagger build

PROTOPATH = todo/transport/gRPC
PROTOFILE = todo.proto

build:
		go build

protoc:
		protoc  ${PROTOPATH}/${PROTOFILE} --go_out=${PROTOPATH} --proto_path=/usr/local/include --proto_path=${PROTOPATH}

swagger:
		swagger generate spec -o ./swagger.json --scan-models && swagger serve -F=swagger swagger.json