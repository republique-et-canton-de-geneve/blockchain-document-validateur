FROM golang

ENV GOBIN $GOPATH/bin
ENV PROJECT_DIR /go/src/github.com/geneva_horodateur

ADD vendor ${PROJECT_DIR}/vendor

ADD main.go ${PROJECT_DIR}/
ADD myservice.cert ${PROJECT_DIR}/
ADD myservice.key ${PROJECT_DIR}/
ADD mockup ${PROJECT_DIR}/mockup

WORKDIR ${PROJECT_DIR}

RUN go build -v -o /go/bin/main ${PROJECT_DIR}/main.go

CMD ["/go/bin/main"]

EXPOSE 8080
