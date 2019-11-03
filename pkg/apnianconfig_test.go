package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApnianConfig(t *testing.T) {
	ac := ApnianConfigurer{"apnian.example", "."}

	t.Run("GetApnianConfig", func(t *testing.T) {
		sut, err := ac.GetApnianConfig()

		assert.Nil(t, err)
		assert.IsType(t, &ApnianConfig{}, sut)
	})

	t.Run("getApnianConfig", func(t *testing.T) {
		sut, err := ac.GetApnianConfig()

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
		ac := ApnianConfigurer{"apnian.badexample", "."}
		sut, err := ac.getApnianConfig()

		assert.Nil(t, sut)
		assert.Error(t, err)
	})


}
