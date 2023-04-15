FROM golang:1.20 AS builder
ENV GO111MODULE=on

RUN mkdir -p /go/src/github.com/ettore83/projeto-anime/backend
WORKDIR /go/src/github.com/ettore83/projeto-anime/backend

COPY . .

RUN go mod download

# Compile
RUN GOOS=linux CGO_ENABLED=0 go build \
      -o api \
      ./cmd/api/

# Build 2
FROM alpine:3.15.0

RUN apk --no-cache --update add ca-certificates curl \
  && rm -rf /var/cache/apk

# Copy binary from builder
COPY --from=builder /go/src/github.com/ettore83/projeto-anime/backend/api /usr/bin/

ENTRYPOINT ["/usr/bin/api"]
