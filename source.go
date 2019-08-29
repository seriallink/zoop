package zoop

import "fmt"

const (
	sourcePath   = "marketplaces/%s/sources"
	sourcePathId = "marketplaces/%s/sources/%s"
)

// Get sources's detail
func (c *Client) GetSource(id string) (*Source, error) {
	model := new(Source)
	err := c.Get(fmt.Sprintf(sourcePathId, c.MarketplaceId, id), nil, nil, model)
	return model, err
}

// Create a new source
func (c *Client) NewSource(params *SourceParams) (*Source, error) {
	model := new(Source)
	err := c.Post(fmt.Sprintf(sourcePath, c.MarketplaceId), params, nil, model)
	return model, err
}

// Delete source by ID
func (c *Client) DelSource(id string) (*DeleteResponse, error) {
	del := new(DeleteResponse)
	err := c.Delete(fmt.Sprintf(sourcePathId, c.MarketplaceId, id), nil, nil, del)
	return del, err
}

type Source struct {
}

type SourceParams struct {
	Usage               UsageType              `json:"usage,omitempty"`
	Amount              int                    `json:"amount,omitempty"`
	Currency            string                 `json:"currency,omitempty"`
	Description         string                 `json:"description,omitempty"`
	Type                SourceType             `json:"type,omitempty"`
	Capture             bool                   `json:"capture,omitempty"`
	OnBehalfOf          string                 `json:"on_behalf_of,omitempty"`
	ReferenceId         string                 `json:"reference_id,omitempty"`
	StatementDescriptor string                 `json:"statement_descriptor,omitempty"`
	Card                *SourceCardParams      `json:"card,omitempty"`
	Customer            *SourceCustomerParams  `json:"customer,omitempty"`
	Token               *SourceTokenParams     `json:"token,omitempty"`
	InstallmentPlan     *InstallmentPlanParams `json:"installment_plan,omitempty"`
	Metadata            interface{}            `json:"metadata,omitempty"`
}

type SourceCardParams struct {
	CardId              string `json:"card.id,omitempty"`
	CardHolderName      string `json:"card.holder_name,omitempty"`
	CardExpirationMonth string `json:"card.expiration_month,omitempty"`
	CardExpirationYear  string `json:"card.expiration_year,omitempty"`
	CardCardNumber      string `json:"card.card_number,omitempty"`
	CardSecurityCode    string `json:"card.security_code,omitempty"`
	CardAmount          int    `json:"card.amount,omitempty"`
}

type InstallmentPlanParams struct {
	Mode               InstallmentPlanType `json:"mode,omitempty"`
	NumberInstallments int                 `json:"number_installments,omitempty"`
}

type SourceCustomerParams struct {
	CustomerId     string `json:"id,omitempty"`
	CustomerAmount int    `json:"amount,omitempty"`
}

type SourceTokenParams struct {
	TokenId     string `json:"id,omitempty"`
	TokenAmount int    `json:"amount,omitempty"`
}
