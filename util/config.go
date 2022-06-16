package util

import (
	"fmt"

	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

type configType struct {
	Mongodb  map[string]interface{}
	isloaded bool
}

var config configType

func LoadConfig() {
	if !config.isloaded {
		viper.AddConfigPath(".")
		viper.SetConfigName("app")
		if err := viper.ReadInConfig(); err != nil {
			fmt.Printf("Error reading config file, %s", err)
		}
		viper.SetEnvPrefix("global")
		runmode := cast.ToString(viper.Get("runmode"))
		config.Mongodb = viper.Get(runmode + ".mongodb").(map[string]interface{})
		config.isloaded = true
	} else {
		//do nothig Just Chill !!!
	}
}

func GetConfig() map[string]interface{} {
	LoadConfig()
	return config.Mongodb
}
