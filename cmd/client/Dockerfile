# Build stage
FROM golang:1.22 AS builder

WORKDIR /app

# Скопируем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Скопируем исходники
COPY ./cmd/client /app/cmd/client
COPY ./pkg /app/pkg

# Собираем приложение

ENV CGO_ENABLED=0

RUN go build -o /app/client ./cmd/client/main.go

# Final stage с использованием scratch
FROM scratch AS runtime

WORKDIR /app

# Копируем скомпилированное приложение из стадии сборки
COPY --from=builder /app/client /app/client

# Открываем порт
EXPOSE 8080

# Запускаем приложение с аргументом
CMD ["/app/client"]
