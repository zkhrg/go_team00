# Team 00 - Go Boot camp

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

