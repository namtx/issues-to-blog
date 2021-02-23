FROM golang:alpine AS build

WORKDIR $GOPATH/src/github.com/namtx/issues-to-blog

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

CMD ["issues-to-blog"]
