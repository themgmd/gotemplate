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
	Port        string `env:"POSTGRES_PORT"`
	Name        string `env:"POSTGRES_NAME"`
	User        string `env:"POSTGRES_USER"`
	Password    string `env:"POSTGRES_PASSWORD"`
	MaxIdleConn int    `env:"MAX_IDLE_CONN" default:"120"`
	MaxOpenConn int    `env:"MAX_OPEN_CONN" default:"60"`
}

func (pc PostgreConfig) GetDSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		pc.Host, pc.Port, pc.User, pc.Name, pc.Password)
}

func (pc PostgreConfig) GetMaxIdleConn() int {
	return pc.MaxIdleConn
}

func (pc PostgreConfig) GetMaxOpenConn() int {
	return pc.MaxOpenConn
}

type AppConfig struct {
	EncryptionKey string `env:"ENCRYPTION_KEY,required"`
}

type RedisConfig struct {
	Host string `env:"REDIS_HOST"`
	Port string `env:"REDIS_PORT"`
}

func (rc RedisConfig) Addr() string {
	return fmt.Sprintf("%s:%s", rc.Host, rc.Port)
}

func Init() {
	once.Do(func() {
		err := env.Parse(config)
		if err != nil {
			log.Fatal(err)
		}
	})
}
