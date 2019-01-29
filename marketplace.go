package zoop

import (
	"encoding/json"
	"fmt"
	"time"
)

type Marketplace struct {
	Id                        string      `json:"id"`
	Name                      string      `json:"name"`
	Type                      string      `json:"type"`
	Description               string      `json:"description"`
	IsActive                  bool        `json:"is_active"`
	IsVerified                bool        `json:"is_verified"`
	AccountBalance            json.Number `json:"account_balance"`
	CurrentBalance            json.Number `json:"current_balance"`
	DeclineOnFailSecurityCode bool        `json:"decline_on_fail_security_code"`
	DeclineOnFailZipCode      bool        `json:"decline_on_fail_zip_code"`
	SupportEmail              string      `json:"support_email"`
	PhoneNumber               string      `json:"phone_number"`
	StatementDescriptor       string      `json:"statement_descriptor"`
	Website                   string      `json:"website"`
	Facebook                  string      `json:"facebook"`
	Twitter                   string      `json:"twitter"`
	TransferEnabled           bool        `json:"transfer_enabled"`
	TransferPolicy            string      `json:"transfer_policy"`
	DebitEnabled              bool        `json:"debit_enabled"`
	DefaultDebit              string      `json:"default_debit"`
	DefaultCredit             string      `json:"default_credit"`
	Customer                  *Customer   `json:"customer"`
	ApiKeys                   []Key       `json:"api_keys"`
	Resource                  string      `json:"resource"`
	Uri                       string      `json:"uri"`
	CreatedAt                 time.Time   `json:"created_at"`
	UpdatedAt                 time.Time   `json:"updated_at"`
}

func (c *Client) GetMarketplace() (*Marketplace, error) {
	model := new(Marketplace)
	err := c.Get(fmt.Sprintf("marketplaces/%s", c.MarketplaceId), nil, nil, model)
	return model, err
}
