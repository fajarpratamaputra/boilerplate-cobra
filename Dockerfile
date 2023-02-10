FROM golang:1.15-alpine

WORKDIR /app
COPY . .

RUN go build -o main .

CMD ["./main"]
