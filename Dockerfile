FROM golang:1.22-alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o webapp .

FROM alpine:3.18 as runner
WORKDIR /app
COPY --from=builder /app/webapp .
EXPOSE 8080
CMD ["./webapp"]
