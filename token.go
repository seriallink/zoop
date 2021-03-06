package zoop

import (
	"fmt"
	"time"
)

func (c *Client) NewCardToken(params *CreditCardParams) (*Token, error) {
	model := new(Token)
	err := c.Post(fmt.Sprintf("marketplaces/%s/cards/tokens", c.MarketplaceId), params, nil, model)
	return model, err
}

func (c *Client) GetToken(id string) (*Token, error) {
	model := new(Token)
	err := c.Get(fmt.Sprintf("marketplaces/%s/tokens/%s", c.MarketplaceId, id), nil, nil, model)
	return model, err
}

type Token struct {
	Id        string      `json:"id"`
	Used      bool        `json:"used"`
	Type      string      `json:"type"`
	Card      *CreditCard `json:"card"`
	Resource  string      `json:"resource"`
	Uri       string      `json:"uri"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}
