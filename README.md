
AWS Services as a Service in your own Service

# Development Environment

Install minikube: 

[https://kubernetes.io/docs/tasks/tools/install-minikube/](https://kubernetes.io/docs/tasks/tools/install-minikube/)

Run it: 

`minikube start`

Save your access credentials locally: 

`kubectl config view > .kubeconfig`

Start CockroachDB: 

`cockroach start --http-port 9080 --insecure --background`
`cockroach sql -d "root@localhost:26257" -e "CREATE DATABASE altwebplatform"`

Start AWP: 

`go run main.go`

# Update Dependencies: 

`glide install`