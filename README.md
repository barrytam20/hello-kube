# Kubernetes Steps
```
minikube start
minikube dashboard

// set up terminal env to use minikube docker env
eval $(minikube docker-env)

// rebuild docker image
docker build -f hello-api/Dockerfile -t hello-api:k8 .

// create k8s deployment
kubectl create -f hello-api-deployments.yaml

// create k8s service
kubectl expose deployment hello-api --type=LoadBalancer --port=5000
kubectl get services
```