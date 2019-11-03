package apnian

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/electronicpanopticon/gobrick"
	"github.com/mitchellh/go-homedir"
	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/token"
	"github.com/spf13/viper"
	"log"
)

type ApnianConfigurer struct {
	ConfigName string
	Root string
}

type Apnian struct {
	P8KeyName  string
	Topic      string
	APNSKeyID  string
	TeamID     string
	Configurer *ApnianConfigurer
	Client     *apns2.Client

}

// New returns an Apnian filed with the values in its config file.
// Locations it looks for are:
//		.
//		..
//		$GOPATH/config
//		$HOME
func New(configName string) (*Apnian, error) {
	apnian := ApnianConfigurer{configName, gobrick.GetGOPATH()}
	return apnian.getApnian()
}

// AuthKeyPath returns the path to the ECDSA private key specified in the Apnian file.
func (apnian Apnian) AuthKeyPath() string {
	rel := fmt.Sprintf("keys/%s", apnian.P8KeyName)
	return fmt.Sprintf("%s/%s", apnian.Configurer.Root, rel)
}

// AuthKey returns the ECDSA private key specified in the Apnian file.
func (apnian Apnian) AuthKey() (*ecdsa.PrivateKey, error) {
	return token.AuthKeyFromFile(apnian.AuthKeyPath())
}

// Token represents an Apple Provider Authentication Token (JSON Web Token) configured
// with the values from the Apnian file.
func (apnian Apnian) Token() (*token.Token, error) {
	authKey, err := apnian.AuthKey()
	if err != nil {
		return &token.Token{}, err
	}
	return &token.Token{
		AuthKey:  authKey,
		KeyID:    apnian.APNSKeyID,
		TeamID:   apnian.TeamID,
	}, nil
}

func (apnian Apnian) Notification(deviceID string, payload *APS) *apns2.Notification {
	notification := &apns2.Notification{}
	notification.DeviceToken = deviceID
	notification.Topic = apnian.Topic
	notification.Payload = payload.ToJsonBytes()
	return notification
}

// loadClient lazy loads a configured apns2.Client.
func (apnian *Apnian) loadClient() error {
	if apnian.Client == nil {
		token, err := apnian.Token()
		if err != nil {
			return err
		}
		apnian.Client = apns2.NewTokenClient(token).Development()
	}
	return nil
}

// func (apnian Apnian) Push

// getApnian returns an Apnian from the configured Viper instance.
func (ac ApnianConfigurer) getApnian() (*Apnian, error) {
	ac.configureViper()

	var c Apnian
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}
	c.Configurer = &ac
	return &c, nil
}



// configureViper
func (apnian ApnianConfigurer) configureViper() {
	viper.SetConfigName(apnian.ConfigName)
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	home, err := homedir.Dir()
	if err == nil {
		viper.AddConfigPath(home)
	} else {
		log.Println("unable to get homedir")
	}
	viper.AddConfigPath(apnian.Root + "/config")
}