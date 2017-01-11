# kubeps

[![Build Status](https://travis-ci.org/koudaiii/kubeps.svg?branch=master)](https://travis-ci.org/koudaiii/kubeps)
[![Docker Repository on Quay](https://quay.io/repository/koudaiii/kubeps/status "Docker Repository on Quay")](https://quay.io/repository/koudaiii/kubeps)
[![GitHub release](https://img.shields.io/github/release/koudaiii/kubeps.svg)](https://github.com/koudaiii/kubeps/releases)

Get container image tag for Kubernetes Pods

As you know, `kubectl get pod -o wide --show-labels` can get only pod( NAME,READY,STATUS, RESTARTS,AGE,IP,NODE,LABELS).
`kubectl get pod` difficult for you to get container image or tag.
`kubeps` enables you to get container image and tag in ALL pods that the specified namespace or labels.

![example](_images/example.png)

## Table of Contents

* [Requirements](#requirements)
* [Installation](#installation)
  + [Using Homebrew (OS X only)](#using-homebrew-os-x-only)
  + [Precompiled binary](#precompiled-binary)
  + [From source](#from-source)
  + [Run in a Docker container](#run-in-a-docker-container)
* [Usage](#usage)
  + [kubeconfig file](#kubeconfig-file)
  + [Options](#options)
* [Development](#development)
* [Author](#author)
* [License](#license)

## Requirements

Kubernetes 1.3 or above

## Installation

### Using Homebrew (OS X only)

Formula is available at [koudaiii/homebrew-tools](https://github.com/koudaiii/homebrew-tools).

```bash
$ brew tap koudaiii/tools
$ brew install kubeps
```

### Precompiled binary

Precompiled binaries for Windows, OS X, Linux are available at [Releases](https://github.com/koudaiii/kubeps/releases).

### From source

```bash
$ go get -d github.com/koudaiii/kubeps
$ cd $GOPATH/src/github.com/koudaiii/kubeps
$ make deps
$ make install
```

### Run in a Docker container

docker image is available at [quay.io/koudaiii/kubeps](https://quay.io/repository/koudaiii/kubeps).

```bash
# -t is required to colorize logs
$ docker run \
    --rm \
    -t \
    -v $HOME/.kube/config:/.kube/config \
    quay.io/koudaiii/kubeps:latest \
      -kubeconfig=/.kube/config
```

## Usage

`kubeps` gets all containers in pod in the specified namespace or labels.

```bash
$ kubeps --namespace docker-hello-world
=== Deployment ===
NAME                    IMAGE                           NAMESPACE
docker-hello-world      koudaiii/hello-world:3959aca    docker-hello-world

=== Pod ===
NAME                                    IMAGE                           STATUS  RESTARTS        START                           NAMESPACE
docker-hello-world-2473057991-9ftt0     koudaiii/hello-world:3959aca    Running 0               2016-11-18 14:27:21 +0900 JST   docker-hello-world
docker-hello-world-2473057991-biyvx     koudaiii/hello-world:3959aca    Running 0               2016-12-25 20:03:34 +0900 JST   docker-hello-world
docker-hello-world-2473057991-dkkv1     koudaiii/hello-world:3959aca    Running 0               2016-12-25 20:03:34 +0900 JST   docker-hello-world
docker-hello-world-2473057991-qtpu7     koudaiii/hello-world:3959aca    Running 0               2016-12-25 20:03:34 +0900 JST   docker-hello-world
docker-hello-world-2473057991-w1st3     koudaiii/hello-world:3959aca    Running 0               2016-12-25 20:03:34 +0900 JST   docker-hello-world
```

With `--labels` option, you can filter pods.


```bash
$ kubeps --labels role=web --namespace docker-hello-world
=== Deployment ===
NAME                    IMAGE                           NAMESPACE
docker-hello-world      koudaiii/hello-world:3959aca    docker-hello-world

=== Pod ===
NAME                                    IMAGE                           STATUS  RESTARTS        START                           NAMESPACE
docker-hello-world-2473057991-9ftt0     koudaiii/hello-world:3959aca    Running 0               2016-11-18 14:27:21 +0900 JST   docker-hello-world
docker-hello-world-2473057991-biyvx     koudaiii/hello-world:3959aca    Running 0               2016-12-25 20:03:34 +0900 JST   docker-hello-world
docker-hello-world-2473057991-dkkv1     koudaiii/hello-world:3959aca    Running 0               2016-12-25 20:03:34 +0900 JST   docker-hello-world
docker-hello-world-2473057991-qtpu7     koudaiii/hello-world:3959aca    Running 0               2016-12-25 20:03:34 +0900 JST   docker-hello-world
docker-hello-world-2473057991-w1st3     koudaiii/hello-world:3959aca    Running 0               2016-12-25 20:03:34 +0900 JST   docker-hello-world
```

### kubeconfig file

`kubeps` uses `~/.kube/config` as default.
You can specify another path by `KUBECONFIG` environment variable or `--kubeconfig` option.
`--kubeconfig` option always overrides `KUBECONFIG` environment variable.

```bash
$ KUBECONFIG=/path/to/kubeconfig kubeps
# or
$ kubeps --kubeconfig=/path/to/kubeconfig
```

### Options

|Option|Description|Required|Default|
|---------|-----------|-------|-------|
|`--kubeconfig=KUBECONFIG`|Path of kubeconfig||`~/.kube/config`|
|`--labels=LABELS`|Label filter query (e.g. `app=APP,role=ROLE`)|||
|`--namespace=NAMESPACE`|Kubernetes namespace||All namespaces|
|`-h`, `-help`|Print command line usage|||
|`-v`, `-version`|Print version|||

## Development

Clone this repository and build using `make`.

```bash
$ go get -d github.com/koudaiii/kubeps
$ cd $GOPATH/src/github.com/koudaiii/kubeps
$ make
```

## Author

[@koudaiii](https://github.com/koudaiii)

## License

[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)
