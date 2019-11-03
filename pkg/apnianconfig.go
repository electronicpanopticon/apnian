package pkg

import (
	"fmt"
	"github.com/electronicpanopticon/gobrick"
	"github.com/spf13/viper"
)

type ApnianConfigurer struct {
	ConfigName string
	Root string
}

type ApnianConfig struct {
	P8KeyName  string
	Topic      string
	APNSKeyID  string
	TeamID     string
	Configurer ApnianConfigurer
}

func GetApnianConfig(configName string) (*ApnianConfig, error) {
	ac := ApnianConfigurer{configName, gobrick.GetGOPATH()}
	return ac.getApnianConfig()
}

func (ac ApnianConfig) AuthKeyPath() string {
	rel := fmt.Sprintf("keys/%s", ac.P8KeyName)
	return fmt.Sprintf("%s/%s", ac.Configurer.Root, rel)
}

func (ac ApnianConfigurer) getApnianConfig() (*ApnianConfig, error) {
	viper.SetConfigName(ac.ConfigName)
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.AddConfigPath(ac.Root + "/config")

	var c ApnianConfig
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}
	c.Configurer = ac
	return &c, nil
}