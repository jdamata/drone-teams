[![Build Status](https://cloud.drone.io/api/badges/jdamata/drone-teams/status.svg)](https://cloud.drone.io/jdamata/drone-teams)
# drone-teams

Drone plugin to send teams notifications for build status

## Build

Build the binary with the following command:

```console
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
export GO111MODULE=on

go build -a -tags netgo -o release/linux/amd64/drone-teams ./cmd/drone-teams
```

## Docker

Build the Docker image with the following command:

```console
docker build \
  --label org.label-schema.build-date=$(date -u +"%Y-%m-%dT%H:%M:%SZ") \
  --label org.label-schema.vcs-ref=$(git rev-parse --short HEAD) \
  --file docker/Dockerfile.linux.amd64 --tag jdamata/drone-teams .
```

## Usage

```
docker run --rm \
  -e PLUGIN_WEBHOOK=<WEBHOOK ENDPOINT> \
  jdamata/drone-teams
```
