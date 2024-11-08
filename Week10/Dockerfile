#single stage dockerfile
FROM golang:1.23

WORKDIR /app

COPY . .

RUN go build -o myApp .

EXPOSE 80

CMD ["./myApp"]
