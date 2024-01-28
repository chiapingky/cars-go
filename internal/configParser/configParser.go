package configparser

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
	Database struct {
		Url      string `yaml:"url"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		SslMode  string `yaml:"ssl-mode"`
	} `yaml:"database"`
}

func ParseConfig(location string) Config {
	file, err := os.ReadFile(location)
	if err != nil {
		log.Fatal(err)
	}
	var marshal Config
	err = yaml.Unmarshal(file, &marshal)
	if err != nil {
		log.Fatal(err)
	}
	return marshal
}
