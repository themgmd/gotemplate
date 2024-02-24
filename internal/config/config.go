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

func Get() *Config {
	return config
}

type Config struct {
	App     AppConfig
	HTTP    HTTPConfig
	Redis   RedisConfig
	Postgre PostgreConfig
}

type HTTPConfig struct {
	Host         string        `env:"HTTP_HOST"`
	Port         string        `env:"HTTP_PORT"`
	WriteTimeout time.Duration `env:"HTTP_WRITE_TIMEOUT"`
	ReadTimeout  time.Duration `env:"HTTP_READ_TIMEOUT"`
	IdleTimeout  time.Duration `env:"HTTP_IDLE_TIMEOUT"`
}

func (h HTTPConfig) GetAddress() string {
	return fmt.Sprintf("%s:%s", h.Host, h.Port)
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

func (pc PostgreConfig) GetDSN() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		pc.User, pc.Password, pc.Host, pc.Port, pc.Name)
}

func (pc PostgreConfig) GetMaxIdleConn() int {
	return pc.MaxIdleConn
}

func (pc PostgreConfig) GetMaxOpenConn() int {
	return pc.MaxOpenConn
}

type AppConfig struct {
	JwtSecret     string `env:"JWT_SECRET"`
	EncryptionKey string `env:"ENCRYPTION_KEY"`
}

type RedisConfig struct {
	Host     string `env:"REDIS_HOST"`
	Port     string `env:"REDIS_PORT"`
	Password string `env:"REDIS_PASSWORD"`
}

func (rc RedisConfig) Addr() string {
	return fmt.Sprintf("%s:%s", rc.Host, rc.Port)
}
func (rc RedisConfig) Pass() string {
	return rc.Password
}

func Init() {
	once.Do(func() {
		err := env.Parse(config)
		if err != nil {
			log.Fatal(err)
		}
	})
}
