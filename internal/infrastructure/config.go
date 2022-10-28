package infrastructure

import (
	"github.com/caarlos0/env/v6"
	"net"
	"time"
)

var (
	AppName      = "Дайте мне футболочку, плес"
	AppNamespace = "edu"
	AppTag       = "v0.0.0"
)

type Config struct {
	Database struct {
		Host     string `env:"DB_HOST,required"`
		Port     string `env:"DB_PORT,required"`
		User     string `env:"DB_USER,required"`
		Password string `env:"DB_PASSWORD,required"`
		Name     string `env:"DB_NAME,required"`
		MaxConn  int32  `env:"DB_POOL_MAX_CON" envDefault:"10"`
	}

	HTTPServer struct {
		Port string `env:"APP_HTTP_PORT,required"`
		Addr string
	}

	HTTPClient struct {
		Timeout           time.Duration
		RetryInterval     time.Duration
		HTTPTimeout       int `env:"HTTP_TIMEOUT" envDefault:"30"`
		RetryCount        int `env:"HTTP_RETRY_COUNT" envDefault:"5"`
		HTTPRetryInterval int `env:"HTTP_RETRY_INTERVAL" envDefault:"2"`
	}
}

func NewConfigFromEnv() (*Config, error) {
	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	cfg.HTTPServer.Addr = net.JoinHostPort("", cfg.HTTPServer.Port)
	cfg.HTTPClient.Timeout = time.Second * time.Duration(cfg.HTTPClient.HTTPTimeout)
	cfg.HTTPClient.RetryInterval = time.Second * time.Duration(cfg.HTTPClient.HTTPRetryInterval)

	return cfg, nil
}
