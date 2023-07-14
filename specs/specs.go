package specs

import "os/exec"

func Command(command string) string {
	cmd := exec.Command("sh", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		return err.Error()
	}
	return string(output)
}
