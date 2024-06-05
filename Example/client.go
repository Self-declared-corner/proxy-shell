// 2024 The Corner. This software is using GPL-3.0 licence. Licence can be found in the LICENCE file, one directory up.

package main

import (
	"fmt"
	"github.com/self-declared-corner/proxy-shell/Client"
	"github.com/self-declared-corner/proxy-shell/Server"
	"log"
	"net/url"
	"os"
	"time"
)

var (
	serverURL, _     = url.Parse(os.Args[1])
	checkDuration, _ = time.ParseDuration("1s")
)

func main() {
	currentMachine := Client.LocalMachine{Stdout: os.Stdout, Stdin: os.Stdin, Stderr: os.Stderr}
	err := currentMachine.GetLocalAddr()
	if err != nil {
		log.Fatal(err)
	}
	err = currentMachine.CollectInfo()
	if err != nil {
		log.Fatal(err)
	}
	server := Server.RemoteMachine{URL: *serverURL}
	err = server.IsAlive(checkDuration)
	if err != nil {
		log.Fatal(err)
	}
	output, err := currentMachine.SendCommand("neofetch", server)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output)
}
