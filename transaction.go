package zoop

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	transactionPath    = "marketplaces/%s/transactions"
	transactionPathId  = "marketplaces/%s/transactions/%s"
	transactionSeller  = "marketplaces/%s/sellers/%s/transactions"
	transactionCapture = "marketplaces/%s/transactions/%s/capture"
	transactionRefund  = "marketplaces/%s/transactions/%s/void"
)

// Enums
type PaymentType string
type UsageType string
type SourceType string
type InstallmentPlanType string
type TransactionStatus string

const (
	// PaymentType
	TypeCredit PaymentType = "credit"
	TypeDebit  PaymentType = "debit"
	TypeWallet PaymentType = "wallet"
	TypeBoleto PaymentType = "boleto"

	// UsageType
	SingleUse UsageType = "single_use"
	Reusable  UsageType = "reusable"

	// SourceType
	SourceWallet        SourceType = "wallet"
	SourceCard          SourceType = "card"
	SourceCardAndWallet SourceType = "card_and_wallet"
	SourceToken         SourceType = "token"
	SourceCustomer      SourceType = "customer"
	SourceThreeDSecure  SourceType = "three_d_secure"
	SourceDebitOnline   SourceType = "debit_online"

	// InstallmentPlanType
	InterestFree InstallmentPlanType = "interest_free"
	WithInterest InstallmentPlanType = "with_interest"

	// TransactionStatus
	StatusNew           = "new"
	StatusPending       = "pending"
	StatusPreAuthorized = "pre_authorized"
	StatusSucceeded     = "succeeded"
	StatusFailed        = "failed"
	StatusReversed      = "reversed"
	StatusCanceled      = "canceled"
	StatusRefunded      = "refunded"
	StatusDispute       = "dispute"
	StatusChargedBack   = "charged_back"
)

func (c *Client) GetTransaction(id string) (*Transaction, error) {
	model := new(Transaction)
	err := c.Get(fmt.Sprintf(transactionPathId, c.MarketplaceId, id), nil, nil, model)
	return model, err
}

func (c *Client) NewTransaction(params *TransactionParams) (*Transaction, error) {
	model := new(Transaction)
	err := c.Post(fmt.Sprintf(transactionPath, c.MarketplaceId), params, nil, model)
	return model, err
}

func (c *Client) CaptureTransaction(id string, amount int) (*Transaction, error) {
	model := new(Transaction)
	params := &Params{"on_behalf_of": c.SellerId, "amount": amount}
	err := c.Post(fmt.Sprintf(transactionCapture, c.MarketplaceId, id), params, nil, model)
	return model, err
}

func (c *Client) RefundTransaction(id string, amount int) (*Transaction, error) {
	model := new(Transaction)
	params := &Params{"on_behalf_of": c.SellerId, "amount": amount}
	err := c.Post(fmt.Sprintf(transactionRefund, c.MarketplaceId, id), params, nil, model)
	return model, err
}

func (c *Client) ListMarketplaceTransactions() (*TransactionsList, error) {
	list := new(TransactionsList)
	err := c.Get(fmt.Sprintf(transactionPath, c.MarketplaceId), nil, nil, list)
	return list, err
}

func (c *Client) ListSellerTransactions() (*TransactionsList, error) {
	list := new(TransactionsList)
	err := c.Get(fmt.Sprintf(transactionSeller, c.MarketplaceId, c.SellerId), nil, nil, list)
	return list, err
}

type TransactionParams struct {
	Amount                     int                  `json:"amount"`
	Currency                   string               `json:"currency"`
	Description                string               `json:"description"`
	PaymentType                PaymentType          `json:"payment_type"`
	Capture                    bool                 `json:"capture"`
	OnBehalfOf                 string               `json:"on_behalf_of"` // seller_id
	ReferenceId                string               `json:"reference_id,omitempty"`
	StatementDescriptor        string               `json:"statement_descriptor,omitempty"`
	Customer                   string               `json:"customer,omitempty"`
	Token                      string               `json:"token,omitempty"`
	Metadata                   interface{}          `json:"metadata,omitempty"`
	SourceUsage                UsageType            `json:"source.usage,omitempty"`
	SourceAmount               int                  `json:"source.amount,omitempty"`
	SourceCurrency             string               `json:"source.currency,omitempty"`
	SourceDescription          string               `json:"source.description,omitempty"`
	SourceType                 SourceType           `json:"source.type"`
	SourceCapture              bool                 `json:"source.capture,omitempty"`
	SourceOnBehalfOf           string               `json:"source.on_behalf_of,omitempty"`
	SourceReferenceId          string               `json:"source.reference_id,omitempty"`
	SourceCardId               string               `json:"source.card.id,omitempty"`
	SourceCardHolderName       string               `json:"source.card.holder_name,omitempty"`
	SourceCardExpirationMonth  string               `json:"source.card.expiration_month,omitempty"`
	SourceCardExpirationYear   string               `json:"source.card.expiration_year,omitempty"`
	SourceCardCardNumber       string               `json:"source.card.card_number,omitempty"`
	SourceCardSecurityCode     string               `json:"source.card.security_code,omitempty"`
	SourceCardAmount           int                  `json:"source.card.amount,omitempty"`
	SourceInstallmentPlanMode  InstallmentPlanType  `json:"source.installment_plan.mode,omitempty"`
	SourceNumberOfInstallments int                  `json:"source.installment_plan.number_installments,omitempty"`
	SourceStatementDescriptor  string               `json:"source.statement_descriptor,omitempty"`
	SourceCustomerId           string               `json:"source.customer.id,omitempty"`
	SourceCustomerAmount       int                  `json:"source.customer.amount,omitempty"`
	SourceTokenId              string               `json:"source.token.id,omitempty"`
	SourceTokenAmount          int                  `json:"source.token.amount,omitempty"`
	SourceMetadata             interface{}          `json:"source.metadata,omitempty"`
	PaymentMethod              *PaymentMethodParams `json:"payment_method,omitempty"`
}

