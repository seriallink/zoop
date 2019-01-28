package zoop

import (
	"encoding/json"
	"time"
)

type Customer struct {
	Id                        string      `json:"id"`
	FirstName                 string      `json:"first_name"`
	LastName                  string      `json:"last_name"`
	Email                     string      `json:"email"`
	PhoneNumber               string      `json:"phone_number"`
	TaxpayerId                string      `json:"taxpayer_id"`
	BirthDate                 string      `json:"birthdate"`
	Description               string      `json:"description"`
	Status                    string      `json:"status"`
	AccountBalance            json.Number `json:"account_balance"`
	CurrentBalance            json.Number `json:"current_balance"`
	BusinessName              string      `json:"business_name"`
	BusinessPhone             string      `json:"business_phone"`
	BusinessEmail             string      `json:"business_email"`
	BusinessWebsite           string      `json:"business_website"`
	BusinessFacebook          string      `json:"business_facebook"`
	BusinessTwitter           string      `json:"business_twitter"`
	BusinessDescription       string      `json:"business_description"`
	BusinessOpeningDate       string      `json:"business_opening_date"`
	Ein                       string      `json:"ein"`
	StatementDescriptor       string      `json:"statement_descriptor"`
	MCC                       string      `json:"mcc"`
	ShowProfileOnline         bool        `json:"show_profile_online"`
	IsMobile                  bool        `json:"is_mobile"`
	DeclineOnFailSecurityCode bool        `json:"decline_on_fail_security_code"`
	DeclineOnFailZipCode      bool        `json:"decline_on_fail_zip_code"`
	Delinquent                bool        `json:"delinquent"`
	PaymentMethods            interface{} `json:"payment_methods"` // TODO: couldn't find the correct datatype
	DefaultDebit              string      `json:"default_debit"`
	DefaultCredit             string      `json:"default_credit"`
	MerchantCode              string      `json:"merchant_code"`
	TerminalCode              string      `json:"terminal_code"`
	MarketplaceId             string      `json:"marketplace_id"`
	BusinessAddress           *Address    `json:"business_address"`
	Metadata                  interface{} `json:"metadata"`
	CreatedAt                 time.Time   `json:"created_at"`
	UpdatedAt                 time.Time   `json:"updated_at"`
	Info
}
