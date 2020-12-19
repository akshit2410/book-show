FROM golang:latest

WORKDIR /app

COPY server1.go .

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN go build server1.go

EXPOSE $PORT

CMD ["./server1"]
