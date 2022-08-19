# Example Golang Kubernetes Devthingy

## Prerequisites:

- skaffold
- ko
- kind

## Starting environment

```bash
$ kind create cluster --config hacking/kind.yaml
$ skaffold dev
$ curl -k https://example.127.0.0.1.nip.io
```
