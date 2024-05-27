package proxy_shell

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"sync"
)

type Cmd string

type Config struct {
	LocalURL string `json:"localURL"`
	Version  string `json:"version"`
}
type ProxyRequest struct {
	Command Cmd
	SendTo  url.URL
}
type LogRequest struct {
	Message   string
	Level     zerolog.Level
	BoolName  string
	BoolValue bool
}
type Client struct {
	Url         url.URL
	ConnectedTo string
	Mutex       sync.Mutex
	Alive       bool
}

func (request ProxyRequest) RunProxyRequest(w *http.ResponseWriter, r *http.Request, command Cmd, group *fiber.Group) error {
	request.Command = command
	if group != nil {
		err := ServeCommand(*group, request.Command)
		if err != nil {
			return err
		}
	}
	if w != nil && r != nil {
		proxy := httputil.NewSingleHostReverseProxy(&request.SendTo)
		proxy.ServeHTTP(*w, r)
	}
	return nil
}
func (rawCfg Config) CreateConfig(name string, dir string) error {
	err := os.Chdir(dir)
	if err != nil {
		return err
	}
	data, err := json.Marshal(rawCfg)
	err = os.WriteFile(name+".psCfg", data, 0666)
	if err != nil {
		return err
	}
	return nil
}
