![Build Status](https://github.com/cloudacademy/web-echo-app/actions/workflows/go.yml/badge.svg) 
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/cloudacademy/web-echo-app)

# Web Echo App
A simple web based application that prints a message on a coloured background, both of which can be configured using environment variables.

## Usage
To start the web application, configure both the `MESSAGE` and `BACKGROUND_COLOR` environment variables.

Startup:
```
MESSAGE=CloudAcademy BACKGROUND_COLOR=yellow ./webapp
```

## Docker
The web application has been packaged into a Docker image. The Docker image can be pulled with the following command:

```
docker pull cloudacademydevops/webappecho:v2
```

Use the following command to launch the web echoing application within Docker:
```
docker run --name webapp --env MESSAGE=CloudAcademy --env BACKGROUND_COLOR=yellow -p 8080:80 --detach cloudacademydevops/webappecho:v2
```

## Kubernetes
Use the following command to launch the web echoing application as a Deployment resource within a cluster:

```
cat << EOF | kubectl apply -f -
apiVersion: apps/v1
kind: Deployment
metadata:
  name: webapp
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
        image: cloudacademydevops/webappecho:v2
        imagePullPolicy: IfNotPresent
        command: ["/go/bin/demo"]
        ports:
        - containerPort: 80
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
docker buildx build --platform=linux/amd64 -t cloudacademydevops/webappecho:v2 .
```