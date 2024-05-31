// 2024 The Corner. This software is using GPL-3.0 licence. Licence can be found in the LICENCE file.

package proxy_shell

import (
	"os"
	"os/exec"
	"strings"
)

func ExecCommand(commands string) error {
	commandAndArgs := strings.Split(commands, " ")
	cmd := exec.Command(commandAndArgs[1], strings.Join(commandAndArgs[2:], " "))
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	defer func(cmd *exec.Cmd) {
		err := cmd.Run()
		if err != nil {
			return
		}
	}(cmd)
	return nil
}
