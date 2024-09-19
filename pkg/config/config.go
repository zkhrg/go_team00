package config

type Config struct {
	GRPCPort string
}

func LoadConfig() *Config {
	return &Config{
		GRPCPort: ":50051",
	}
}
