# Simple HTTP Server

A Simple Http server for devops demo.

## Usage

```sh
Upload Server

Usage:
  demo-devops [flags]
  demo-devops [command]

Available Commands:
  help        Help about any command
  server      Start server

Flags:
      --config string      config file (default is $HOME/.demo-devops.yaml)
      --debug              Set log level to Debug
      --disabletrace       Disable the trace
  -h, --help               help for demo-devops
      --jaegerurl string   Set jaegger agent endpoint (without port, without http://)
  -t, --toggle             Help message for toggle
      --version            version
```

For example in order to launch the server without opentracing, and store img in /tmp

```sh
demo-devops server --outputdir /tmp --disabletrace
```


## Deployment


```sh
kubectl apply -f deployment/deployment.yml
kubectl apply -f deployment/svc.yml
```

## Build

Have a look on the `Makefile`, you can generate docker, binary, or artifact for other plateform.

### Build Docker

```sh
make build
```

### Build locally

```sh
make build_local
```
