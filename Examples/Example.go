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
func CreateConfig(name string) error {
	addr, _ := GetIP()
	config := proxyshell.Config{LocalURL: addr.String(), Version: "1"}
	directory, err := os.Getwd()
	if err != nil {
		return err
	}
	err = config.CreateConfig("config", directory)
	if err != nil {
		return err
	}
	return nil
}
func main() {
	app := fiber.New(fiber.Config{DisableStartupMessage: true})
	args := os.Args[2:]
	group := app.Group("/ps")
	commands := strings.Join(args, " ")
	err := proxyshell.ServeCommand(app, group, proxyshell.Cmd(commands))
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
