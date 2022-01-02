## Docker Golang

https://hub.docker.com/_/golang

https://hub.docker.com/repository/docker/gustcoder/go-grpc

## GRPC

https://grpc.io/docs/languages/go/quickstart/
https://grpc.io/docs/protoc-installation/

### Recursos necessários para o ambiente Go (já estão inclusos no Dockerfile)

```
apt-get install -y protobuf-compiler && \
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26 && \
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1 && \
```

### Criar estrutura GRPC
```
protoc --proto_path=proto proto/*.proto --go_out=pb --go-grpc_out=pb
```

### Instalar pacote do Go para conseguir criar o Server.go

https://pkg.go.dev/google.golang.org/grpc#section-readme

```
go get -u google.golang.org/grpc
```

### Instalar Evans (GRPC Client)

```
go install github.com/ktr0731/evans@latest
evans -r repl --host localhost --port 50051
```