package config

import (
	"log"
	"os"
	"path/filepath"
	"sync"

	"github.com/spf13/viper"
)

var once sync.Once

func init() {
	once.Do(func() {
		LoadConfig()
	})
}

type WebConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}

type JWTSecret string

type RedisConfig struct {
	Addr     string `mapstructure:"addr"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type Config struct {
	Web       WebConfig      `mapstructure:"web"`
	Database  DatabaseConfig `mapstructure:"database"`
	JWTSecret JWTSecret      `mapstructure:"jwt_secret"`
	Redis     RedisConfig    `mapstructure:"redis"`
}

var AppConfig *Config

func LoadConfig() {
	ex, err := os.Executable()
	if err != nil {
		log.Fatalf("get exe dir failed: %s", err.Error())
	}
	exeDir := filepath.Dir(ex)

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(exeDir + "/config")
	viper.AddConfigPath("/home/jeanfo/workspace/releases/jeanfo_mix_api/config")
	// viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("Effor unmarshaling config: %v", err)
	}
}
