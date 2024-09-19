# Team 00 - Go Boot camp

#### Project layout

```txt
.
├── LICENSE
├── README.md
├── bin
├── cmd
│   ├── client
│   │   └── main.go
│   └── server
│       └── main.go
├── docs
│   └── task
│       ├── README_ENG.md
│       └── README_RUS.md
├── go.mod
├── go.sum
├── pkg
│   ├── api
│   │   ├── data_stream.proto
│   │   └── pb
│   │       ├── data_stream.pb.go
│   │       └── data_stream_grpc.pb.go
│   ├── config
│   │   └── config.go
│   ├── domain
│   │   ├── model
│   │   │   └── data.go
│   │   └── repository
│   │       └── data_repository.go
│   ├── infrastructure
│   │   ├── grpc
│   │   │   └── server.go
│   │   └── logger
│   │       └── logger.go
│   ├── usecase
│   │   └── data_service.go
│   └── utils
│       └── utils.go
├── scripts
│   └── setup.sh
└── tests
    ├── integration
    │   └── data_stream_integration_test.go
    └── unit
        └── data_stream_unit_test.go
```

Шаги:

```sh
# Установить protoc-gen-go
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# Установить protoc-gen-go-grpc
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

накидал в data_stream.proto того что требуют по заданию

запустил

```sh
protoc --go_out=. --go-grpc_out=. pkg/api/data_stream.proto
```

собрались клиент-сервер на го в пакете pb

затем по clean architecture раскидал исходники и смог запустить и сервер и клиент 

