package zoop

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddBuyer(t *testing.T) {

	params := &BuyerParams{
		FirstName:  "John",
		LastName:   "Doe",
		Email:      "john.doe@email.com",
		TaxpayerId: "00000000000",
	}

	buyer, err := client.NewBuyer(params)
	assert.NoError(t, err)
	assert.NotEmpty(t, buyer.Id)
	assert.Equal(t, buyer.Email, params.Email)

}

func TestGetBuyer(t *testing.T) {
	buyer, err := client.GetBuyer("")
	assert.NoError(t, err)
	assert.Equal(t, buyer.Id, "")
}

func TestSetBuyer(t *testing.T) {

	params := &BuyerParams{
		PhoneNumber: "999999999",
	}

	buyer, err := client.SetBuyer("", params)
	assert.NoError(t, err)
	assert.Equal(t, buyer.Id, "")
	assert.Equal(t, buyer.PhoneNumber, params.PhoneNumber)

}

func TestListBuyers(t *testing.T) {
	list, err := client.ListBuyer()
	assert.NoError(t, err)
	assert.NotZero(t, len(list.Buyers))
	spew.Dump(list)
}

func TestDelBuyer(t *testing.T) {
	del, err := client.DelBuyer("")
	assert.NoError(t, err)
	assert.Equal(t, "", del.Id)
}
