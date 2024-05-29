package proxy_shell

type Cmd string

type Config struct {
	LocalURL  string `json:"localURL"`
	RemoteURL string `json:"remoteURL"`
	Version   string `json:"version"`
}
