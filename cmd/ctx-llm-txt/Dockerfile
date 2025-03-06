FROM golang:1.23-alpine AS builder
WORKDIR /src
COPY . .
RUN go build -o /app/ctx-llm-txt

FROM alpine:latest
COPY --from=builder /app/ctx-llm-txt /app/ctx-llm-txt
EXPOSE 8080
CMD ["/app/ctx-llm-txt"]

