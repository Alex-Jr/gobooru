FROM golang:1.21-alpine3.18 AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /app/cmd/main /app/cmd

EXPOSE 8080

FROM alpine:3.18

RUN apk update && \
  apk add --no-cache ffmpeg

COPY --from=builder ["/app/cmd/main", "/app/"]

COPY --from=builder ["/app/migrations", "/app/migrations"]

CMD [ "/app/main" ]