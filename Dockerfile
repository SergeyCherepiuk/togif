FROM golang:1.21-alpine
RUN apk update && \
    apk upgrade && \
    apk add ffmpeg

WORKDIR /togif
COPY . .

ENTRYPOINT [ "go", "run", "cmd/cli/main.go" ]