# note: call scripts from /scripts

generate:
	protoc ./internal/api/*.proto \
			--go_out=. \
			--go_opt=paths=source_relative \
			--go-grpc_out=. \
			--go-grpc_opt=paths=source_relative \
            --proto_path=. \
            --proto_path=./third_party/proto \
	        --grpc-gateway_out . \
	        --grpc-gateway_opt logtostderr=true \
	        --grpc-gateway_opt paths=source_relative \
	        --grpc-gateway_opt generate_unbound_methods=true \
