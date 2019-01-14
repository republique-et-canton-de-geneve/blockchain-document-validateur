FROM golang as builder
MAINTAINER Sylvain Laurent
ENV GOBIN $GOPATH/bin
ENV PROJECT_DIR github.com/Magicking/rc-ge-validator
ENV PROJECT_NAME r-c-g-horodatage-validateur-server
ADD vendor /usr/local/go/src
ADD cmd /go/src/${PROJECT_DIR}/cmd
ADD models /go/src/${PROJECT_DIR}/models
ADD restapi /go/src/${PROJECT_DIR}/restapi
ADD merkle /go/src/${PROJECT_DIR}/merkle
ADD internal /go/src/${PROJECT_DIR}/internal
WORKDIR /go/src/${PROJECT_DIR}
RUN go build -v -o /go/bin/main /go/src/${PROJECT_DIR}/cmd/${PROJECT_NAME}/main.go

FROM alpine:latest
RUN apk --no-cache add libc6-compat ca-certificates
COPY --from=builder /go/bin/main /go/bin/main
EXPOSE 8090
CMD ["/go/bin/main","--host","0.0.0.0","--port=8090"]

