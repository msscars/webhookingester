package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type MatcherConfig struct {
	Type  string `yaml:"type"`
	Token string `yaml:"token"`
	Key   string `yaml:"key"`
}

type PublisherConfig struct {
	Type         string `yaml:"type"`
	Uri          string `yaml:"uri"`
	Subject      string `yaml:"subject"`
	BodySelector string `yaml:"bodySelector"`
}

type RoutingConfig struct {
	Matcher   MatcherConfig   `yaml:"matcher"`
	Publisher PublisherConfig `yaml:"publisher"`
}

type Config struct {
	Routings []RoutingConfig `yaml:"routings"`
}

func (c *Config) GetConf() *Config {

	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}
