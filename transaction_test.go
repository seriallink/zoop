package zoop

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetTransaction(t *testing.T) {
	transaction, err := client.GetTransaction("")
	assert.NoError(t, err)
	assert.Equal(t, transaction.Id, "")
}

func TestNewTransactionWithToken(t *testing.T) {

	params := &TransactionParams{
		Amount:      100,
		Currency:    "BRL",
		Description: "It's a token test!",
		PaymentType: TypeCredit,
		OnBehalfOf:  client.SellerId,
		Capture:     true,
		SourceUsage: SingleUse,
		SourceType:  SourceToken,
		Token:       "",
	}

	transaction, err := client.NewTransaction(params)
	assert.NoError(t, err)
	assert.NotEmpty(t, transaction.Id)

}

func TestNewTransactionWithVault(t *testing.T) {

	params := &TransactionParams{
		Amount:      100,
		Currency:    "BRL",
		Description: "It's a vault test!",
		PaymentType: TypeCredit,
		OnBehalfOf:  client.SellerId,
		Capture:     true,
		SourceUsage: Reusable,
		SourceType:  SourceCard,
		Customer:    "",
	}

	transaction, err := client.NewTransaction(params)
	assert.NoError(t, err)
	assert.NotEmpty(t, transaction.Id)

}

func TestNewTransactionWithTicket(t *testing.T) {

	params := &TransactionParams{
		Amount:      300,
		Currency:    "BRL",
		Description: "It's a ticket test!",
		PaymentType: TypeBoleto,
		OnBehalfOf:  client.SellerId,
		SourceUsage: SingleUse,
		SourceType:  SourceCustomer,
		PaymentMethod: &PaymentMethodParams{
			ExpirationDate: NewDate(time.Now().Add(time.Hour * 48)).Pointer(),
			TopInstructions: []string{
				"- SR. CAIXA, NÃO RECEBER APÓS O VENCIMENTO",
			},
			BodyInstructions: []string{
				"- PEDIDO 1234567890",
				"",
			},
		},
		Customer: "",
	}

	transaction, err := client.NewTransaction(params)
	assert.NoError(t, err)
	assert.NotEmpty(t, transaction.Id)
	assert.NotNil(t, transaction.PaymentMethod)
	assert.NotEmpty(t, transaction.PaymentMethod.Url)

}

func TestCaptureTransaction(t *testing.T) {
	transaction, err := client.CaptureTransaction("", 50)
	assert.NoError(t, err)
	assert.Equal(t, transaction.Id, "")
}

func TestRefundTransaction(t *testing.T) {
	transaction, err := client.RefundTransaction("", 50)
	assert.NoError(t, err)
	assert.Equal(t, transaction.Id, "")
}

func TestListMarketplaceTransactions(t *testing.T) {
	list, err := client.ListMarketplaceTransactions()
	assert.NoError(t, err)
	assert.NotZero(t, len(list.Transactions))
	spew.Dump(list)
}

func TestListSellerTransactions(t *testing.T) {
	list, err := client.ListSellerTransactions()
	assert.NoError(t, err)
	assert.NotZero(t, len(list.Transactions))
	spew.Dump(list)
}