type PaymentMethodParams struct {
	ExpirationDate  time.Time `json:"expiration_date,omitempty"`
	TopInstructions []string  `json:"top_instructions,omitempty"`
}

type Transaction struct {
	Id                   string                `json:"id"`
	Status               string                `json:"status"`
	Amount               json.Number           `json:"amount"`
	OriginalAmount       json.Number           `json:"original_amount"`
	Currency             string                `json:"currency"`
	Description          string                `json:"description"`
	PaymentType          PaymentType           `json:"payment_type"`
	TransactionNumber    string                `json:"transaction_number"`
	GatewayAuthorizer    string                `json:"gateway_authorizer"`
	AppTransactionUid    string                `json:"app_transaction_uid"`
	Refunds              json.Number           `json:"refunds"`
	Rewards              json.Number           `json:"rewards"`
	Discounts            json.Number           `json:"discounts"`
	PreAuthorization     string                `json:"pre_authorization"`
	SalesReceipt         string                `json:"sales_receipt"`
	OnBehalfOf           string                `json:"on_behalf_of"`
	Customer             string                `json:"customer"`
	StatementDescriptor  string                `json:"statement_descriptor"`
	PaymentMethod        *PaymentMethod        `json:"payment_method"`
	Source               interface{}           `json:"source"` // TODO: set correct datatype
	PointOfSale          *PointOfSale          `json:"point_of_sale"`
	InstallmentPlan      interface{}           `json:"installment_plan"` // TODO: set correct datatype
	Refunded             bool                  `json:"refunded"`
	Voided               bool                  `json:"voided"`
	Captured             bool                  `json:"captured"`
	Fees                 json.Number           `json:"fees"`
	FeeDetails           []FeeDetail           `json:"fee_details"`
	LocationLatitude     string                `json:"location_latitude"`
	LocationLongitude    string                `json:"location_longitude"`
	PaymentAuthorization *PaymentAuthorization `json:"payment_authorization"`
	History              []History             `json:"history"`
	Resource             string                `json:"resource"`
	Uri                  string                `json:"uri"`
	Metadata             interface{}           `json:"metadata"`
	ExpectedOn           time.Time             `json:"expected_on"`
	CreatedAt            time.Time             `json:"created_at"`
	UpdatedAt            time.Time             `json:"updated_at"`
}

type PaymentMethod struct {
	Id               string      `json:"id"`
	Status           string      `json:"status"`
	Resource         string      `json:"resource"`
	Description      string      `json:"description"`
	CardBrand        string      `json:"card_brand"`
	First4Digits     string      `json:"first4_digits"`
	ExpirationMonth  string      `json:"expiration_month"`
	ExpirationYear   string      `json:"expiration_year"`
	HolderName       string      `json:"holder_name"`
	IsActive         bool        `json:"is_active"`
	IsValid          bool        `json:"is_valid"`
	IsVerified       bool        `json:"is_verified"`
	Customer         string      `json:"customer"`
	Fingerprint      string      `json:"fingerprint"`
	ReferenceNumber  string      `json:"reference_number"`
	DocumentNumber   string      `json:"document_number"`
	Recipient        string      `json:"recipient"`
	BankCode         string      `json:"bank_code"`
	Sequence         string      `json:"sequence"`
	Accepted         bool        `json:"accepted"`
	Printed          bool        `json:"printed"`
	Downloaded       bool        `json:"downloaded"`
	PaidAt           string      `json:"paid_at"`
	Barcode          string      `json:"barcode"`
	Url              string      `json:"url"`
	Uri              string      `json:"uri"`
	Metadata         interface{} `json:"metadata"`
	PaymentLimitDate time.Time   `json:"payment_limit_date"`
	ExpirationDate   time.Time   `json:"expiration_date"`
	CreatedAt        time.Time   `json:"created_at"`
	UpdatedAt        time.Time   `json:"updated_at"`
	Address          *Address    `json:"address"`
	Checklist        *Checklist  `json:"verification_checklist"`
}

type PointOfSale struct {
	EntryMode            string `json:"entry_mode"`
	IdentificationNumber string `json:"identification_number"`
}

type FeeDetail struct {
	Amount       json.Number `json:"amount"`
	Prepaid      bool        `json:"prepaid"`
	Currency     string      `json:"currency"`
	Type         string      `json:"type"`
	IsGatewayFee bool        `json:"is_gateway_fee"`
	Description  string      `json:"description"`
}

type PaymentAuthorization struct {
	AuthorizerId      string `json:"authorizer_id"`
	AuthorizationCode string `json:"authorization_code"`
	AuthorizationNsu  string `json:"authorization_nsu"`
}

type History struct {
	Id                string      `json:"id"`
	Transaction       string      `json:"transaction"`
	Amount            json.Number `json:"amount"`
	OperationType     string      `json:"operation_type"`
	Status            string      `json:"status"`
	ResponseCode      string      `json:"response_code"`
	ResponseMessage   string      `json:"response_message"`
	AuthorizationCode string      `json:"authorization_code"`
	AuthorizerId      string      `json:"authorizer_id"`
	AuthorizationNsu  string      `json:"authorization_nsu"`
	CreatedAt         string      `json:"created_at"`
}

type TransactionsList struct {
	Transactions []Transaction `json:"items"`
	Resource     string        `json:"resource"`
	Uri          string        `json:"uri"`
	Pagination
}
