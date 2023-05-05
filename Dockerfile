FROM golang:1.19 as builder

# run
RUN mkdir -p $GOPATH/src/github.com/husanmusa/mh_api_gateway
WORKDIR $GOPATH/src/github.com/husanmusa/mh_api_gateway
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.50.0

# Copy the local package files to the container's workspace.
COPY . ./

# installing depends and build
RUN export CGO_ENABLED=0 && \
    export GOOS=linux && \
    go mod vendor && \
#    make linter && \
    make build && \
    mv ./bin/mh_api_gateway/

FROM alpine

COPY --from=builder mh_api_gateway .
RUN apk update && apk add -U tzdata && cp /usr/share/zoneinfo/Asia/Tashkent /etc/localtime && apk del tzdata

ENTRYPOINT ["/mh_api_gateway"]
