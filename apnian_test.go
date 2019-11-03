package apnian

import (
	"crypto/ecdsa"
	"github.com/sideshow/apns2/token"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestApnianConfig(t *testing.T) {
	apnianConfigurer := ApnianConfigurer{"apnian.example", "files/test"}

	t.Run("New", func(t *testing.T) {
		sut, err := New("apnian.example")

		assert.Nil(t, err)
		assert.IsType(t, &Apnian{}, sut)
	})

	t.Run("getApnian", func(t *testing.T) {
		sut, err := New("apnian.example")

		assert.Nil(t, err)
		assert.IsType(t, &Apnian{}, sut)
		assert.NotEmpty(t, sut.P8KeyName)
		assert.NotEmpty(t, sut.Topic)
		assert.NotEmpty(t, sut.APNSKeyID)
		assert.NotEmpty(t, sut.TeamID)
	})

	t.Run("getApnian GOROOT/config path", func(t *testing.T) {
		ac := ApnianConfigurer{"apnian.example.pathtest", "files/test"}
		sut, err := ac.getApnian()

		assert.Nil(t, err)
		assert.IsType(t, &Apnian{}, sut)
		assert.NotEmpty(t, sut.P8KeyName)
		assert.NotEmpty(t, sut.Topic)
		assert.NotEmpty(t, sut.APNSKeyID)
		assert.NotEmpty(t, sut.TeamID)
	})

	t.Run("getApnian bad config name returns error", func(t *testing.T) {
		ac := ApnianConfigurer{"apnian.example.nope", "."}
		sut, err := ac.getApnian()

		assert.Nil(t, sut)
		assert.Error(t, err)
	})

	t.Run("getApnian bad config file returns error", func(t *testing.T) {
		ac := ApnianConfigurer{"apnian.badexample", "../files/test"}
		sut, err := ac.getApnian()

		assert.Nil(t, sut)
		assert.Error(t, err)
	})

	t.Run("AuthKeyPath()", func(t *testing.T) {
		sut, err := apnianConfigurer.getApnian()

		keyPath := sut.AuthKeyPath()

		info, err := os.Stat(keyPath)
		assert.Nil(t, err)
		assert.Equal(t, info.Name(), sut.P8KeyName)
	})

	t.Run("AuthKey()", func(t *testing.T) {
		sut, err := apnianConfigurer.getApnian()

		authKey, err := sut.AuthKey()

		assert.Nil(t, err)
		assert.IsType(t, &ecdsa.PrivateKey{}, authKey)
		assert.NotNil(t, authKey)
	})

	t.Run("AuthKey() bad key", func(t *testing.T) {
		ac := ApnianConfigurer{"apnian.badkey", "../files/test"}
		sut, err := ac.getApnian()

		_, err2 := sut.AuthKey()

		assert.Nil(t, err)
		assert.NotNil(t, err2)
	})

	t.Run("Token()", func(t *testing.T) {
		sut, err := apnianConfigurer.getApnian()

		toke, err2 := sut.Token()

		assert.Nil(t, err)
		assert.Nil(t, err2)
		assert.IsType(t, &token.Token{}, toke)
		assert.Equal(t, sut.APNSKeyID, toke.KeyID)
		assert.Equal(t, sut.TeamID, toke.TeamID)
	})

	t.Run("Token() bad key", func(t *testing.T) {
		ac := ApnianConfigurer{"apnian.badkey", "../files/test"}
		sut, err := ac.getApnian()

		_, err2 := sut.Token()

		assert.Nil(t, err)
		assert.NotNil(t, err2)
	})

	t.Run("Notification()", func(t *testing.T) {
		deviceID := "123456"
		sut, err := apnianConfigurer.getApnian()
		payload := testAPS()

		notification := sut.Notification(deviceID, payload)

		assert.Nil(t, err)
		assert.Equal(t, payload.ToJsonBytes(), notification.Payload)
		//notification := sut.Notification(deviceId)
		//
		//assert.Nil(t, err)
	})
}
