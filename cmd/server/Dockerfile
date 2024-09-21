
FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY ./cmd/server /app/cmd/server
COPY ./pkg /app/pkg

RUN go build -o server ./cmd/server/main.go

EXPOSE 50051

CMD ["./server"]
