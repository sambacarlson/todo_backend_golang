# FROM alpine:3.18
FROM golang:latest
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .
expose 4356
CMD ["./main"]