FROM golang:1.24 AS builder

WORKDIR /app

COPY ../go.mod ../go.sum ./
RUN go mod download

COPY ../ ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/weather-cli ./cmd/main.go

FROM alpine

COPY --from=builder /app/weather-cli /usr/local/bin/weather-cli

RUN chmod +x /usr/local/bin/weather-cli

ENTRYPOINT ["weather-cli"]