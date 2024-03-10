package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Port         string `yaml:"port"`
		Host         string `yaml:"host"`
		JwtSecretKey string `yaml:"jwtSecretKey"`
		CookieName   string `yaml:"cookieName"`
	} `yaml:"server"`

	Cookie struct {
		Name     string `yaml:"name"`
		MaxAge   int    `yaml:"maxAge"`
		Secure   bool   `yaml:"secure"`
		HTTPOnly bool   `yaml:"httpOnly"`
	} `yaml:"cookie"`
	
	JsonDB struct {
		Filename string `yaml:"filename"`
	} `yaml:"jsonDB"`
}

func Decode(cfgFile string) (*Config, error) {
	f, err := os.Open(cfgFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
