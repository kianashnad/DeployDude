package foundation

import (
	"os/exec"
)

func RunCmd(commands string) ([]byte, error) {
	out, err := exec.Command("bash", "-c", commands).Output()
	if err != nil {
		return nil, err
	}
	return out, nil
}
