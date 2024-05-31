// 2024 The Corner. This software is using GPL-3.0 licence. Licence can be found in the LICENCE file.

package latest

type Cmd string

type Config struct {
	LocalURL  string `json:"localURL"`
	RemoteURL string `json:"remoteURL"`
	Version   string `json:"version"`
}
