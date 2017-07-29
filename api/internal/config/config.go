package config

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
)

type Config struct {
	SendGrid SendGrid `toml:"sendgrid"`
}

type SendGrid struct {
	Key         string `toml:"apiKey"`
	FromAddress string `toml:"fromAddress"`
}

func ReadConfig(filePath string) (*Config, error) {

	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	Config := &Config{}
	_, err = toml.Decode(os.ExpandEnv(string(bytes)), Config)
	if err != nil {
		return nil, err
	}

	return Config, nil
}
