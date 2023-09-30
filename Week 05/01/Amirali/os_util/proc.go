package os_util

import (
	"io"
	"os/exec"
)

func Run(command string, outputWriter io.Writer) error {
	out, err := exec.Command("bash", "-c", command).Output()

	if err != nil {
		return err
	}
	outputWriter.Write(out)

	return nil
}
