package disk

import (
	"fmt"
	"log/slog"
	"os/exec"
)

func MakeExt4DiskFormatting(devPath, uuid string) error {
	if devPath == "" {
		return fmt.Errorf("devPath is empty")
	}

	args := []string{"-F", devPath}
	if uuid != "" {
		args = append(args, "-U", uuid)
	}
	cmd := exec.Command("mkfs.ext4", args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to format %s. err: %v", devPath, err)
	}
	slog.Info("formatted disk", "devPath", devPath, "uuid", uuid, "output", string(output))
	return nil
}
