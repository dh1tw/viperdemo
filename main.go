package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("ars")            // name of config file (without extension)
	viper.AddConfigPath("$HOME/.appname") // call multiple times to add many search paths
	viper.AddConfigPath(".")              // optionally look for config in the working directory
	err := viper.ReadInConfig()           // Find and read the config file
	if err != nil {                       // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	if viper.IsSet("serial.baudrate") {
		fmt.Println("baudrate:", viper.GetInt("serial.baudrate"))
	} else {
		fmt.Println("serial.baudrate not found in config file")
	}

	if viper.IsSet("serial.port") {
		fmt.Println("port:", viper.GetString("serial.port"))
	} else {
		fmt.Println("serial.port not found in config file")
	}
}
