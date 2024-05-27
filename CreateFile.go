package proxy_shell

import (
	"errors"
	"os"
)

func CreateFile(name string) (*os.File, error) {
	file, err := os.Create(name)
	if err != nil {
		return file, err
	}
	if file != nil {
		return file, nil
	}
	return nil, errors.New("something went wrong")
}
