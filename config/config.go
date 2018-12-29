package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

func Initialize(fileName string) error {
	splits := strings.Split(filepath.Base(fileName), ".")
	viper.SetConfigName(filepath.Base(splits[0]))
	viper.AddConfigPath(filepath.Dir(fileName))
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}

func checkKey(key string) {
	if !viper.IsSet(key) {
		fmt.Printf("Configuration key %s not found; aborting \n", key)
		os.Exit(1)
	}
}

func MustGetString(key string) string {
	checkKey(key)
	return viper.GetString(key)
}

func MustGetInt(key string) int {
	checkKey(key)
	return viper.GetInt(key)
}

func MustGetBool(key string) bool {
	checkKey(key)
	return viper.GetBool(key)
}

func GetString(key string) string {
	return viper.GetString(key)
}
