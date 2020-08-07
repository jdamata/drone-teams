FROM golang AS build-env

ENV CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go/src/github.com/jdamata/drone-teams
ADD . /go/src/github.com/jdamata/drone-teams
RUN go build -a -tags netgo -ldflags '-w' -o /bin/drone-teams

FROM alpine

RUN apk update && apk add git

COPY --from=build-env /bin/drone-teams /usr/bin/drone-teams