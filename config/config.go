package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var err error

type Config struct {
	Host      string `yaml:"service_host"`
	Port      int    `yaml:"service_port"`
	DataFile  string `yaml:"db_file"`
	Reset     int    `yaml:"main_window"`
	CountUrl  string `yaml:"count_URL"`
	LimitSize int    `yaml:"limit_size"`
	LimitRate int    `yaml:"limit_window"`
}

type cfg struct {
	file *os.File
}

type Configuration interface {
	Get() (*Config, error)
}

func New(file *os.File) Configuration {
	return &cfg{file}
}

func (c *cfg) Get() (*Config, error) {

	f, err := ioutil.ReadAll(c.file)
	if err != nil {
		return nil, err
	}
	config := &Config{}
	err = yaml.Unmarshal(f, config)
	if err != nil {
		return config, err
	}
	return config, nil
}
