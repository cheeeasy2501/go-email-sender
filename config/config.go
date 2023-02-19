package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	grpc *GRPC
}

func (c *Config) GRPC() *GRPC {
	return c.grpc
}

type GRPC struct {
	addr string `mapstructure:"GRPC_HOST"`
	port string `mapstructure:"GRPC_PORT"`
}

func (c *GRPC) GetListenerAddr() string {
	return c.addr + ":" + c.port
}

func NewConfig(path string, t string) (*Config, error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType(t)
	
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	return &Config{
		grpc: &GRPC{
			addr: viper.GetString("GRPC_HOST"),
			port: viper.GetString("GRPC_PORT"),
		},
	}, nil
}
