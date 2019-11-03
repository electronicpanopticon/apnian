package pkg

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/electronicpanopticon/gobrick"
	"github.com/mitchellh/go-homedir"
	"github.com/sideshow/apns2/token"
	"github.com/spf13/viper"
	"log"
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

// GetApnianConfig returns an ApnianConfig filed with the values in its config file.
// Locations it looks for are:
//		.
//		..
//		$GOPATH/config
//		$HOME
func GetApnianConfig(configName string) (*ApnianConfig, error) {
	ac := ApnianConfigurer{configName, gobrick.GetGOPATH()}
	return ac.getApnianConfig()
}

// AuthKeyPath returns the path to the ECDSA private key specified in the ApnianConfig file.
func (ac ApnianConfig) AuthKeyPath() string {
	rel := fmt.Sprintf("keys/%s", ac.P8KeyName)
	return fmt.Sprintf("%s/%s", ac.Configurer.Root, rel)
}

// AuthKey returns the ECDSA private key specified in the ApnianConfig file.
func (ac ApnianConfig) AuthKey() (*ecdsa.PrivateKey, error) {
	return token.AuthKeyFromFile(ac.AuthKeyPath())
}

// Token represents an Apple Provider Authentication Token (JSON Web Token) configured
// with the values from the ApnianConfig file.
func (ac ApnianConfig) Token() (*token.Token, error) {
	authKey, err := ac.AuthKey()
	if err != nil {
		return nil, err
	}
	return &token.Token{
		AuthKey:  authKey,
		KeyID:    ac.APNSKeyID,
		TeamID:   ac.TeamID,
	}, nil
}

// getApnianConfig returns an ApnianConfig from the configured Viper instance.
func (ac ApnianConfigurer) getApnianConfig() (*ApnianConfig, error) {
	ac.configureViper()

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

// configureViper
func (ac ApnianConfigurer) configureViper() {
	viper.SetConfigName(ac.ConfigName)
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	home, err := homedir.Dir()
	if err == nil {
		viper.AddConfigPath(home)
	} else {
		log.Println("unable to get homedir")
	}
	viper.AddConfigPath(ac.Root + "/config")
}