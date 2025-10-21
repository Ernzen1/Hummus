package config

import (
	"io/ioutil"
	"log"
	"os"
)

func GetConfigMarshalled() []byte {
	configFd, err := os.Open(DefaultConfigPath)
	defer configFd.Close()
	if err != nil {
		log.Fatal("Error opening file:", err)
	}

	marshalledFd, err := ioutil.ReadAll(configFd)

	return marshalledFd
}
