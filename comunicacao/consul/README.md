### Criando um server consul

```
consul agent -dev
```

### Consultando membros

```
consul members
```

### Acessando catálogo Consul (datacenter / gossip protocol)

```
curl localhost:8500/v1/catalog/nodes
```

### Consultas ao servidor DNS - Dig

```
apk -U add bind-tools
dig @localhost -p 8600
```