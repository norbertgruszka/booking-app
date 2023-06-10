#!/bin/bash

KIND_CLUSTER_NAME=booking-app

# Create Kind cluster
kind create cluster --config ../kind-config.yml --name $KIND_CLUSTER_NAME

# Install CNI - Cilium
helm upgrade --install \
  --namespace cilium \
  --create-namespace \
  --kube-context kind-$KIND_CLUSTER_NAME \
  --repo https://helm.cilium.io/ \
  cilium cilium \
  --values - <<EOF

EOF

# Install and configure Metallb
KIND_NET_CIDR=$(docker network inspect kind -f '{{(index .IPAM.Config 0).Subnet}}')
METALLB_IP_START=$(echo ${KIND_NET_CIDR} | sed "s@0.0/16@255.200@")
METALLB_IP_END=$(echo ${KIND_NET_CIDR} | sed "s@0.0/16@255.250@")
METALLB_IP_RANGE="${METALLB_IP_START}-${METALLB_IP_END}"
helm upgrade --install \
  --namespace metallb-system \
  --create-namespace \
  --kube-context kind-$KIND_CLUSTER_NAME \
  --repo https://metallb.github.io/metallb \
  metallb metallb \
  --values - <<EOF
configInline:
  address-pools:
  - name: default
    protocol: layer2       # use layer2 protocol
    addresses:
    - ${METALLB_IP_RANGE}  # configure the ip address range
EOF
