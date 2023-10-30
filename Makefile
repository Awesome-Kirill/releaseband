.PHONY: build-src build test cover lint run stop run-app

NAME=api
LD_FLAGS?="-X ${VERSION_PACKAGE_PATH}.commitHash=${COMMIT_HASH} -X ${VERSION_PACKAGE_PATH}.version=${VERSION}"

build:
	go build -o bin/${NAME} cmd/${NAME}.go

run-app: build
	SERVER_LISTEN_PORT=8080 \
	./bin/${NAME}

run:
	 docker-compose up

down:
	 docker-compose down

test:
	go test  ./...

cover:
	go test -coverpkg=./internal/... -coverprofile=coverage.txt ./...
	go tool cover -func coverage.txt
	rm coverage.txt

lint:
	golangci-lint run --fix

