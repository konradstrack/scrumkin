FROM golang:1.6

ENV APP_PATH /go/src/scrumway
ENV GOBIN $GOPATH/bin

COPY . $APP_PATH
WORKDIR $APP_PATH

CMD ["scrumway"]
