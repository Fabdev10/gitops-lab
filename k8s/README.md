# Kubernetes deployment guide

This project can be deployed on Kubernetes with either the raw manifests under k8s/manifests or the Helm chart under charts/blog-app.

## 1. Create a local cluster

### With k3d

```bash
k3d cluster create dev-cluster --agents 1 --port '8081:80@loadbalancer'
```

### With kind

```bash
kind create cluster --name dev-cluster
```

## 2. Install ingress controller and cert-manager

```bash
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.10.0/deploy/static/provider/cloud/deploy.yaml
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.14.2/cert-manager.yaml
```

Wait until the pods are ready:

```bash
kubectl get pods -n ingress-nginx
kubectl get pods -n cert-manager
```

## 3. Install the manifests

```bash
kubectl apply -f k8s/manifests/
```

## 4. Install the Helm chart

```bash
helm install blog-app ./charts/blog-app -f charts/blog-app/values-dev.yaml
```

For staging:

```bash
helm upgrade --install blog-app ./charts/blog-app -f charts/blog-app/values-staging.yaml
```

## 5. Verify the deployment

```bash
kubectl get pods
kubectl get svc
kubectl get ingress
kubectl describe ingress blog-app
```

## 6. Test the services

Add `example.local` to `/etc/hosts` or use a local DNS override:

```bash
echo "127.0.0.1 example.local" | sudo tee -a /etc/hosts
curl -k https://example.local/api/posts
curl -k https://example.local/api/comments
```

If the ingress is not ready yet, wait a few minutes and rerun the curls.

## 7. GitOps with Flux

This project includes a Flux GitOps configuration under `k8s/flux`.

### Bootstrap Flux

Install the Flux CLI and controllers:

```bash
flux install
```

### Apply the Flux GitOps manifests

```bash
kubectl apply -f k8s/flux/flux-system/gotk-components.yaml
```

### Verify deployment

```bash
flux get sources git
flux get kustomizations
flux get helmreleases -A
```

Flux will sync the `blog-app` Helm release from `charts/blog-app` using the repository `https://github.com/Fabdev10/gitops-lab.git` and deploy it into the `blog-app` namespace.
