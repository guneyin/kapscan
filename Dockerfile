FROM golang:1.23-bookworm as builder

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . ./

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g cmd/app/main.go
RUN go build -v -o server cmd/app/main.go

FROM debian:bookworm-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /app/server /app/server
COPY --from=builder /app/docs/ /app/docs/

WORKDIR /app

CMD ["/app/server"]
