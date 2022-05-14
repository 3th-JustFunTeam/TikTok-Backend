package config

import (
	"errors"
	"fmt"
	"github.com/ghodss/yaml"
	"io/ioutil"
	"os"
)

type Config struct {
	MySQL MySQL `yaml:"mysql"`
	Redis Redis `yaml:"redis"`
}

type MySQL struct {
	Host     string `yaml:"host"`
	Port     int32  `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type Redis struct {
	Host string `yaml:"host"`
	Port int32  `yaml:"port"`
}

func (c *Config) GetConfig() {
	if _, err := os.Stat("./config.yml"); errors.Is(err, os.ErrNotExist) {
		if _, err := os.Stat("./config.development.yml"); errors.Is(err, os.ErrNotExist) {
			panic("ERROR: cannot find config file")
		}
		fmt.Println("INFO: find config.development.yml")
		c.readConfig("./config.development.yml")
	} else {
		fmt.Println("INFO: find config.yml")
		c.readConfig("./config.yml")
	}
}

func (c *Config) readConfig(filename string) {
	ymlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err.Error())
	}
	err = yaml.Unmarshal(ymlFile, c)
	if err != nil {
		panic(err.Error())
	}
}
