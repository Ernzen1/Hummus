package config

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Paths    []string `mapstructure:"paths"`
	Tick     int      `mapstructure:"tick"`
	Location []string `mapstructure:"location"`
}

var AppConfig Config

func LoadConfig() error {

	userpath, _ := os.UserHomeDir()
	viper.SetDefault("paths", []string{userpath})
	viper.SetDefault("tick", 3600)
	viper.SetDefault("location", []string{})

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath(userpath + "/.config/humus")
	viper.AddConfigPath(userpath)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("No config file found")
		} else {
			return fmt.Errorf("Error reading config file: %s", err)
		}
	}
	err := viper.Unmarshal(&AppConfig)
	if err != nil {
		return fmt.Errorf("Unable to decode into struct, %v", err)
	}
	log.Println("Config loaded successfully")
	log.Printf("Configuração carregada com sucesso.")
	return nil
}
