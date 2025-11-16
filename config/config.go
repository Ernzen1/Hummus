package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

type Config struct {
	Paths       []string `mapstructure:"paths"`
	Tick        int      `mapstructure:"tick"`
	Location    []string `mapstructure:"location"`
	ServiceName string   `mapstructure:"service_name"`
	LogFile     string   `mapstructure:"logfile"`
}

var AppConfig Config

func LoadConfig() error {

	userpath, _ := os.UserHomeDir()

	userhummus := filepath.Join(userpath, ".hummus")
	configDir := filepath.Join(userpath, ".config", "hummus")
	configFilePath := filepath.Join(configDir, "config.yaml")
	defaultLogFile := filepath.Join(configDir, "hummus.log")

	// defaults
	viper.SetDefault("paths", []string{})
	viper.SetDefault("tick", 10)
	viper.SetDefault("location", []string{userhummus})
	viper.SetDefault("service_name", "hummus")
	viper.SetDefault("logfile", defaultLogFile)

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configDir)
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("No config file found")

			if err := os.MkdirAll(configDir, 0755); err != nil {
				return fmt.Errorf("error creating config dir: %s", err)
			}

			if err := viper.WriteConfigAs(configFilePath); err != nil {
				return fmt.Errorf("error writing config file: %s", err)
			}
		} else {
			return fmt.Errorf("error reading config file: %s", err)
		}
	} else {
		log.Printf("config file successfully loaded: %s", configFilePath)
	}
	err := viper.Unmarshal(&AppConfig)
	if err != nil {
		return fmt.Errorf("unable to decode into struct, %s", err)
	}
	log.Println("Config loaded successfully")
	return nil
}
