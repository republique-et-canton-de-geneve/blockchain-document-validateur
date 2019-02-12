FROM golang
MAINTAINER Sylvain Laurent

ENV GOBIN $GOPATH/bin
ENV PROJECT_DIR github.com/geneva_validateur/
ENV PROJECT_NAME r-c-g-horodatage-validateur-server

ADD vendor /usr/local/go/src
ADD cmd /go/src/${PROJECT_DIR}/cmd
ADD models /go/src/${PROJECT_DIR}/models
ADD restapi /go/src/${PROJECT_DIR}/restapi
ADD merkle /go/src/${PROJECT_DIR}/merkle
ADD internal /go/src/${PROJECT_DIR}/internal

WORKDIR /go/src/${PROJECT_DIR}

RUN go build -v -o /go/bin/main /go/src/${PROJECT_DIR}/cmd/${PROJECT_NAME}/main.go
ADD run.sh /go/src/${PROJECT_DIR}/
ENTRYPOINT /go/src/${PROJECT_DIR}/run.sh

EXPOSE 8090
