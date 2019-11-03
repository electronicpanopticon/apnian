package pkg

import (
	"crypto/ecdsa"
	"github.com/sideshow/apns2/token"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestApnianConfig(t *testing.T) {
	apnianConfigurer := ApnianConfigurer{"apnian.example", "../files/test"}

	t.Run("GetApnianConfig", func(t *testing.T) {
		sut, err := GetApnianConfig("apnian.example")

		assert.Nil(t, err)
		assert.IsType(t, &ApnianConfig{}, sut)
	})

	t.Run("getApnianConfig", func(t *testing.T) {
		sut, err := GetApnianConfig("apnian.example")

		assert.Nil(t, err)
		assert.IsType(t, &ApnianConfig{}, sut)
		assert.NotEmpty(t, sut.P8KeyName)
		assert.NotEmpty(t, sut.Topic)
		assert.NotEmpty(t, sut.APNSKeyID)
		assert.NotEmpty(t, sut.TeamID)
	})

	t.Run("getApnianConfig GOROOT/config path", func(t *testing.T) {
		ac := ApnianConfigurer{"apnian.example.pathtest", "../files/test"}
		sut, err := ac.getApnianConfig()

		assert.Nil(t, err)
		assert.IsType(t, &ApnianConfig{}, sut)
		assert.NotEmpty(t, sut.P8KeyName)
		assert.NotEmpty(t, sut.Topic)
		assert.NotEmpty(t, sut.APNSKeyID)
		assert.NotEmpty(t, sut.TeamID)
	})

	t.Run("getApnianConfig bad config name returns error", func(t *testing.T) {
		ac := ApnianConfigurer{"apnian.example.nope", "."}
		sut, err := ac.getApnianConfig()

		assert.Nil(t, sut)
		assert.Error(t, err)
	})

	t.Run("getApnianConfig bad config file returns error", func(t *testing.T) {
		ac := ApnianConfigurer{"apnian.badexample", "../files/test"}
		sut, err := ac.getApnianConfig()

		assert.Nil(t, sut)
		assert.Error(t, err)
	})

	t.Run("AuthKeyPath()", func(t *testing.T) {
		sut, err := apnianConfigurer.getApnianConfig()

		keyPath := sut.AuthKeyPath()

		info, err := os.Stat(keyPath)
		assert.Nil(t, err)
		assert.Equal(t, info.Name(), sut.P8KeyName)
	})

	t.Run("AuthKey()", func(t *testing.T) {
		sut, err := apnianConfigurer.getApnianConfig()

		authKey, err := sut.AuthKey()

		assert.Nil(t, err)
		assert.IsType(t, &ecdsa.PrivateKey{}, authKey)
		assert.NotNil(t, authKey)
	})

	t.Run("Token()", func(t *testing.T) {
		sut, err := apnianConfigurer.getApnianConfig()

		toke, err := sut.Token()

		assert.Nil(t, err)
		assert.IsType(t, &token.Token{}, toke)
		assert.Equal(t, sut.APNSKeyID, toke.KeyID)
		assert.Equal(t, sut.TeamID, toke.TeamID)
	})

	t.Run("Notification()", func(t *testing.T) {
		//deviceId := "123456"
		//sut, err := apnianConfigurer.getApnianConfig()

		//notification := sut.Notification(deviceId)
		//
		//assert.Nil(t, err)
	})
}
