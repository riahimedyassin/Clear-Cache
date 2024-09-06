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

func GetConfig() (*Config, error) {
	file, err := os.OpenFile("./config.json", os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	byteValue, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	res := Config{}
	if err := json.Unmarshal(byteValue, &res); err != nil {
		return nil, err
	}
	if err = ValidateConfig(res); err != nil {
		return nil, err
	}
	return &res, nil
}

func ValidateConfig(config Config) error {
	if config.ROAMING == "" || config.CACHE == "" {
		return errors.New("invalid config")
	}
	return nil
}

func InitConfig() {

}
