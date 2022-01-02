## Docker Golang

https://hub.docker.com/_/golang

## GRPC

https://grpc.io/docs/languages/go/quickstart/
https://grpc.io/docs/protoc-installation/

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