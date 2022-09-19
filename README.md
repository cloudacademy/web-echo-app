![Build Status](https://github.com/cloudacademy/web-echo-app/actions/workflows/go.yml/badge.svg) 
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/cloudacademy/web-echo-app)

# Web Echo App
A simple web based application that prints a message on a coloured background, both of which can be configured using environment variables.

## Usage
To start the web application, configure the following environment variables:
 - `HOSTPORT=0.0.0.0:8080`
 - `MESSAGE="CloudAcademy ❤ DevOps"`
 - `BACKGROUND_COLOR=yello`

Startup:
```
HOSTPORT=0.0.0.0:8080 MESSAGE="CloudAcademy ❤ DevOps" BACKGROUND_COLOR=yellow ./webapp
```
![webapp](./docs/webapp.png)

## Docker
The web application has been packaged into a Docker image. The Docker image can be pulled with the following command:

```
docker pull cloudacademydevops/webappecho:v3
```

Use the following command to launch the web echoing application within Docker:
```
docker run --name webapp --env MESSAGE=CloudAcademy --env BACKGROUND_COLOR=yellow -p 8080:8080 --detach cloudacademydevops/webappecho:v3
```

## Kubernetes
Use the following command to launch the web echoing application as a Deployment resource within a cluster:

```
cat << EOF | kubectl apply -f -
apiVersion: apps/v1
kind: Deployment
metadata:
  name: frontend-v1
  namespace: webapp
  labels:
    role: frontend
    version: v1
spec:
  replicas: 2
  selector:
    matchLabels:
      role: frontend
      version: v1
  template:
    metadata:
      labels:
        role: frontend
        version: v1
    spec:
      containers:
      - name: webapp
        image: cloudacademydevops/webappecho:v3
        imagePullPolicy: IfNotPresent
        command: ["/go/bin/webapp"]
        ports:
        - containerPort: 8080
        env:
        - name: MESSAGE
          valueFrom:
            configMapKeyRef:
              name: webapp-cfg-v1
              key: message
        - name: BACKGROUND_COLOR
          valueFrom:
            configMapKeyRef:
              name: webapp-cfg-v1
              key: bgcolor
EOF
```

## Build
The following commands can be used to build and package the source code:

Current operating system:
```
go build .
```

Linux operating system:
```
CGO_ENABLED=0 GOOS=linux go build -o webapp .
```

Docker:
```
docker buildx build --platform=linux/amd64 -t cloudacademydevops/webappecho:v3 .
```