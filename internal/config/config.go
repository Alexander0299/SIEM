package config

type Config struct {
	Host string
	Port int
}

func NewConfig() *Config {
	return &Config{
		Host: "0.0.0.0",
		Port: 8585,
	}

}
