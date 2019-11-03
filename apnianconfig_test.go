package apnian_go

import (
	"crypto/ecdsa"
	"github.com/sideshow/apns2/token"
	. "github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestApnianConfig(t *testing.T) {
	apnianConfigurer := ApnianConfigurer{"apnian.example", "files/test"}

	t.Run("GetApnianConfig", func(t *testing.T) {
		sut, err := GetApnianConfig("apnian.example")

		Nil(t, err)
		IsType(t, &ApnianConfig{}, sut)
	})

	t.Run("getApnianConfig", func(t *testing.T) {
		sut, err := GetApnianConfig("apnian.example")

		Nil(t, err)
		IsType(t, &ApnianConfig{}, sut)
		NotEmpty(t, sut.P8KeyName)
		NotEmpty(t, sut.Topic)
		NotEmpty(t, sut.APNSKeyID)
		NotEmpty(t, sut.TeamID)
	})

	t.Run("getApnianConfig GOROOT/config path", func(t *testing.T) {
		ac := ApnianConfigurer{"apnian.example.pathtest", "files/test"}
		sut, err := ac.getApnianConfig()

		Nil(t, err)
		IsType(t, &ApnianConfig{}, sut)
		NotEmpty(t, sut.P8KeyName)
		NotEmpty(t, sut.Topic)
		NotEmpty(t, sut.APNSKeyID)
		NotEmpty(t, sut.TeamID)
	})

	t.Run("getApnianConfig bad config name returns error", func(t *testing.T) {
		ac := ApnianConfigurer{"apnian.example.nope", "."}
		sut, err := ac.getApnianConfig()

		Nil(t, sut)
		Error(t, err)
	})

	t.Run("getApnianConfig bad config file returns error", func(t *testing.T) {
		ac := ApnianConfigurer{"apnian.badexample", "../files/test"}
		sut, err := ac.getApnianConfig()

		Nil(t, sut)
		Error(t, err)
	})

	t.Run("AuthKeyPath()", func(t *testing.T) {
		sut, err := apnianConfigurer.getApnianConfig()

		keyPath := sut.AuthKeyPath()

		info, err := os.Stat(keyPath)
		Nil(t, err)
		Equal(t, info.Name(), sut.P8KeyName)
	})

	t.Run("AuthKey()", func(t *testing.T) {
		sut, err := apnianConfigurer.getApnianConfig()

		authKey, err := sut.AuthKey()

		Nil(t, err)
		IsType(t, &ecdsa.PrivateKey{}, authKey)
		NotNil(t, authKey)
	})

	t.Run("Token()", func(t *testing.T) {
		sut, err := apnianConfigurer.getApnianConfig()

		toke, err2 := sut.Token()

		Nil(t, err)
		Nil(t, err2)
		IsType(t, &token.Token{}, toke)
		Equal(t, sut.APNSKeyID, toke.KeyID)
		Equal(t, sut.TeamID, toke.TeamID)
	})

	t.Run("Notification()", func(t *testing.T) {
		deviceID := "123456"
		sut, err := apnianConfigurer.getApnianConfig()
		payload := testAPS()

		notification := sut.Notification(deviceID, payload)

		Nil(t, err)
		Equal(t, payload.ToJsonBytes(), notification.Payload)
		//notification := sut.Notification(deviceId)
		//
		//assert.Nil(t, err)
	})
}
