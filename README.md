# drone-teams

Drone plugin to send teams notifications for build status

## Build

Build the binary with the following command:

```console
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
export GO111MODULE=on

go build -a -tags netgo -o release/linux/amd64/drone-download ./cmd/drone-teams
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
  -e DRONE_BUILD_STATUS=Failed \
  -e DRONE_BUILD_ACTION=test \
  -e DRONE_REPO_NAME=test \
  -e DRONE_COMMIT_AUTHOR=test \
  -e DRONE_COMMIT_MESSAGE=test \
  -e DRONE_COMMIT_LINK=test \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  jdamata/drone-teams
```
