FROM golang:1.6

ENV APP_PATH /go/src/scrumkin
ENV GOBIN $GOPATH/bin

COPY . $APP_PATH
WORKDIR $APP_PATH

CMD ["scrumkin"]
