package backupEngine

import (
	"bytes"
	"crypto/sha256"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
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
			log.Printf("Error reading from file: %s %s", source, err)
			continue
		}

		sourcehash := createSHA256Hash(sourceData)
		sourceFilename := filepath.Base(source)

		for _, destination := range destinations {

			if err := os.MkdirAll(destination, 0755); err != nil {
				log.Printf("Error creating directory: %s %s", destination, err)
				continue
			}

			fulldestination := filepath.Join(destination, sourceFilename)
			err = backupFile(source, fulldestination, sourceData, sourcehash)
			if err != nil {
				log.Printf("Error creating backup file: %s %s", source, err)
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
			return fmt.Errorf("error reading from file '%s':  %w", destination, err)
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
		return fmt.Errorf("error writing to folder '%s':  %w", destination, err)
	}

	log.Printf("Successfully backed up file %s to %s", source, destination)
	return nil
}
