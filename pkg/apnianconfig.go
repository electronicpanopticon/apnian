package pkg

import (
	"github.com/electronicpanopticon/gobrick"
	"github.com/spf13/viper"
)

type ApnianConfig struct {
	P8KeyName string
}

func GetApnianConfig() (*ApnianConfig, error) {
	return getApnianConfig("apnian")
}

func getApnianConfig(filename string) (*ApnianConfig, error) {
	viper.SetConfigName(filename)
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.AddConfigPath(gobrick.GetGOPATH() + "/config")

	var c ApnianConfig
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}
	return &c, nil
}