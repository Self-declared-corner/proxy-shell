// 2024 The Corner. This software is using GPL-3.0 licence. Licence can be found in the LICENCE file.

package proxy_shell

import (
	"github.com/creack/pty"
	"os"
	"os/exec"
	"strings"
)

func ExecCommand(commands string) ([]byte, error) {
	var result []byte
	commandAndArgs := strings.Split(commands, " ")
	cmd := exec.Command(commandAndArgs[1], strings.Join(commandAndArgs[2:], " "))
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	defer func(cmd *exec.Cmd) {
		file, err := pty.Start(cmd)
		if err != nil {
			return
		}
		_, err = file.Read(result)
		if err != nil {
			return
		}
	}(cmd)
	return result, nil
}
