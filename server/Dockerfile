FROM golang:1.21-alpine3.18

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /app/cmd/main /app/cmd

EXPOSE 8080

WORKDIR /app/cmd

CMD [ "./main" ]