package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAPS(t *testing.T) {
	t.Run("GenerateAPS()", func(t *testing.T) {
		alert := "Bawk bawk, bitches! 🐔"
		sound := "default"
		linkUrl := "https://electronicpanopticon.com"
		payload := Payload{
			Alert:   alert,
			Sound:   sound,
			LinkUrl: linkUrl,
		}
		expected := &APS{
			APS: payload,
		}

		aps := GenerateAPS(alert, sound, linkUrl)

		assert.Equal(t, expected, aps)
	})
}