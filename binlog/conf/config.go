package conf

import (
	"fmt"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("conf") // name of config file (without extension)
	viper.SetConfigType("json")   //

	viper.AddConfigPath(".")               // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func GetChildConf(str string) string {
	return viper.GetString(str)
}
