FROM golang:1.24 AS builder

WORKDIR /app

COPY ../go.mod ../go.sum ./
RUN go mod download

COPY ../ ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ./weather-cli ./cmd/main.go

FROM alpine

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /usr/local/bin

COPY --from=builder /app/weather-cli ./weather-cli
COPY --from=builder /app/config/prod/prod.yaml ./config/prod/prod.yaml

RUN chmod +x /usr/local/bin/weather-cli

ENTRYPOINT ["weather-cli"]

CMD [ "--help" ]