package zoop

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetToken(t *testing.T) {
	token, err := client.GetToken("")
	assert.NoError(t, err)
	assert.NotEmpty(t, token.Id)
}

func TestNewCardToken(t *testing.T) {

	params := &CreditCardParams{
		HolderName:      "John Doe",
		ExpirationMonth: "10",
		ExpirationYear:  "2020",
		CardNumber:      "4539003370725497",
		SecurityCode:    "123",
	}

	token, err := client.NewCardToken(params)
	assert.NoError(t, err)
	assert.NotEmpty(t, token.Id)
	assert.NotEmpty(t, token.Card.Id)

}
