FROM golang:alpine AS build-env

ENV GOBIN $GOPATH/bin
ENV PROJECT_DIR /go/src/github.com/geneva_validateur

ADD main.go ${PROJECT_DIR}/
ADD mockup ${PROJECT_DIR}/mockup

WORKDIR ${PROJECT_DIR}

RUN go build -o webapp


FROM alpine
# Create a group and user
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Run the app
WORKDIR /app
COPY --from=build-env /go/src/github.com/geneva_validateur/webapp /app/
COPY --from=build-env /go/src/github.com/geneva_validateur/mockup /app/mockup

RUN chown appuser:appgroup -R .
# Tell docker that all future commands should run as the appuser user
USER appuser

EXPOSE 8080
