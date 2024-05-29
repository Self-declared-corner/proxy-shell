package Client

import (
	"fmt"
	proxyshell "github.com/Self-declared-corner/proxy-shell"
	"github.com/Self-declared-corner/proxy-shell/Server"
	"github.com/fasthttp/websocket"
	"net"
	"net/url"
)

type LocalMachine struct {
	URL    url.URL     //client's url
	Stdin  interface{} // stdin stream
	Stdout interface{} //stdout stream
	Stderr interface{} // stderr stream
	OS     string      //linux, windows, macos (darwin)
}

func (lm LocalMachine) GetLocalAddr() error { //gets local machine's address
	conn, err := net.Dial("udp", "example.com:80")
	if err != nil {
		return err
	}
	localURL, err := url.Parse(conn.LocalAddr().String())
	if err != nil {
		return err
	}
	lm.URL = *localURL
	return nil
}
func (lm LocalMachine) SendCommand(command proxyshell.Cmd, rm Server.RemoteMachine) error {
	serverURL := url.URL{Scheme: "ws", Host: rm.URL.String(), Path: "/ws"}
	conn, _, err := websocket.DefaultDialer.Dial(serverURL.String(), nil)
	if err != nil {
		return err
	}
	go func() {
		for {
			err := conn.WriteJSON(command)
			if err != nil {
				return
			}
		}
	}()
	return nil
}
func (lm LocalMachine) PrintInfo() error {
	fmt.Println("URL: ", lm.URL, "Operating system: ", lm.OS)
	return nil
}