package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	DSN       string
	JWTSecret []byte
	Admin     struct {
		Username string
		Password string
	}
}

// C holds global config
var C Config

var cIsLoaded bool = false

func loadConfig() {
	viper.SetConfigName("config")   // name of config file (without extension)
	viper.AddConfigPath("./config") // optionally look for config in the working directory
	err := viper.ReadInConfig()     // Find and read the config file

	if err != nil {
		panic(fmt.Errorf("Fatal error reading config file: %s", err))
	}

	C.DSN = viper.GetString("postgresql.dsn")
	C.JWTSecret = []byte(viper.GetString("jwt.secret"))
	C.Admin.Username = viper.GetString("admin.username")
	C.Admin.Password = viper.GetString("admin.password")
	cIsLoaded = true
}

// Get returns the config
func Get() *Config {
	if cIsLoaded == false {
		// load config
		loadConfig()
	}

	return &C
}
