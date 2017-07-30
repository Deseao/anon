package config

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"os"
)

type Config struct {
	SendGrid SendGrid `toml:"sendgrid"`
	Twilio   Twilio   `toml:"twilio"`
}

type SendGrid struct {
	Key         string `toml:"apiKey"`
	FromAddress string `toml:"fromAddress"`
}

type Twilio struct {
	SID        string `toml:"sid"`
	Key        string `toml:"key"`
	AccountID  string `toml:"accountId"`
	FromNumber string `toml:"fromNumber"`
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
