package lib

import (
	"encoding/json"
	"errors"
	"io"
	"os"
)

type Config struct {
	CACHE   string `json:"cache"`
	ROAMING string `json:"roaming"`
}

type ConfigLoader struct {
	Config Config
}

func NewConfigLoader() *ConfigLoader {
	return &ConfigLoader{}
}

func (c *ConfigLoader) validateConfig(config Config) error {
	if config.ROAMING == "" || config.CACHE == "" {
		return errors.New("invalid config")
	}
	return nil
}

func (c *ConfigLoader) InitConfig() error {
	file, err := os.OpenFile("./config.json", os.O_RDONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	byteValue, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	res := Config{}
	if err := json.Unmarshal(byteValue, &res); err != nil {
		return err
	}
	if err = c.validateConfig(res); err != nil {
		return err
	}
	c.Config = res
	return nil
}
