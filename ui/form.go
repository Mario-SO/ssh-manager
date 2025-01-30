package ui

import (
	"fmt"
	"os"

	"github.com/Mario-SO/ssh-manager/ssh"
	"github.com/charmbracelet/huh"
)

func ShowKeySelectionForm(files []os.DirEntry) (string, error) {
	var selectedKey string

	options := makeFormOptions(ssh.GetPublicKeys(files))

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Choose a SSH key to copy").
				Options(options...).
				Value(&selectedKey),
		),
	)

	fmt.Printf("\x1bc") // Clear screen

	err := form.Run()
	if err != nil {
		return "", err
	}

	return selectedKey, nil
}

func makeFormOptions(keys []string) []huh.Option[string] {
	var options []huh.Option[string]
	for _, key := range keys {
		options = append(options, huh.NewOption(key, key))
	}
	return options
}
