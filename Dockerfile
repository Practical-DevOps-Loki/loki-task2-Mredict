FROM golang:1.22-alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o webapp .

FROM alpine:3.18 as runner
ENV PORT=3000 \
    LOG_PATH=/app/log/app.log
WORKDIR /app
COPY --from=builder /app/webapp .
COPY --from=builder /app/public ./public
EXPOSE ${PORT}
CMD ["./webapp"]
