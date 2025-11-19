package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HttpServer struct {
	Addr string `yaml:"address" env-required:"true"`
}

type Config struct {
	Env        string     `yaml:"env" env:"ENV" env-required:"true" env-default:"production"`
	Storage    string     `yaml:"storage-path" env-required:"true"`
	HttpServer HttpServer `yaml:"http_server"`
}

func MustLoad() *Config {
	var ConfigPath string

	ConfigPath = os.Getenv("CONFIG_PATH")

	if ConfigPath == "" {
		flags := flag.String("config", "", "path to config")
		flag.Parse()

		ConfigPath = *flags

		if ConfigPath == "" {
			log.Fatal("config path is not set")
		}
	}

	if _, err := os.Stat(ConfigPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist at path: %s", ConfigPath)
	}

	var cfg Config

	err := cleanenv.ReadConfig(ConfigPath, &cfg)

	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	return &cfg
}
