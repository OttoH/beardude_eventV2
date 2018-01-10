package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Server   string
	Database string
}

// Read and parse the configuration file
func (c *Config) Read() {
  viper.SetConfigName("config")
  viper.SetConfigType("toml")
  viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	} else {
    c.Server = viper.GetString("Server")
    c.Database = viper.GetString("Database")
  }
}
