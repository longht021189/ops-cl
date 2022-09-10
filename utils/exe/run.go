package exe

import (
	e "os/exec"
)

func Run(cmdArgs ...string) ([]byte, error) {
	all := prefixCmd
	all = append(all, cmdArgs...)

	cmd := e.Command(all[0], all[1:]...)
	return cmd.Output()
}
