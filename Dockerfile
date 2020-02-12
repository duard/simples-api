FROM golang:1.13.4

WORKDIR /go-lang/src/github.com/duard/fiber-1

COPY . .

RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]

ENTRYPOINT CompileDaemon -log-prefix=false -build="go build ./cmd/api/" -command="./api"