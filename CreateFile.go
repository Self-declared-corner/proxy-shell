package proxy_shell

import (
	"errors"
	"os"
)

func CreateFile(name string) (error, *os.File) {
	file, err := os.Create(name)
	if err != nil {
		return err, file
	}
	if file != nil {
		return nil, file
	}
	return errors.New("something went wrong"), nil
}
