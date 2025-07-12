package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string `yaml:"addr" env-required:"true"`
}
type Config struct {
	Env          string `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
	Storage_path string `yaml:"storage_path" env-required:"true"`
	HTTPServer   `yaml:"http_server"`
}

func MustLoad() *Config {
	var config_path string
	config_path = os.Getenv("CONFIG_PATH")
	if config_path == "" {
		flags := flag.String("config", "", "path to config file")
		flag.Parse()
		config_path = *flags
		if config_path == "" {
			log.Fatal("config path is not set")
		}
	}
	if _, err := os.Stat(config_path); os.IsNotExist(err) {
		log.Fatalf("config file %s does not exist", config_path)
	}
	var cfg Config
	err := cleanenv.ReadConfig(config_path, &cfg)
	if err != nil {
		log.Fatalf("can not read config file %s", err.Error())
	}
	return &cfg
}
