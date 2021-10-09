protos:
		protoc \
		 --proto_path=./common \
		 --proto_path=. \
		 --go-grpc_out=. \
		 --go_out=. \
		 ./recipe/*.proto