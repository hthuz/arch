package main

import (
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"
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

	x.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config changed")
	})

	x.WatchConfig()

	fmt.Println(x.AllSettings())

	time.Sleep(10 * time.Second)

	fmt.Println(x.AllSettings())

}
