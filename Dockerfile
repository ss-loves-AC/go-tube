FROM golang:1.23-alpine

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o go-tube main.go

CMD ["./go-tube"]