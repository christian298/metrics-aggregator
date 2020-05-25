FROM golang:1.14.0-alpine AS builder

RUN apk update && apk add --no-cache git

WORKDIR $GOPATH/src/mypackage/myapp/
COPY .. .

RUN go mod download
RUN go mod verify

RUN GOOS=linux GOARCH=amd64 go build -o /go/bin/metrics-aggregator ./cmd/app/main.go

FROM scratch

COPY --from=builder /go/bin/metrics-aggregator /go/bin/metrics-aggregator

ENTRYPOINT ["/go/bin/metrics-aggregator"]