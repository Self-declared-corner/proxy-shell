package proxy_shell

func RunCommand(commands string) error {
	err := ExecCommand(commands)
	if err != nil {
		return err
	}
	return nil
}
