BINARY_NAME=starter

build:
	 GOARCH=amd64 GOOS=darwin go build -o ${BINARY_NAME}-darwin main.go
	 GOARCH=amd64 GOOS=linux go build -o ${BINARY_NAME}-linux main.go
	 GOARCH=amd64 GOOS=windows go build -o ${BINARY_NAME}-windows main.go

run: build
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}-darwin
	rm ${BINARY_NAME}-linux
	rm ${BINARY_NAME}-windows

fmt:
	go fmt

lint:
	golangci-lint run --enable-all

test:
	go test ./...

coverage:
	go test ./... -coverprofile=coverage.out

dep:
	go mod download

vet:
	go vet

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
