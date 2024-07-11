FROM golang:latest AS builder
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .
CMD ["./main"]

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .
CMD [ "./main" ]