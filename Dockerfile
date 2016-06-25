FROM golang:1.6

RUN go get github.com/nlopes/slack

COPY . /go/src/app
WORKDIR /go/src/app

RUN go install
CMD ["app"]
