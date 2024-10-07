dev: 
	go run ./cmd/*

genProtobuf:
	protoc --go_out=./internal/store/grpc/gen \
		--go-grpc_out=./internal/store/grpc/gen \
		--go_opt=paths=source_relative \
		--go-grpc_opt=paths=source_relative \
		--proto_path=./internal/store/grpc/proto plugBack.proto
