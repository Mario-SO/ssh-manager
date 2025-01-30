package main

import (
	"fmt"
	"log"

	"github.com/Mario-SO/ssh-manager/ssh"
	"github.com/Mario-SO/ssh-manager/ui"
)

func main() {
	files := ssh.GetSSHKeys()
	key, err := ui.ShowKeySelectionForm(files)
	if err != nil {
		log.Fatal(err)
	}

	err = ssh.CopyKeyToClipboard(key)
	if err != nil {
		log.Fatal("Failed to copy to clipboard:", err)
	}

	fmt.Printf("SSH key '%s' has been copied to clipboard!\n", key)
}
