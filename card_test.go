package zoop

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetCustomerCard(t *testing.T) {
	card, err := client.LinkCreditCard("", "")
	assert.NoError(t, err)
	assert.NotEmpty(t, card.Id)
}

func TestGetCreditCard(t *testing.T) {
	card, err := client.GetCreditCard("")
	assert.NoError(t, err)
	assert.NotEmpty(t, card.Id)
}

func TestDelCreditCard(t *testing.T) {
	del, err := client.DelCreditCard("")
	assert.NoError(t, err)
	assert.Equal(t, "", del.Id)
}
