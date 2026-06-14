FROM golang:1.25

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main ./cmd/server

EXPOSE 3000

CMD ["./main"]