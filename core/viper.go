package core

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"scry/global"
)

const (
	ConfigEnv  = "CONFIG"
	ConfigFile = "config.yml"
)

func Viper(path ...string) *viper.Viper {
	v := viper.New()
	v.SetConfigFile(ConfigFile)
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	fmt.Println(v)
	v.OnConfigChange(func(e fsnotify.Event) {
		var c map[string]interface{}
		fmt.Println("config file changed:")
		if err := v.Unmarshal(&c); err != nil {
			fmt.Println(err)
		}
	})
	if err := v.Unmarshal(&global.Config); err != nil {
		fmt.Println(err)
	}
	return v
}
