FROM golang:1.23.0 as builder
WORKDIR /app

COPY go.mod ./
RUN go mod download
COPY . ./
RUN go build -o main .

FROM debian:bookworm-slim
WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/certs ./certs
COPY --from=builder /app/templates ./templates

EXPOSE 443

CMD ["./main"]