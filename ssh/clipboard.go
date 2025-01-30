package ssh

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func CopyKeyToClipboard(key string) error {
	fullPath := filepath.Join(os.Getenv("HOME"), ".ssh", key)
	cmd := exec.Command("sh", "-c", fmt.Sprintf("cat %s | pbcopy", fullPath))
	return cmd.Run()
}
