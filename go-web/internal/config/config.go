package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

// Init config
func Init(confPath string) error{
	err:=initConfig(confPath)
	if err!=nil{
		return err
	}
	return nil
}

func initConfig(confPath string) error{
	if confPath!=""{
		viper.SetConfigFile(confPath)
	}else{
		viper.AddConfigPath("conf")
		viper.SetConfigName("config")
	}
	viper.SetConfigType("json")
	viper.AutomaticEnv()
	if err:=viper.ReadInConfig();err!=nil{
		return errors.WithStack(err)
	}

	return nil
}