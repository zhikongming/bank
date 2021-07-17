package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ghodss/yaml"
)

type DBConfig struct {
	WriteDSN string `yaml:"WriteDSN"`
	ReadDSN  string `yaml:"ReadDSN"`
}

type Config struct {
	DB DBConfig `yaml:"DB"`
}

var c *Config

func GetConfig() *Config {
	return c
}

func InitConfig() {
	configFilePath := "conf/config.yaml"
	if _, err := os.Stat(configFilePath); err != nil {
		panic(fmt.Sprintf("file %s not exist", configFilePath))
	}
	fileContent, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		panic(err)
	}
	if err := yaml.Unmarshal(fileContent, &c); err != nil {
		panic(err)
	}
}
