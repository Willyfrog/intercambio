BINARY_NAME=intercambio

build:
	GOARCH=amd64 GOOS=darwin go build -o ./out/${BINARY_NAME} main.go

run:
	./out/${BINARY_NAME}

clean:
	go clean
	rm ./out/${BINARY_NAME}

dep:
	go mod download

vet:
	go vet

lint:
	golangci-lint run --enable-all