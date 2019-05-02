FROM golang:1.12-alpine AS src

# Install git
RUN apk update && apk add --no-cache git ca-certificates && update-ca-certificates

# Copy Repository
WORKDIR /go/src/github.com/MaksRybalkin/TravisTest/
COPY . ./

# Build Go Binary
RUN set -ex; \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo \
    -ldflags "-s -w -X main.RELEASE=0.0.1 -X main.DATE=`date +"%Y-%m-%dT%H:%M:%SZ"` -X main.REPO=`git config --get remote.origin.url` -X main.COMMIT=`git rev-parse --short HEAD`" \
    -o ./travistest

# Final image, no source code
FROM alpine:latest

RUN apk update && apk add --no-cache ca-certificates && update-ca-certificates

COPY --from=src /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8080

WORKDIR /opt/
COPY --from=src /go/src/github.com/MaksRybalkin/TravisTest/travistest .

# Run Go Binary
CMD /opt/travistest