package proxy_shell

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"time"
)

func CreateLog(file *os.File) (error, *zerolog.Logger) {
	if file != nil {
		output := zerolog.ConsoleWriter{Out: file, TimeFormat: time.RFC850}
		output.FormatLevel = func(i interface{}) string {
			return fmt.Sprintf("| Level: %s | proxy-shell 1.0 |", i)
		}
		logger := zerolog.New(output)
		return nil, &logger
	}
	return nil, nil
}
