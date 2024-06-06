// 2024 The Corner. This software is using GPL-3.0 licence. Licence can be found in the LICENCE file.

package proxy_shell

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"log"
	"net"
	"net/http"
	"time"
)

type RemoteMachine struct {
	IP    net.IP
	Alive bool
}

func (rm RemoteMachine) IsAlive(duration time.Duration) error {
	tick := time.Tick(duration)
	go func() {
		for range tick {
			if resp, err := http.Get(rm.IP.String()); err != nil || resp.StatusCode >= 500 {
				rm.Alive = false
			}
		}
	}()
	return nil
}
func (rm RemoteMachine) ListenForCommands() (*Cmd, error) {
	app := fiber.New(fiber.Config{DisableStartupMessage: true})
	var command Cmd
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
	app.Post("/ws/:cmd", websocket.New(func(conn *websocket.Conn) {
		defer func(conn *websocket.Conn) {
			err := conn.Close()
			if err != nil {
				panic(err)
			}
		}(conn)
		go func() {
			for {
				err := conn.ReadJSON(command)
				if err != nil {
					break
				}
			}
		}()
		result, err := ExecCommand(string(command))
		if err != nil {
			return
		}
		defer func() {
			for {
				err := conn.WriteJSON(result)
				if err != nil {
					break
				}
			}
		}()
	}, websocket.Config{ReadBufferSize: 2048}))
	go func() {
		log.Fatal(app.Listen(":21507"))
	}()
	return &command, nil
}
