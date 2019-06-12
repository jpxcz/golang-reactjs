# Documentation taken from https://blog.golang.org/docker
FROM golang:1.12.5

WORKDIR /$GOPATH/src/github.com/josepablo93/api
COPY . .
RUN go build -o /go/bin/api

CMD ["/go/bin/api"]