package proxy_shell

import (
	"encoding/json"
	"os"
)

func ReadConfig(file string) (*Config, error) {
	var config Config
	data, err := os.ReadFile(file)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
func CreateConfig(name string) (*os.File, error) {
	file, err := os.Create(name + ".psCfg") // The name parameter must be without file extension.
	if err != nil {
		return nil, err
	}
	return file, nil
}
