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
  "message": "pong:
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
