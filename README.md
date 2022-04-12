Proof of concept Kubernetes stack with tenancy.

Start by installing the [nginx ingress controller](https://kubernetes.github.io/ingress-nginx/deploy/#quick-start)
```
helm upgrade --install ingress-nginx ingress-nginx \
  --repo https://kubernetes.github.io/ingress-nginx \
  --namespace ingress-nginx --create-namespace
```

Each tenant has a tenant id. Any number of them can be deployed with

```
helm upgrade --install my-new-deploy-name ./chart --set tenant=tenant_id
```

The dashboard can be deployed with
```
kubectl apply -f ./kubernetes-dashboard-ingress.yaml
```