FROM golang:1.22-alpine as builder
WORKDIR /app
COPY ./ ./
RUN go mod tidy && CGO_ENABLED=0 GOOS=linux go build -o webapp

FROM alpine:3.18 as runner
ENV PORT=3000 \
    LOG_PATH=/app/log/app.log
WORKDIR /app
COPY --from=builder /app/webapp /app
EXPOSE ${PORT}
CMD ["/app/webapp"]
