package lib

import (
	"encoding/json"
	"io"
	"os"
)

type Config struct {
	CACHE   string `json:"cache"`
	ROAMING string `json:"roaming"`
}

func GetConfig() (*map[string]interface{}, error) {
	file, err := os.OpenFile("./config.json", os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	byteValue, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	res := make(map[string]interface{})
	if err := json.Unmarshal(byteValue, &res); err != nil {
		return nil, err
	}
	return &res, nil
}
