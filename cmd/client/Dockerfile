
FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY ./cmd/client /app/cmd/client
COPY ./pkg /app/pkg

RUN go build -o client ./cmd/client/main.go

EXPOSE 8080

# Run the client
CMD ["./client"]
