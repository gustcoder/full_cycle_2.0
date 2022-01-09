### Criando o modulo

```
go mod init graphql
```

https://gqlgen.com/getting-started/

### Adicionando gqlgen

```
go get github.com/99designs/gqlgen
```

### Gerando modelos de exemplo (ou caso já exista no projeto algum graph/schema.graphqls, irá criar estrutura baseada nele)
```
go run github.com/99designs/gqlgen init
```

### Iniciar servidor

```
go run server.go
```