[![license](http://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/udhos/prime/blob/main/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/udhos/prime)](https://goreportcard.com/report/github.com/udhos/prime)
[![Go Reference](https://pkg.go.dev/badge/github.com/udhos/prime.svg)](https://pkg.go.dev/github.com/udhos/prime)
[![Artifact Hub](https://img.shields.io/endpoint?url=https://artifacthub.io/badge/repository/prime)](https://artifacthub.io/packages/search?repo=prime)
[![Docker Pulls](https://img.shields.io/docker/pulls/udhos/prime)](https://hub.docker.com/r/udhos/prime)

# prime

## Build

```
./build.sh
```

## Usage

Start the server.

```
$ prime
```

Get the 100.000.000th prime.

```
$ time curl localhost:8080/prime/100000000
2038074743

real	0m1.300s
user	0m0.006s
sys	0m0.005s
```

## Docker

Docker hub:

https://hub.docker.com/r/udhos/prime

Run from docker hub:

```
docker run -p 8080:8080 --rm udhos/prime:0.0.0
```

Build recipe:

```
./docker/build.sh

docker push udhos/prime:0.0.0
```

### Create

```
mkdir charts
cd charts
helm create prime
```

Then edit files.

### Lint

```
helm lint ./charts/prime --values charts/prime/values.yaml
```

### Test rendering chart templates locally

```
helm template prime ./charts/prime --values charts/prime/values.yaml
```

### Render templates at server

```
helm install prime ./charts/prime --values charts/prime/values.yaml --dry-run
```

### Generate files for a chart repository

A chart repository is an HTTP server that houses one or more packaged charts.
A chart repository is an HTTP server that houses an index.yaml file and optionally (*) some packaged charts.

(*) Optionally since the package charts could be hosted elsewhere and referenced by the index.yaml file.

    docs
    ├── index.yaml
    └── prime-0.1.0.tgz

See script [update-charts.sh](update-charts.sh):

    # generate chart package from source
    helm package ./charts/prime -d ./docs

    # regenerate the index from existing chart packages
    helm repo index ./docs --url https://udhos.github.io/prime/

### Install

```
helm install prime ./charts/prime --values charts/prime/values.yaml
```

### Upgrade

```
helm upgrade prime ./charts/prime --values charts/prime/values.yaml
```

### Uninstall

```
helm uninstall prime
```
