package zoop

import (
	"fmt"
	"time"
)

func (c *Client) NewCardToken(params *CreditCardParams) (*Token, error) {
	model := new(Token)
	err := c.Post(fmt.Sprintf("marketplaces/%s/cards/tokens", c.MarketPlaceId), params, nil, model)
	return model, err
}

func (c *Client) GetToken(id string) (*Token, error) {
	model := new(Token)
	err := c.Get(fmt.Sprintf("marketplaces/%s/tokens/%s", c.MarketPlaceId, id), nil, nil, model)
	return model, err
}

type Token struct {
	Id        string      `json:"id"`
	Used      bool        `json:"used"`
	Type      string      `json:"type"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Card      *CreditCard `json:"card"`
	Info
}
