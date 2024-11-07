package config

import (
	"fmt"
	"sync/atomic"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `yaml:"server" mapstructure:"server"`
	Database DatabaseConfig `yaml:"database" mapstructure:"database"`
	Storage  StorageConfig  `yaml:"storage" mapstructure:"storage"`
	Cors     CorsConfig     `yaml:"cors" mapstructure:"cors"`
}

type ServerConfig struct {
	HTTPPort int32  `yaml:"http_port" mapstructure:"http_port"`
	GRPCPort int32  `yaml:"grpc_port" mapstructure:"grpc_port"`
	Host     string `yaml:"host" mapstructure:"host"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host" mapstructure:"host"`
	Port     int    `yaml:"port" mapstructure:"port"`
	User     string `yaml:"user" mapstructure:"user"`
	Password string `yaml:"password" mapstructure:"password"`
	DBName   string `yaml:"dbname" mapstructure:"dbname"`
}

type StorageConfig struct {
	Path    string `yaml:"path" mapstructure:"path"`
	MaxSize int64  `yaml:"max_size" mapstructure:"max_size"`
}

type CorsConfig struct {
	AllowedOrigins []string `yaml:"allowed_origins" mapstructure:"allowed_origins"`
	AllowedMethods []string `yaml:"allowed_methods" mapstructure:"allowed_methods"`
}

var cfg atomic.Value

// LoadConfig 载配置文件
func LoadConfig(path string) (staticCfg *Config, err error) {
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	if err = viper.ReadInConfig(); err != nil {
		return
	}
	staticCfg = &Config{}
	if err = viper.Unmarshal(&staticCfg); err != nil {
		return
	}
	cfg.Store(staticCfg)
	return
}

func GetConfig() *Config {
	return cfg.Load().(*Config)
}

// GetDSN 获取数据库连接字符串
func GetDSN() string {
	c := GetConfig()
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		c.Database.User, c.Database.Password, c.Database.Host, c.Database.Port, c.Database.DBName)
}
