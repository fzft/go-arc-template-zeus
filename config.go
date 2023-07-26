package template

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strings"
)

func loadConfig() {

	// load env variable
	env := os.Getenv("env")
	path, err := os.Getwd()
	if err != nil {
		panic(fmt.Sprintf("Failed to get current working directory: %v", err))
	}

	viper.SetConfigType("yaml")

	// if the env is not set, use dev as default
	if env == "" {
		env = "dev"
	}

	viper.SetConfigFile(filepath.Join(path, "config", fmt.Sprintf("config-%s", env)))

	if err = viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("Failed to load config file: %v", err))
	}

	// override config with env variables, this is useful for docker deployment
	// e.g. env variable "db_host" will override "db.host" in config file
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()
}
