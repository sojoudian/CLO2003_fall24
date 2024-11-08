# Multi-stage dockerfile
FROM golang:1.23 AS builder
WORKDIR /app
COPY . .
RUN go build -o myApp .


FROM alpine:3.20
WORKDIR /server
COPY --from=builder /app/myApp .
EXPOSE 80
CMD ["./myApp"]
