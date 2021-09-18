package huwlte

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	t.Run("Minimal", func(t *testing.T) {
		c := NewClient("http://192.168.1.8:8080")
		assert.NotNil(t, c, "client is not created")
		assert.NotNil(t, c.getDoer(), "client is not created")
	})

	t.Run("Extended", func(t *testing.T) {
		myClient := &http.Client{}

		c := NewClient("http://192.168.8.8",
			WithDoer(myClient),
		)

		assert.NotNil(t, c, "client is not created")
		assert.Equal(t, myClient, c.getDoer(), "client Doer is not set")
	})
}
