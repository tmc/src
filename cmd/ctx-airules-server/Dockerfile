FROM golang:1.23-alpine AS builder
WORKDIR /src
COPY . .
RUN go build -o /app/ctx-airules-server

FROM alpine:latest
COPY --from=builder /app/ctx-airules-server /app/ctx-airules-server
RUN mkdir -p /app/outputs
EXPOSE 8081
CMD ["/app/ctx-airules-server"]

