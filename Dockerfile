# Install go-lang for Linux
# A workspace (GOPATH) configured at /go.
FROM golang:1.14-alpine

RUN apk update

COPY ./myweb.go /

WORKDIR /

RUN go build myweb.go

ENTRYPOINT [ "./myweb" ]

# Document that the service listens on port 8080.
EXPOSE 8080
