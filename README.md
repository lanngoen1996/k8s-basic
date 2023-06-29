# k8s-basic

Example kubernetes basic

### Install Golang

1. Install Golang by Google "install golang"

2. Download installer according to your platform (Mac, Windows, Ubuntu, etc)

3. After install open terminal and run command

```cmd
$ go version

go version go1.20.5 darwin/arm64
```

### Run service go

1. Open directory ./service

2. After run command

```cmd
$ go run main.go

...
[GIN-debug] Listening and serving HTTP on :3333
```

3. Open browser http://localhost:3333/ping

```json
{
  "message": "pong"
}
```

### Build go image

1. Install docker from https://docs.docker.com/get-docker/

- After install open terminal and run command

```cmd
$ docker version

Client: Docker Engine - Community
 Version:           20.10.2
 API version:       1.41
 Go version:        go1.13.15
 Git commit:        2291f61
 Built:             Mon Dec 28 16:12:42 2020
 OS/Arch:           darwin/amd64
 Context:           default
 Experimental:      true
...
```

2. Open directory ./service

3. After run command

```cmd
$ docker build -f Dockerfile.prod -t go-service .

[+] Building 20.5s (17/17) FINISHED
 => [internal] load build definition from Dockerfile.prod                 0.0s
 => => transferring dockerfile: 337B                                      0.0s
 => [internal] load .dockerignore                                         0.0s
 => => transferring context: 2B                                           0.0s
 => [internal] load metadata for docker.io/library/golang:1.20.5-alpine   2.5s
 => [internal] load metadata for docker.io/library/alpine:latest          3.6s
 => [auth] library/golang:pull token for registry-1.docker.io             0.0s
 => [auth] library/alpine:pull token for registry-1.docker.io             0.0s
 => [internal] load build context                                         2.0s
...
```

4. Check images

```cmd
$ docker images

REPOSITORY                    TAG       IMAGE ID       CREATED          SIZE
go-service                    latest    68c89bb52600   20 minutes ago   19MB
...
```

### Install Kubernetes

1. Open Docker Preference > Kubernetes > Check box at Enable Kubernetes,
   wait until Kubernetes is running (If Kubernetes starting forever, try reset factory default)

2. After Kubernetes is running, Run command

```cmd
$ kubectl version

Client Version: version.Info{
  Major:"1",
  Minor:"25",
  GitVersion:"v1.25.9",
  GitCommit:"a1a87a0a2bcd605820920c6b0e618a8ab7d117d4",
  GitTreeState:"clean",
  BuildDate:"2023-04-12T12:16:51Z",
  GoVersion:"go1.19.8",
  Compiler:"gc",
  Platform:"darwin/arm64"
}
Server Version: version.Info{
  Major:"1",
  Minor:"25",
  GitVersion:"v1.25.9",
  GitCommit:"a1a87a0a2bcd605820920c6b0e618a8ab7d117d4",
  GitTreeState:"clean",
  BuildDate:"2023-04-12T12:08:36Z",
  GoVersion:"go1.19.8",
  Compiler:"gc",
  Platform:"linux/arm64"
}
```

### Install nginx ingress

1. Install nginx-ingress follow the step from this URL
   https://kubernetes.github.io/ingress-nginx/deploy

```cmd
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.8.0/deploy/static/provider/cloud/deploy.yaml
```

2. Run command
   You must see ingress-nginx in the list of namespace

```cmd
$ kubectl get ns

NAME              STATUS   AGE
default           Active   53m
ingress-nginx     Active   33s
...
```

3. Run command

```cmd
$ kubectl get po -n ingress-nginx

NAME                                        READY   STATUS      RESTARTS   AGE
ingress-nginx-admission-create-mlwmr        0/1     Completed   0          5m19s
ingress-nginx-admission-patch-fpk4v         0/1     Completed   0          5m19s
ingress-nginx-controller-56c75d774d-xwjhq   1/1     Running     0          5m19s
```

4. Run command

```cmd
$ kubectl get svc -n ingress-nginx

NAME              TYPE           CLUSTER-IP       EXTERNAL-IP   PORT(S)                      AGE
nginx             LoadBalancer   10.106.200.139   localhost     80:31482/TCP,443:31129/TCP   2m11s
nginx-admission   ClusterIP      10.102.57.197    <none>        443/TCP                      2m11s
```

### Deploy Service to Kubernetes

1. Run command from root directory

```cmd
$ kubectl apply -f ./k8s

namespace/basic-k8s created
deployment.apps/service created
service/service created
ingress.networking.k8s.io/ingress created
```

2. Run command

```cmd
$ kubectl get po -n basic-k8s

NAME                       READY   STATUS    RESTARTS   AGE
service-5fbfcfb745-xgcjr   1/1     Running   0          10m
```

- ** Default ingress URL endpoint is http://kubernetes.docker.internal **

2. Open browser http://kubernetes.docker.internal/ping (endpoint form service route)
