FROM golang

ENV GOBIN $GOPATH/bin
ENV PROJECT_DIR /go/src/github.com/geneva_horodateur

ADD vendor ${PROJECT_DIR}/vendor

ADD main.go ${PROJECT_DIR}/
ADD myservice.cert ${PROJECT_DIR}/
ADD myservice.key ${PROJECT_DIR}/
ADD mockup ${PROJECT_DIR}/mockup

EXPOSE 8080
