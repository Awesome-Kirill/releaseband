.PHONY: build-src build test test-cov lint

NAME=api
LD_FLAGS?="-X ${VERSION_PACKAGE_PATH}.commitHash=${COMMIT_HASH} -X ${VERSION_PACKAGE_PATH}.version=${VERSION}"

build-src:
	go build -o bin/${NAME} cmd/${NAME}.go

build: build-src
	SERVER_LISTEN_PORT=8000 \
	./bin/${NAME}

test:
	go test  ./... --tags=tests

test-cov:
	go test --tags=tests -coverpkg=./internal/... -coverprofile=coverage.txt ./...
	go tool cover -func coverage.txt
	rm coverage.txt

lint:
	golangci-lint run --fix

