// 2024 The Corner. This software is using GPL-3.0 licence. Licence can be found in the LICENCE file.

package proxy_shell

import (
	"github.com/fasthttp/websocket"
	"github.com/matishsiao/goInfo"
	"net"
	"net/url"
	"os"
	"strings"
)

type LocalMachine struct {
	IP          net.IP   //client's url
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
	localAddr := strings.Split(conn.LocalAddr().(*net.UDPAddr).String(), ":")
	lm.IP = net.IP(localAddr[1])
	return nil
}
func (lm LocalMachine) SendCommand(command Cmd, rm RemoteMachine) (interface{}, error) {
	var result interface{}
	serverURL := url.URL{Scheme: "ws", Host: rm.IP.String(), Path: "/ws"}
	conn, _, err := websocket.DefaultDialer.Dial(serverURL.String(), nil)
	if err != nil {
		return nil, err
	}
	go func() {
		for {
			err := conn.WriteJSON(command)
			if err != nil {
				return
			}
		}
	}()
	go func() {
		for {
			err := conn.ReadJSON(result)
			if err != nil {
				return
			}
		}
	}()

	return result, nil
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
