// 2024 The Corner. This software is using GPL-3.0 licence. Licence can be found in the LICENCE file.

package Client

import (
	"github.com/fasthttp/websocket"
	"github.com/matishsiao/goInfo"
	proxyshell "github.com/self-declared-corner/proxy-shell"
	"github.com/self-declared-corner/proxy-shell/Server"
	"net"
	"net/url"
	"os"
)

type LocalMachine struct {
	URL         url.URL  //client's url
	Stdin       *os.File // stdin stream
	Stdout      *os.File //stdout stream
	Stderr      *os.File // stderr stream
	Information goInfo.GoInfoObject
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
	lm.Information.VarDump()
	return nil
}
func (lm LocalMachine) CollectInfo() error {
	gi, err := goInfo.GetInfo()
	if err != nil {
		return err
	}
	lm.Information = gi
	return nil
}
