package proxy_shell

import (
	"github.com/rs/zerolog"
	"os"
)

func WriteLog(request LogRequest, file *os.File) error {
	if file != nil {
		err, log := CreateLog(file)
		if err != nil {
			return err
		}
		zerolog.SetGlobalLevel(request.Level)
		log.Error().Bool(request.BoolName, request.BoolValue).Msg(request.Message)
		log.Debug().Bool(request.BoolName, request.BoolValue).Msg(request.Message)
		log.Fatal().Bool(request.BoolName, request.BoolValue).Msg(request.Message)
		log.Info().Bool(request.BoolName, request.BoolValue).Msg(request.Message)
	}
	return nil
}
