package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Load - loads configuration
func Load(fileName, fileType string, paths ...string) error {
	viper.SetConfigName(fileName)
	viper.SetConfigType(fileType)
	for _, path := range paths {
		viper.AddConfigPath(path)
	}
	return viper.ReadInConfig()
	// viper.AutomaticEnv()
}

// ReadEnvString - reads variable from environment
func ReadEnvString(key string) string {
	checkIfSet(key)
	return viper.GetString(key)
}

// ReadEnvInt - reads int variable from environment
func ReadEnvInt(key string) int {
	checkIfSet(key)
	return viper.GetInt(key)
}

func checkIfSet(key string) {
	if !viper.IsSet(key) {
		err := fmt.Errorf("Key %s is not set", key)
		panic(err)
	}
}
