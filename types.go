package proxy_shell

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
)

type Cmd string
type CmdJSON struct {
	command string
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
