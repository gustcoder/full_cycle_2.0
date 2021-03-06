### Obter Pods
```
kubectl get pods
kubectl get po
```

### Info sobre Pods
```
kubectl describe pod pod_name
```

### Deletando Pods
```
kubectl delete pod pod_name
```

### Obter ReplicaSets
```
kubectl get replicasets
```

### Deletando ReplicaSets
```
kubectl delete replicasets replicaset_name
```

### Criando um Pod / ReplicaSet / Deployment
```
kubectl apply -f ./pod.yaml
kubectl apply -f ./replicaset.yaml
kubectl apply -f ./deployment.yaml
```

## Deployment
#### Deployments servem para que replicasets sejam recriados quando a versão da imagem dos containers for alterada.
#### Ao rodar um novo Deployment, será gerada um novo replicaset, mas o anterior será mantido (porém com 0 pods).

### Rollback de Deployments
```
kubectl rollout undo deployment deployment_name
kubectl rollout undo deployment deployment_name --to-revision=
```
#### Ao fazer rollout em um Deployment, se já existir uma replicaset da versão anterior que foi mantida anteriormente ela será reaproveitada, e o deployment atual que ficará com 0 pods desta vez. Bem legal!