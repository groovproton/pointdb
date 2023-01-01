FROM golang:1.19-alpine3.16 as init
RUN apk update && apk upgrade && \
    apk add --no-cache  git

FROM init as reload
WORKDIR /app
COPY ./app/go.mod  /app
RUN go mod download
RUN go install github.com/githubnemo/CompileDaemon@master

FROM reload as app
COPY ./app .

ENTRYPOINT CompileDaemon --build="go build -o /bin/main /app/main.go" --command=/bin/main