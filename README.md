# Example Golang Kubernetes Devthingy

## Prerequisites:

- [docker](https://www.docker.com/)
- [skaffold](https://skaffold.dev/)
- [ko](https://github.com/google/ko)
- [kind](https://kind.sigs.k8s.io/)
- [kustomize](https://kustomize.io/)

## Development environment

### Kubernetes cluster
I've included [kind](https://kind.sigs.k8s.io/) configuration for initializing a minimalistic
Kubernetes cluster which forwards host port 80 (`http`) and 443 (`https`) to the kubernetes cluster.

```bash
$Β kind create cluster --config hacking/kind.yaml
Creating cluster "kind" ...
 β Ensuring node image (kindest/node:v1.21.1) πΌ
 β Preparing nodes π¦  
 β Writing configuration π 
 β Starting control-plane πΉοΈ 
 β Installing CNI π 
 β Installing StorageClass πΎ 
Set kubectl context to "kind-kind"
You can now use your cluster with:

kubectl cluster-info --context kind-kind

Thanks for using kind! π
```

### Running example application
This example project utilizes [skaffold](https://skaffold.dev/), [ko](https://github.com/google/ko) and
[kustomize](https://kustomize.io/) to build and deploy the example application.

```bash
$ skaffold dev
$ curl -k https://example.127.0.0.1.nip.io
```

If you change the source code while [skaffold](https://skaffold.dev/) is running, it will rebuild and
redeploy the application for you.

### Debugging the application
I've included example vscode debugging support, allowing you to step through the execution
of a process running as a `Pod` in a remote Kubernetes cluster.

```bash
$ skaffold debug
```

You should now be able to set a breakpoint in the handler defined in  `pkg/service/service.go` using `VSCode`.
After setting up the breakpoint you can use the `Skaffold Debug` profile included in the repository. Hot reloading
won't work in skaffold debug mode.

