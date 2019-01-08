FROM golang:1.10

WORKDIR /go/src/github.com/gpng/go-docker-api-boilerplate

COPY . .

RUN ["go", "get", "github.com/githubnemo/CompileDaemon"]

ENTRYPOINT CompileDaemon -log-prefix=false -build="go build -o bin/application cmd/api/main.go" -command="./bin/application"