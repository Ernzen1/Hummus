package backupEngine

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"
	"log"
	"os"
)

func createSHA256Hash(data []byte) []byte {
	hasher := sha256.New()
	hasher.Write(data)
	return hasher.Sum(nil)
}

func BackupHandler(sources, destinations []string) error {

	for _, source := range sources {
		log.Printf("Backing up file %s", source)
		sourceData, err := os.ReadFile(source)

		if err != nil {
			log.Printf("Error reading from file: %s %e", source, err)
		}

		sourcehash := createSHA256Hash(sourceData)

		for _, destination := range destinations {
			err = backupFile(source, destination, sourceData, sourcehash)
			if err != nil {
				log.Printf("Error creating backup file: %s %e", source, err)
			}
		}
	}
	return nil
}

func backupFile(source, destination string, sourceContent []byte, sourceHash []byte) error {

	destData, err := os.ReadFile(destination)

	if err != nil {

		if errors.Is(err, os.ErrNotExist) {
			log.Printf("Destionation files %s does not exist", destination)
		} else {
			return fmt.Errorf("Error reading from file: %s %e", destination, err)
		}
	} else {
		destHash := createSHA256Hash(destData)
		log.Printf("Backing up file %s to %s", source, destination)

		if bytes.Equal(sourceHash, destHash) {
			return nil // no changes in the file.
		}
	}

	log.Printf("Data has changed!")

	if err := os.WriteFile(destination, sourceContent, 0644); err != nil {
		return fmt.Errorf("Error writing to file: %s %e", destination, err)
	}

	log.Printf("Successfully backed up file %s to %s", source, destination)
	return nil
}
