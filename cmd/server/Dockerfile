# Первая стадия: сборка приложения
FROM golang:1.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./cmd/server /app/cmd/server
COPY ./pkg /app/pkg

ENV CGO_ENABLED=0

RUN go build -o server ./cmd/server/main.go

# Вторая стадия: минимальный образ
FROM scratch AS runtime

COPY --from=builder /app/server /app/server

EXPOSE 50051

CMD ["/app/server"]
