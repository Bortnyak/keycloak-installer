package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	BaseURL       string `yaml:"baseURL"`
	ApiType       string `yaml:"apiType"`
	Realm         string `yaml:"realm"`
	Client        string `yaml:"clinet"`
	AdminLogin    string `yaml:"adminLogin"`
	AdminPassword string `yaml:"adminPassword"`
}

func GetConf() *Config {
	c := &Config{}
	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		fmt.Printf("yamlFile.Get err   #%v ", err)
	}

	fmt.Println(string(yamlFile))

	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		fmt.Printf("Unmarshal: %v", err)
	}

	return c
}

func (c *Config) GetBaseURL() string {
	return c.BaseURL
}
