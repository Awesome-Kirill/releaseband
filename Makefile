.PHONY: build-src build test cover lint run stop

NAME=api
LD_FLAGS?="-X ${VERSION_PACKAGE_PATH}.commitHash=${COMMIT_HASH} -X ${VERSION_PACKAGE_PATH}.version=${VERSION}"

build-src:
	go build -o bin/${NAME} cmd/${NAME}.go

build: build-src
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

