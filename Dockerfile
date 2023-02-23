FROM golang:alpine3.17
ENV APP_HOME /go/src/plate-stack
RUN mkdir -p "$APP_HOME"
WORKDIR "$APP_HOME"
COPY plate-stack "$APP_HOME"/plate-stack
COPY ./views "$APP_HOME"/views
ENTRYPOINT ["/bin/plate-stack"]