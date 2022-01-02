FROM golang:1.17

WORKDIR /go/src/app

COPY . .

RUN apt-get update && apt-get install protobuf-compiler -y

RUN apt-get update && apt-get install -y protobuf-compiler && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26 && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1 && \
    export PATH="$PATH:$(go env GOPATH)/bin"

CMD tail -f /dev/null