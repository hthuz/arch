package main

import (
	"fmt"

	"github.com/spf13/viper"
)

func main() {
	x := viper.New()
	x.SetConfigName("config")
	x.SetConfigType("yaml")

	x.AddConfigPath(".")
	if err := x.ReadInConfig(); err != nil {
		fmt.Println(err)
		return
	}
	// Read from viper
	// Transfer into a map
	fmt.Println(x.Get("open"))
	x.Set("open.0x8888", "open")
	x.Set("notopen.0x8888", "open")
	fmt.Println(x.AllSettings())

	// write to config
	x.WriteConfig()
}
