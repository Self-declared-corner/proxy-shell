package proxy_shell

import (
	"encoding/json"
	"os"
)

func ReadConfig(file string) (error, *Config) {
	var config Config
	data, err := os.ReadFile(file)
	if err != nil {
		return err, nil
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		return err, nil
	}
	return nil, &config
}
