package Server

import (
	proxyshell "github.com/Self-declared-corner/proxy-shell"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"net/url"
	"time"
)

type RemoteMachine struct {
	URL   url.URL
	Alive bool
}

func (rm RemoteMachine) IsAlive(duration time.Duration) error {
	tick := time.Tick(duration)
	go func() {
		for range tick {
			if resp, err := http.Get(rm.URL.String()); err != nil || resp.StatusCode >= 500 {
				rm.Alive = false
			}
		}
	}()
	return nil
}
func (rm RemoteMachine) ListenForCommands(app *fiber.App) (*proxyshell.Cmd, error) {
	var command proxyshell.Cmd
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
		for {
			err := conn.ReadJSON(command)
			if err != nil {
				break
			}
		}
	}, websocket.Config{ReadBufferSize: 2048}))
	return &command, nil
}