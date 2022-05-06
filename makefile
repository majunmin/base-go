.PHONY: gen
gen:
	protoc --proto_path=. \
		--proto_path=./third_party/ \
		--go_out=paths=source_relative:. \
		--go-grpc_out=./proto \
		--go-http_out=./proto \
		--openapi_out==paths=source_relative:. \
		proto/helloworld.proto



