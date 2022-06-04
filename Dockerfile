FROM golang:alpine

WORKDIR /app

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o cron-parser cmd/cli/main.go

FROM alpine:latest

COPY --from=0 /app ./

ENTRYPOINT ["./cron-parser"]
