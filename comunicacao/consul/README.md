### Criando um server consul

```
consul agent -dev
```

### Consultando membros

```
consul members
```

### Acessando catálogo Consul (datacenter / gossip protocol)

https://www.consul.io/api-docs/catalog

```
curl localhost:8500/v1/catalog/nodes
curl localhost:8500/v1/catalog/services

Outras formas:

consul catalog nodes -service nginx
consul catalog nodes -detailed
```

### Consultas ao servidor DNS - Dig

```
apk -U add bind-tools
dig @localhost -p 8600
dig @localhost -p 8600 consul01.node.consul
dig @localhost -p 8600 nginx.service.consul
```

### Comando para subir consul server
```
consul agent -server -bootstrap-expect=3 -node=consulserver01 -bind=172.25.0.3 -data-dir=/var/lib/consul -config-dir=/etc/consul.d

*** Necessário criar os diretórios /etc/consul.d e /var/lib/consul com mkdir
```

### Comando para subir consul client
```
consul agent -bind=172.25.0.5 -data-dir=/var/lib/consul -config-dir=/etc/consul.d

Conectar automaticamente em um server (join):
-retry-join=IP_DE_UM_SERVER

Podemos passar mais de um:
-retry-join=IP_DE_UM_SERVER -retry-join=IP_DE_OUTRO_SERVER

*** consul agent client ou consul agent (ocultando ele já sabe que é client)
```

### Informações

```
-data-dir: onde ficam os arquivos do servidor consul
-config-dir: diretório onde tem os arquivos de configuração do nosso consul (.json / .hcl (hashicorp configuration language))
```

### Fazendo join entre consulserver para se enxergarem no cluster
```
consul join 172.25.0.2
```
#### onde 172.25.0.2 é o IP do server consul ao qual se quer fazer o join

### Reload no arquivo services.json
```
consul reload
```

### Subindo nginx para testar
```
apk add nginx
mkdir /run/nginx # pode ser que já exista
nginx # para rodar o nginx
ps #listar processos
```

### Configurando nginx para não retornar 404 para tudo
```
apk add vim # para editar os arquivos
mkdir /usr/share/nginx/html -p # -p criar pastas recursivamente
vim /etc/nginx/conf.d/default.conf # remover o bloco do 404 e colocar root /usr/share/nginx/html;
vim /usr/share/nginx/html/index.html # inserir qualquer texto
nginx -s reload # reinicia servico nginx
curl localhost # deverá exibir novo conteúdo

consul reload
```

### Para subir baseado no server.jon

```
consul agent -config-dir=/etc/consul.d
```
