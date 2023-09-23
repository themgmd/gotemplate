FROM golang:1.21.1-alpine3.18 AS builder

RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates

WORKDIR /src

COPY /go.mod /go.sum ./

RUN go mod graph | awk '{if ($1 !~ "@") print $2}' | xargs go get

COPY . .

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -a -o ./go/bin/main ./cmd/main

FROM alpine:3.14

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /app

COPY --from=builder /src/go/bin/main .
COPY --from=builder /src/api/http/schema/ api/http/schema/
COPY --from=builder /src/logs/ logs/

CMD ["./main"]
