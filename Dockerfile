# Build app:
FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod .
RUN go mod download

COPY main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -o bye .

# Create minimal image with only app binary:
FROM alpine:edge

WORKDIR /app

COPY --from=builder /app/bye .

ENTRYPOINT ["/app/bye"]
