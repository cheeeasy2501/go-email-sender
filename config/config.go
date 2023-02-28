package config

import (
	"github.com/cheeeasy2501/go-email-sender/pkg/logger"
	"github.com/spf13/viper"
)

/* All configurations */
type Config struct {
	app  *App
	grpc *GRPC
	mail IMail
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
		app:  newApp(),
		grpc: NewGRPC(),
		mail: NewMailConfig(),
	}, nil
}

func (c *Config) GetAppEnv() string {
	return c.app.env
}

func (c *Config) GetAppLoggerType() logger.LoggerType {
	return c.app.logger
}

func (c *Config) App() *App {
	return c.app
}

func (c *Config) GRPC() *GRPC {
	return c.grpc
}

func (c *Config) Mail() IMail {
	return c.mail
}

type App struct {
	env    string
	logger logger.LoggerType
}

func newApp() *App {
	return &App{
		env:    viper.GetString("APP_ENV"),
		logger: logger.LoggerType(viper.GetString("APP_LOGGER")),
	}
}

/* GRPC Configuration */
type GRPC struct {
	enable bool   `mapstructure:"GRPC_ENABLE"`
	addr   string `mapstructure:"GRPC_HOST"`
	port   string `mapstructure:"GRPC_PORT"`
}

func NewGRPC() *GRPC {
	return &GRPC{
		enable: viper.GetBool("GRPC_ENABLE"),
		addr:   viper.GetString("GRPC_HOST"),
		port:   viper.GetString("GRPC_PORT"),
	}
}

func (c *GRPC) IsGRPCEnable() bool {
	return c.enable
}

func (c *GRPC) GetListenerAddr() string {
	return c.addr + ":" + c.port
}

/** Mail configuration */
type IMail interface {
	GetHost() string
	GetPort() string
	GetAddress() string
	GetUsername() string
	GetPassword() string
	GetEncryption() string
	GetAddressFrom() string
}

type Mail struct {
	host        string
	port        string
	username    string
	password    string
	encryption  string
	addressFrom string
}

func NewMailConfig() *Mail {
	return &Mail{
		host:        viper.GetString("MAIL_HOST"),
		port:        viper.GetString("MAIL_PORT"),
		username:    viper.GetString("MAIL_USERNAME"),
		password:    viper.GetString("MAIL_PASSWORD"),
		encryption:  viper.GetString("MAIL_ENCRYPTION"),
		addressFrom: viper.GetString("MAIL_FROM"),
	}
}

func (c *Mail) GetHost() string {
	return c.host
}

func (c *Mail) GetPort() string {
	return c.port
}

func (c *Mail) GetAddress() string {
	return c.host + ":" + c.port
}

func (c *Mail) GetUsername() string {
	return c.username
}

func (c *Mail) GetPassword() string {
	return c.password
}

func (c *Mail) GetEncryption() string {
	return c.encryption
}

func (c *Mail) GetAddressFrom() string {
	return c.addressFrom
}
