# Team 00 - Go Boot camp

## **Обзор**

Этот проект создан для симуляции связи с инопланетным кораблем через анализ передаваемых частот. Он включает в себя несколько ключевых компонентов: gRPC-сервер, который передает поток данных о частотах, клиент для обнаружения аномалий в этом потоке, а также сохранение обнаруженных аномалий в базе данных PostgreSQL с использованием ORM.

Проект разделен на четыре основные задачи:
1. **Передатчик** - gRPC-сервер, генерирующий поток частот, распределенных по нормальному закону.
2. **Обнаружение аномалий** - клиент, обнаруживающий статистические аномалии в потоке частот.
3. **Отчет** - сохранение данных об аномалиях в базе данных PostgreSQL с использованием ORM.
4. **Интеграция** - объединение передатчика, клиента и базы данных для работы в единой системе.

## **Структура проекта**

```txt
.
├── LICENSE
├── Makefile
├── README.md
├── bin
├── cmd
│   ├── client
│   │   ├── Dockerfile
│   │   └── main.go
│   └── server
│       ├── Dockerfile
│       └── main.go
├── data.txt
├── docker-compose.yml
├── docs
│   └── task
│       ├── README_ENG.md
│       └── README_RUS.md
├── go.mod
├── go.sum
└── pkg
    ├── api
    │   ├── data_stream.proto
    │   └── pb
    │       ├── data_stream.pb.go
    │       └── data_stream_grpc.pb.go
    ├── config
    │   └── config.go
    ├── database
    │   ├── anomaly.go
    │   └── database.go
    ├── domain
    │   ├── model
    │   │   └── data.go
    │   └── repository
    │       └── data_repository.go
    ├── infrastructure
    │   ├── grpc
    │   │   └── server.go
    │   └── logger
    │       └── logger.go
    ├── usecase
    │   ├── anomaly.go
    │   ├── data_service.go
    │   └── detector.go
    └── utils
        └── utils.go
```

1. **Задача 00: Передатчик**
    - Реализован gRPC-сервер, который симулирует работу военного устройства, передающего данные о частоте.
    - При каждом новом подключении клиента сервер генерирует уникальный идентификатор сессии (UUID), случайное среднее значение и стандартное отклонение для нормального распределения.
    - Сервер логирует все сгенерированные данные (session_id, среднее значение, стандартное отклонение) и передает клиенту поток сообщений с полями:
        - `session_id` (строка)
        - `frequency` (случайная величина, распределенная по нормальному закону)
        - `timestamp` (время в формате UTC).

2. **Задача 01: Обнаружение аномалий**
    - Клиент получает поток частот от gRPC-сервера и динамически вычисляет параметры нормального распределения (среднее значение и стандартное отклонение).
    - После накопления достаточного количества данных клиент переключается в режим обнаружения аномалий.
    - Аномалия определяется как частота, отклоняющаяся от ожидаемого значения более чем на `k * STD`, где `k` — это коэффициент, переданный через параметр командной строки.
    - Клиент логирует количество обработанных данных и найденные аномалии.

3. **Задача 02: Отчет**
    - Реализовано сохранение данных об аномалиях в базе данных PostgreSQL.
    - Для взаимодействия с базой данных используется ORM (можно выбрать `go-pg` или `GORM`).
    - Структура записи в базе данных:
        - `session_id` (строка)
        - `frequency` (дробное число)
        - `timestamp` (время UTC).

4. **Задача 03: Интеграция**
    - Все компоненты объединены в одну систему:
        - Сервер передает поток частот.
        - Клиент получает поток, вычисляет параметры распределения, обнаруживает аномалии.
        - Обнаруженные аномалии записываются в базу данных PostgreSQL.

## **Как запустить проект**

### Установка зависимостей

Проект использует **Docker Compose**. Для установки зависимостей выполните:

```bash
docker pull golang:1.22
docker pull postgres:latest
```

### Сборка и запуск проекта

Для сборки и запуска всего проекта используйте:

```bash
docker compose up --build
```
Запустится и добавятся значения из `.env` файла:
* База данных PostgreSQL
* Сервер
* Клиент который будет принимать поток от сервера и сразу же анализировать данные на предмет аномалий

## **Используемые технологии**

* **gRPC** - для реализации сервер-клиентной архитектуры.
* **PostgreSQL** - база данных для хранения аномалий.
* **ORM GORM** - для безопасной работы с базой данных без прямого написания SQL-запросов
* **Go** - основной язык программирования для всех компонентов.
* **Docker compose** - для упаковывания проекта в независимые среды


