package core

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func InitViper(path string, config any) *viper.Viper {
	// Init viper
	v := viper.New()
	v.SetConfigFile(path)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err = v.Unmarshal(config); err != nil {
			fmt.Println(err)
		}
	})
	err = v.Unmarshal(config)
	if err != nil {
		panic(err)
	}
	return v
}
