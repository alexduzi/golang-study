- build the image
```
docker build -t alexduzi/golang-deploy-k8s:latest -f Dockerfile.prod .
```

- verify the size of the generated image
```
docker images | grep golang
```

- run the generated image to test it
```
docker run --rm -p 8080:8080 alexduzi/golang-deploy-k8s:latest
```

- apply generated image to kind or mini-kube (k8s)
```
kubectl apply -f k8s/deployment.yaml

kubectl get pods
```

