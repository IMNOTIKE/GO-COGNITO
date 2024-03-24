package config

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type ServerConfig struct {
	Port     int    `mapstructure:"port"`
	HostName string `mapstructure:"hostname"`
}

type DbConfig struct {
	Url      string `mapstructure:"url"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Pwd      string `mapstructure:"pwd"`
	Schema   string `mapstructure:"schema"`
	DbName   string `mapstructure:"db_name"`
	PoolSize int    `mapstructure:"pool_size"`
}

type AppConfig struct {
	Db     DbConfig     `mapstructure:"db"`
	Server ServerConfig `mapstructure:"server"`
}

func LoadConfig() (AppConfig, error) {

	viper.SetConfigName("config")
	viper.AddConfigPath("./env")
	if err := viper.ReadInConfig(); err != nil {
		log.Err(err)
		fmt.Print(err)
		return AppConfig{}, err
	}

	var out AppConfig
	if err := viper.Unmarshal(&out); err != nil {
		return AppConfig{}, err
	}

	return out, nil
}
