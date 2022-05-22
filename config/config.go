package config

import (
	"errors"
	"fmt"
	"github.com/ghodss/yaml"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
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

var DB *gorm.DB
var RDB *redis.Client

func (c *Config) GetConfig() error {
	if _, err := os.Stat("./config.yml"); errors.Is(err, os.ErrNotExist) {
		if _, err := os.Stat("./config.development.yml"); errors.Is(err, os.ErrNotExist) {
			return err
		}
		fmt.Println("INFO: find config.development.yml")
		return c.ReadConfig("./config.development.yml")
	} else {
		fmt.Println("INFO: find config.yml")
		return c.ReadConfig("./config.yml")
	}
}

func (c *Config) ReadConfig(filename string) error {
	ymlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(ymlFile, c)
	if err != nil {
		return err
	}
	return nil
}
