FROM golang:1.14-alpine as base
WORKDIR /app

COPY . .

RUN go mod tidy
RUN go get github.com/githubnemo/CompileDaemon
ENTRYPOINT CompileDaemon -exclude-dir=.git -command="./myomer"
