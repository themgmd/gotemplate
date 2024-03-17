package config

import (
	"fmt"
	"github.com/caarlos0/env/v9"
	"log"
	"sync"
	"time"
)

var (
	config = &Config{}
	once   sync.Once
)

type Config struct {
	App     AppConfig
	HTTP    HTTPConfig
	Redis   RedisConfig
	Postgre PostgreConfig
}

type HTTPConfig struct {
	Host              string        `env:"HTTP_HOST"`
	ReadHeaderTimeout time.Duration `env:"READ_HEADER_TIMEOUT"`
}

type PostgreConfig struct {
	Host        string `env:"POSTGRES_HOST"`
	Port        int    `env:"POSTGRES_PORT"`
	Name        string `env:"POSTGRES_NAME"`
	User        string `env:"POSTGRES_USER"`
	Password    string `env:"POSTGRES_PASSWORD"`
	MaxIdleConn int    `env:"POSTGRES_MAX_IDLE_CONN"`
	MaxOpenConn int    `env:"POSTGRES_MAX_OPEN_CONN"`
}

func (pc PostgreConfig) DSN() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		pc.User, pc.Password, pc.Host, pc.Port, pc.Name)
}

type AppConfig struct {
	JwtSecret     string `env:"JWT_SECRET"`
	EncryptionKey string `env:"ENCRYPTION_KEY"`
}

type RedisConfig struct {
	Host     string `env:"REDIS_HOST"`
	Password string `env:"REDIS_PASSWORD"`
}

func Get() *Config {
	once.Do(func() {
		err := env.Parse(config)
		if err != nil {
			log.Fatal(err)
		}
	})

	return config
}
