package ssh

import (
	"log"
	"os"
	"path"
)

func GetSSHKeys() []os.DirEntry {
	files, err := os.ReadDir("/Users/mario/.ssh")
	if err != nil {
		log.Fatal("SSH directory not found: ", err)
	}
	return files
}

func GetPublicKeys(files []os.DirEntry) []string {
	var publicKeys []string
	for _, file := range files {
		filename := file.Name()
		if path.Ext(filename) == ".pub" {
			publicKeys = append(publicKeys, filename)
		}
	}
	return publicKeys
}
