package zoop

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	sellerPath   = "marketplaces/%s/sellers"
	sellerPathId = "marketplaces/%s/sellers/%s"
)

// Get seller's detail
func (c *Client) GetSeller() (*Seller, error) {
	model := new(Seller)
	err := c.Get(fmt.Sprintf(sellerPathId, c.MarketplaceId, c.SellerId), nil, nil, model)
	return model, err
}

type Seller struct {
	Id                        string         `json:"id"`
	Status                    string         `json:"status"`
	Resource                  string         `json:"resource"`
	Type                      string         `json:"type"`
	Description               string         `json:"description"`
	AccountBalance            json.Number    `json:"account_balance"`
	CurrentBalance            json.Number    `json:"current_balance"`
	BusinessName              string         `json:"business_name"`
	BusinessPhone             string         `json:"business_phone"`
	BusinessEmail             string         `json:"business_email"`
	BusinessWebsite           string         `json:"business_website"`
	BusinessDescription       string         `json:"business_description"`
	BusinessOpeningDate       Date           `json:"business_opening_date"`
	BusinessFacebook          string         `json:"business_facebook"`
	BusinessTwitter           string         `json:"business_twitter"`
	Ein                       string         `json:"ein"`
	StatementDescriptor       string         `json:"statement_descriptor"`
	MCC                       string         `json:"mcc"`
	ShowProfileOnline         bool           `json:"show_profile_online"`
	IsMobile                  bool           `json:"is_mobile"`
	DeclineOnFailSecurityCode bool           `json:"decline_on_fail_security_code"`
	DeclineOnFailZipCode      bool           `json:"decline_on_fail_zipcode"`
	Delinquent                bool           `json:"delinquent"`
	DefaultDebit              string         `json:"default_debit"`
	DefaultCredit             string         `json:"default_credit"`
	MerchantCode              string         `json:"merchant_code"`
	TerminalCode              string         `json:"terminal_code"`
	URI                       string         `json:"uri"`
	MarketplaceId             string         `json:"marketplace_id"`
	Metadata                  interface{}    `json:"metadata"`
	BusinessAddress           *Address       `json:"business_address"`
	Owner                     *Owner         `json:"owner"`
	PaymentMethods            *PaymentMethod `json:"payment_methods"`
	CreatedAt                 time.Time      `json:"created_at"`
	UpdatedAt                 time.Time      `json:"updated_at"`
}
