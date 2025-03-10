package config

import (
	"fmt"
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
	Host                string `mapstructure:"host"`
	Port                int    `mapstructure:"port"`
	ProjRoot            string `mapstructure:"proj_root"`
	UploadDir           string `mapstructure:"upload_dir"`
	UploadDirStaticPath string `mapstructure:"upload_dir_static_path"`
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
	Web       WebConfig      `yaml:"web"`
	Database  DatabaseConfig `yaml:"database"`
	JWTSecret JWTSecret      `mapstructure:"jwt_secret"`
	Redis     RedisConfig    `yaml:"redis"`
	Log       LogConfig      `yaml:"log"`
}

type LogConfig struct {
	Dir     string        `yaml:"dir"`
	Console bool          `yaml:"console"`
	Level   string        `yaml:"level"`
	Normal  LogFileConfig `yaml:"normal"`
	Error   LogFileConfig `yaml:"error"`
}

type LogFileConfig struct {
	MaxSize    int `mapstructure:"max_size"`
	MaxBackups int `mapstructure:"max_backups"`
}

var AppConfig *Config

// GetConfig 获取配置实例
func GetConfig() *Config {
	if AppConfig == nil {
		LoadConfig()
	}
	return AppConfig
}

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	if configDir := os.Getenv("JMPConfigDir"); len(configDir) > 0 {
		fmt.Println("Specified ConfigDir: " + configDir)
		viper.AddConfigPath(configDir)
	} else {
		fmt.Println("No Specified ConfigDir - Now Search Some Default Dirs...")

		ex, _ := os.Executable()
		exeDir := filepath.Dir(ex)
		viper.AddConfigPath(exeDir + "/config")

		viper.AddConfigPath("/Users/jeanfo/workspace/jeanfo_mix_api/config")
		viper.AddConfigPath("/home/jeanfo/workspace/releases/jeanfo_mix_api/config")
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("Error unmarshaling config: %v", err)
	}
}
