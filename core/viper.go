package core

import (
	"gin-web/global"
	"fmt"

	"github.com/spf13/viper"
)

func Viper() *viper.Viper {
	v := viper.New()
	v.SetConfigFile("./config.yaml")
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Sprintf("read config err:%+v\n", err))
	}
	if err = v.Unmarshal(&global.CONFIG); err != nil {
		fmt.Println(err)
	}
	return v
}
