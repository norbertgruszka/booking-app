# booking-app
A simple microservice app using service mesh

## Deployment

### Create Kind cluster

Start by creating Kind cluster with Cilium as CNI.

```bash
kind create cluster --config ./kind-config.yml --name booking-app
```

### Install CNI

For that project I wanted to use Cilium as CNI. Execute following commands to install it on previously created Kind cluser

```bash
cilium install
helm upgrade --install --kube-context kind-booking-app --repo https://helm.cilium.io/ cilium cilium --version 1.13.3 --namespace kube-system --values ./helm/cilium.yml
```

### Install metallb

For easier access to cluster applications I deployed MetalLB. This step might be different for each machines, so please follow steps from [official documentation](https://kind.sigs.k8s.io/docs/user/loadbalancer/).

In general it all comes to executing following commands:

```bash
kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.13.7/config/manifests/metallb-native.yaml
kubectl apply -f https://kind.sigs.k8s.io/examples/loadbalancer/metallb-config.yaml
```


