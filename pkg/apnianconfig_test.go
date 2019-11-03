package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApnianConfig(t *testing.T) {
	t.Run("getApnianConfig", func(t *testing.T) {
		sut, err := getApnianConfig("apnian.example")

		assert.Nil(t, err)
		assert.IsType(t, &ApnianConfig{}, sut)
		assert.NotNil()
	})

	t.Run("getApnianConfig bad config name returns error", func(t *testing.T) {
		_, err := getApnianConfig("apnian.example.nope")

		assert.Error(t, err)
	})
}
