BINARY_NAME=takeoff

build:
	GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME}-linux main.go
	GOARCH=amd64 GOOS=windows go build -o bin/${BINARY_NAME}-windows main.go

run: build
	go run bin/${BINARY_NAME}

clean: 
	go clean
	rm ./bin/${BINARY_NAME}-*