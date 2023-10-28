FROM golang:1.21 as builder
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -v -ldflags="-s -w" -o /app cmd/api.go


FROM alpine:3.15
WORKDIR /usr/src/app
COPY --from=builder app /usr/local/bin/app
COPY site.txt .
ENTRYPOINT ["app"]
