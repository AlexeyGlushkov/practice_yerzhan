FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o app ./cmd

EXPOSE 8080
EXPOSE 50051

CMD ["./app"]
