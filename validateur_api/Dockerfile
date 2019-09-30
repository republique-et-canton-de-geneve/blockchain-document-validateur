FROM golang:alpine AS build-env


ENV GOBIN $GOPATH/bin
ENV PROJECT_DIR github.com/geneva_validateur/
ENV PROJECT_NAME r-c-g-horodatage-validateur-server

ADD vendor /usr/local/go/src
ADD cmd /go/src/${PROJECT_DIR}/cmd
ADD models /go/src/${PROJECT_DIR}/models
ADD restapi /go/src/${PROJECT_DIR}/restapi
ADD merkle /go/src/${PROJECT_DIR}/merkle
ADD internal /go/src/${PROJECT_DIR}/internal

RUN apk add build-base

WORKDIR /go/src/${PROJECT_DIR}/cmd/r-c-g-horodatage-validateur-server/
RUN go build -o api


FROM alpine
# Create a group and user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

RUN apk update \
        && apk upgrade \
        && apk add --no-cache \
        ca-certificates \
        && update-ca-certificates 2>/dev/null || true

WORKDIR /app
COPY --from=build-env /go/src/github.com/geneva_validateur/cmd/r-c-g-horodatage-validateur-server/api /app/

RUN chown appuser:appgroup -R .
# Tell docker that all future commands should run as the appuser user
USER appuser

EXPOSE 8090