# Kubernetes Steps
```
minikube start
minikube dashboard
minikube addons enable ingress

// set up terminal env to use minikube docker env
eval $(minikube docker-env)

// rebuild docker image
cd hello-api
docker build -t hello-api:k8 .
cd ../hello-web
docker build -t hello-web:k8 .

// create k8s deployment
kubectl create -f hello-deployments.yml
kubectl expose deployment hello --type=NodePort

// create ingress
kubectl create -f ingress.yml
kubectl describe ing ingress

//route etc hosts to ingress
echo "$(minikube ip) hello.web hello.api" | sudo tee -a /etc/hosts

// create k8s service
kubectl get services
```
