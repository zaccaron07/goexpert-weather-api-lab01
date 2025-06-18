FROM golang:latest AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /goexpert-weather-api ./cmd/lab01

FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

COPY --from=builder /goexpert-weather-api .

EXPOSE 8080

ENTRYPOINT ["./goexpert-weather-api"]
