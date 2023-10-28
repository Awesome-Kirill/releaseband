FROM golang:1.21-alpine as builder
WORKDIR /build
COPY go.mod .
RUN go mod download
COPY . .
RUN go build -v -ldflags="-s -w" -o /api cmd/api.go

FROM alpine:3
COPY --from=builder api /bin/api
ENTRYPOINT ["/bin/api"]
