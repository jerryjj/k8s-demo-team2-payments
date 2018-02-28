# Payments -service

This is a demo service which is used in the presentation and setup of
[Kubernetes - Happily in production](https://github.com/jerryjj/k8s-gke-deployment-tpl).

## Development

**Get dependencies**

```sh
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u google.golang.org/grpc
go get -u github.com/op/go-logging
```

### Generating new Protoc

```sh
./protoc.sh
```

### Building

```sh
./build.sh
```

**Building and Deploying to the GCP Project**

```sh
export PROJECT_ID=[YOUR_GCP_PROJECT_ID]
./cloud-build.sh
```

### Running

```sh
./payments-service
```
