package zoop

import "time"

type CreditCardParams struct {
	HolderName      string `json:"holder_name"`
	ExpirationMonth string `json:"expiration_month"`
	ExpirationYear  string `json:"expiration_year"`
	CardNumber      string `json:"card_number"`
	SecurityCode    string `json:"security_code"`
}

type CreditCardModel struct {
	Id              string      `json:"id"`
	HolderName      string      `json:"holder_name"`
	ExpirationMonth string      `json:"expiration_month"`
	ExpirationYear  string      `json:"expiration_year"`
	First4Digits    string      `json:"first4_digits"`
	CardBrand       string      `json:"card_brand"`
	Description     string      `json:"description"`
	IsActive        bool        `json:"is_active"`
	IsValid         bool        `json:"is_valid"`
	IsVerified      bool        `json:"is_verified"`
	Customer        string      `json:"customer"`
	Fingerprint     string      `json:"fingerprint"`
	Address         *Address    `json:"address"`
	Checklist       *Checklist  `json:"verification_checklist"`
	Metadata        interface{} `json:"metadata"`
	CreatedAt       time.Time   `json:"created_at"`
	UpdatedAt       time.Time   `json:"updated_at"`
	Info
}

type Checklist struct {
	PostalCodeCheck   string `json:"postal_code_check"`
	SecurityCodeCheck string `json:"security_code_check"`
	AddressLine1Check string `json:"address_line1_check"`
}
