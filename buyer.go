package zoop

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	buyerPath   = "marketplaces/%s/buyers"
	buyerPathId = "marketplaces/%s/buyers/%s"
)

func (c *Client) GetBuyer(id string) (*Buyer, error) {
	model := new(Buyer)
	err := c.Get(fmt.Sprintf(buyerPathId, c.MarketPlaceId, id), nil, nil, model)
	return model, err
}

func (c *Client) AddBuyer(params *BuyerParams) (*Buyer, error) {
	model := new(Buyer)
	err := c.Post(fmt.Sprintf(buyerPath, c.MarketPlaceId), params, nil, model)
	return model, err
}

func (c *Client) SetBuyer(id string, params *BuyerParams) (*Buyer, error) {
	model := new(Buyer)
	err := c.Put(fmt.Sprintf(buyerPathId, c.MarketPlaceId, id), params, nil, model)
	return model, err
}

func (c *Client) DelBuyer(id string) (*DeleteResponse, error) {
	del := new(DeleteResponse)
	err := c.Delete(fmt.Sprintf(buyerPathId, c.MarketPlaceId, id), nil, nil, del)
	return del, err
}

func (c *Client) ListBuyer() (*BuyersList, error) {
	list := new(BuyersList)
	err := c.Get(fmt.Sprintf(buyerPath, c.MarketPlaceId), nil, nil, list)
	return list, err
}

type BuyerParams struct {
	FirstName   string   `json:"first_name,omitempty"`
	LastName    string   `json:"last_name,omitempty"`
	Email       string   `json:"email,omitempty"`
	PhoneNumber string   `json:"phone_number,omitempty"`
	TaxpayerId  string   `json:"taxpayer_id,omitempty"`
	BirthDate   string   `json:"birthdate,omitempty"`
	Description string   `json:"description,omitempty"`
	Address     *Address `json:"address,omitempty"`
}

type Buyer struct {
	Id             string      `json:"id"`
	FirstName      string      `json:"first_name"`
	LastName       string      `json:"last_name"`
	Email          string      `json:"email"`
	PhoneNumber    string      `json:"phone_number"`
	TaxpayerId     string      `json:"taxpayer_id"`
	BirthDate      string      `json:"birthdate"`
	Description    string      `json:"description"`
	Status         string      `json:"status"`
	AccountBalance json.Number `json:"account_balance"`
	CurrentBalance json.Number `json:"current_balance"`
	Facebook       string      `json:"facebook"`
	Twitter        string      `json:"twitter"`
	Delinquent     bool        `json:"delinquent"`
	PaymentMethods interface{} `json:"payment_methods"` // TODO: set correct type
	DefaultDebit   string      `json:"default_debit"`
	DefaultCredit  string      `json:"default_credit"`
	DeliveryMethod string      `json:"default_receipt_delivery_method"`
	Metadata       interface{} `json:"metadata"`
	Address        *Address    `json:"address"`
	CreatedAt      time.Time   `json:"created_at"`
	UpdatedAt      time.Time   `json:"updated_at"`
	Info
}

type BuyersList struct {
	Buyers []Buyer `json:"items"`
	Info
	Pagination
}
