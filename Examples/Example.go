package main

import (
	proxyshell "github.com/Self-declared-corner/proxy-shell"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"net"
	"os"
	"strings"
)

func GetIP() (*net.IP, error) {
	conn, err := net.Dial("udp", "example.com:80")
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	addr := conn.LocalAddr().(*net.UDPAddr)
	return &addr.IP, nil
}
func main() {
	app := fiber.New(fiber.Config{DisableStartupMessage: true})
	args := os.Args[2:]
	group := app.Group("/ps")
	commands := strings.Join(args, " ")
	localAddr, _ := GetIP()
	dir, _ := os.Getwd()
	rawCfg := proxyshell.Config{LocalURL: localAddr.String(), RemoteURL: os.Args[1], Version: "1"}
	err := rawCfg.CreateConfig("config", dir)
	if err != nil {
		return
	}
	err = proxyshell.ServeCommand(app, group, proxyshell.Cmd(commands))
	if err != nil {
		request := proxyshell.LogRequest{Message: "couldn't serve a command", Level: zerolog.ErrorLevel, BoolName: "isBad", BoolValue: true}
		dir, _ := os.Getwd()
		err := os.Chdir(dir)
		if err != nil {
			panic("couldn't change directory!")
		}
		file, err := os.Open("ps.log")
		err = proxyshell.WriteLog(request, file)
		if err != nil {
			panic("couldn't write a log")
		}
	}

}
