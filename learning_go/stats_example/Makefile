BINARY_NAME=anscombe-quartet

build:
	GOARCH=amd64 GOOS=darwin go build -o ./bin/${BINARY_NAME}-darwin stats_example.go
	GOARCH=amd64 GOOS=linux go build -o ./bin/${BINARY_NAME}-linux stats_example.go
	GOARCH=amd64 GOOS=windows go build -o ./bin/${BINARY_NAME}-windows stats_example.go

run_osx: build
	./bin/${BINARY_NAME}-darwin

clean:
	go clean
	rm ./bin/${BINARY_NAME}-darwin
	rm ./bin/${BINARY_NAME}-linux
	rm ./bin/${BINARY_NAME}-windows

vet:
	go vet

