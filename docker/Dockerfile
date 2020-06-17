FROM golang:1.14.4-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod download
RUN go mod verify

RUN GOOS=linux GOARCH=amd64 go build -o metrics-aggregator ./cmd/app/main.go


FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

EXPOSE 4000

COPY --from=builder /app/metrics-aggregator .

ENTRYPOINT ["./metrics-aggregator"]