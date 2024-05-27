package proxy_shell

import (
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"time"
)

func CreateLog(file *os.File) (*zerolog.Logger, error) {
	if file != nil {
		output := zerolog.ConsoleWriter{Out: file, TimeFormat: time.RFC850}
		output.FormatLevel = func(i interface{}) string {
			return fmt.Sprintf("| Level: %s | thanks for using this good (actually bad) software! |", i)
		}
		logger := zerolog.New(output)
		return &logger, nil
	}
	return nil, nil
}
