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
      --kubeconfig=/.kube/config
```

## Usage

`kubeps` gets all containers in pod in the specified namespace or labels.

```bash
Namespace:
Labels:

=== Deployment ===
NAME		IMAGE								NAMESPACE
kube-dns	gcr.io/google_containers/k8s-dns-kube-dns-amd64:1.14.4		kube-system
kube-dns	gcr.io/google_containers/k8s-dns-dnsmasq-nanny-amd64:1.14.4	kube-system
kube-dns	gcr.io/google_containers/k8s-dns-sidecar-amd64:1.14.4		kube-system
myapp		quay.io/koudaiii/myapp:latest					myapp
myapp		nginx:1.13.3-alpine						myapp
postgres	postgres:9.6.5							myapp

=== Pod ===
NAME							IMAGE								STATUS		READY	RESTARTS	START				LAST				NAMESPACE
kube-addon-manager-minikube				gcr.io/google-containers/kube-addon-manager:v6.4-beta.2		Running		1/1	1		2017-09-10 23:25:33 +0900 JST	<none>				kube-system
kube-dns-910330662-hkvmq				gcr.io/google_containers/k8s-dns-kube-dns-amd64:1.14.4		Running		3/3	1		2017-09-10 23:25:36 +0900 JST	<none>				kube-system
kube-dns-910330662-hkvmq				gcr.io/google_containers/k8s-dns-dnsmasq-nanny-amd64:1.14.4	Running		3/3	1		2017-09-10 23:25:36 +0900 JST	<none>				kube-system
kube-dns-910330662-hkvmq				gcr.io/google_containers/k8s-dns-sidecar-amd64:1.14.4		Running		3/3	1		2017-09-10 23:25:36 +0900 JST	<none>				kube-system
kubernetes-dashboard-2jl6t				gcr.io/google_containers/kubernetes-dashboard-amd64:v1.6.3	Running		1/1	1		2017-09-10 23:25:36 +0900 JST	<none>				kube-system
db-migrate-koudaiii-2017091116251505114716-707pt	quay.io/koudaiii/myapp:946bd19					Succeeded	0/1	0		2017-09-11 16:16:56 +0900 JST	2017-09-11 16:29:01 +0900 JST	myapp
myapp-2136627869-2qlm1					quay.io/koudaiii/myapp:latest					Running		2/2	1		2017-09-11 18:07:43 +0900 JST	<none>				myapp
myapp-2136627869-2qlm1					nginx:1.13.3-alpine						Running		2/2	1		2017-09-11 18:07:43 +0900 JST	<none>				myapp
myapp-2136627869-6w3mj					quay.io/koudaiii/myapp:latest					Running		2/2	1		2017-09-11 18:06:26 +0900 JST	<none>				myapp
myapp-2136627869-6w3mj					nginx:1.13.3-alpine						Running		2/2	1		2017-09-11 18:06:26 +0900 JST	<none>				myapp
myapp-2136627869-dcvw3					quay.io/koudaiii/myapp:latest					Running		2/2	1		2017-09-11 18:07:20 +0900 JST	<none>				myapp
myapp-2136627869-dcvw3					nginx:1.13.3-alpine						Running		2/2	1		2017-09-11 18:07:20 +0900 JST	<none>				myapp
myapp-2136627869-kwlhx					quay.io/koudaiii/myapp:latest					Running		2/2	1		2017-09-11 18:06:26 +0900 JST	<none>				myapp
myapp-2136627869-kwlhx					nginx:1.13.3-alpine						Running		2/2	1		2017-09-11 18:06:26 +0900 JST	<none>				myapp
myapp-job-1505118600-sr7wn				quay.io/koudaiii/myapp:latest					Succeeded	0/1	0		2017-09-11 17:30:05 +0900 JST	2017-09-11 17:30:15 +0900 JST	myapp
myapp-job-1505119500-dgc2d				quay.io/koudaiii/myapp:latest					Succeeded	0/1	0		2017-09-11 17:45:06 +0900 JST	2017-09-11 17:45:15 +0900 JST	myapp
myapp-job-1505120400-zbp9s				quay.io/koudaiii/myapp:latest					Succeeded	0/1	0		2017-09-11 18:00:08 +0900 JST	2017-09-11 18:01:08 +0900 JST	myapp
myapp-job-1505121300-cjgrt				quay.io/koudaiii/myapp:latest					Succeeded	0/1	0		2017-09-11 18:15:01 +0900 JST	2017-09-11 18:15:09 +0900 JST	myapp
myapp-job-1505122200-0pn60				quay.io/koudaiii/myapp:latest					Succeeded	0/1	0		2017-09-11 18:30:02 +0900 JST	2017-09-11 18:30:10 +0900 JST	myapp
postgres-2312165663-5vzcs				postgres:9.6.5							Running		1/1	1		2017-09-11 09:52:18 +0900 JST	<none>				myapp
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
|`-l`, `--labels=LABELS`|Label filter query (e.g. `app=APP,role=ROLE`)|||
|`-n`,`--namespace=NAMESPACE`|Kubernetes namespace||All namespaces|
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
