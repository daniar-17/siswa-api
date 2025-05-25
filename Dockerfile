FROM golang:1.24

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

RUN ls -la /app   # cek apakah file main ada

CMD ["./main"]
