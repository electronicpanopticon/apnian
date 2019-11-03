package pkg

import (
	"github.com/electronicpanopticon/gobrick"
	"github.com/spf13/viper"
)

type ApnianConfigurer struct {
	ConfigName string
	Root string
}

type ApnianConfig struct {
	P8KeyName string
	Topic string
	APNSKeyID string
	TeamID string
}

func (ac ApnianConfigurer) GetApnianConfig() (*ApnianConfig, error) {
	return ac.getApnianConfig()
}

func (ac ApnianConfigurer) getApnianConfig() (*ApnianConfig, error) {
	viper.SetConfigName(ac.ConfigName)
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