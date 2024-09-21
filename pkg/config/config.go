package config

import "os"

type Config struct {
	GRPCPort   string
	PGConn     string
	GRPCAddr   string
	GRPCProt   string
	RetryCount int
}

func NewConfig() Config {
	return Config{
		GRPCPort:   ":50051",
		GRPCAddr:   "data-stream-server",
		PGConn:     os.Getenv("POSTGRES_CONN"),
		GRPCProt:   "tcp",
		RetryCount: 5,
	}
}
