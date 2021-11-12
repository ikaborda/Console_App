package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Config struct {
	Countries string `yaml:"countries"`
	APIurl string `yaml:"url"`
}

func New(nameFile string) *Config{
	c:= Config{}
	return c.GetConfig(nameFile)
}

func (c *Config) GetConfig(nameFile string) *Config{
	yamlFile, err := ioutil.ReadFile(nameFile)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal:%v ", err)
	}

	return c
}
