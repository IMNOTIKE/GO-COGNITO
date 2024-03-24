package config

import (
	"log"

	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

var k = koanf.New(".")

type ServerConfig struct {
	Port     int    `koanf:"server.port"`
	HostName string `koanf:"server.hostname"`
}

type DbConfig struct {
	Url      string `koanf:"db.url"`
	Port     int    `koanf:"db.port"`
	User     string `koanf:"db.user"`
	Pwd      string `koanf:"db.pwd"`
	Schema   string `koanf:"db.schema"`
	DbName   string `koanf:"db.db_name"`
	PoolSize int    `koanf:"db.pool_size"`
}

type AppConfig struct {
	Db     DbConfig     `koanf:"db"`
	Server ServerConfig `koanf:"server"`
}

func LoadConfig() AppConfig {
	// Load JSON config.
	f := file.Provider("env/config.yaml")
	if err := k.Load(f, yaml.Parser()); err != nil {
		log.Fatalf("error loading config: %v", err)
	}

	var out AppConfig

	k.UnmarshalWithConf("", &out, koanf.UnmarshalConf{Tag: "koanf"})

	return out
}
