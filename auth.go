package zoop

import (
	"encoding/base64"
	"fmt"
)

type Client struct {
	ApiKey        string
	MarketplaceId string
	SellerId      string
	BasicAuth     string
}

func NewClient(key, mkt, seller string) *Client {

	return &Client{
		ApiKey:        key,
		MarketplaceId: mkt,
		SellerId:      seller,
		BasicAuth:     fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(key+":"))),
	}

}

func (c *Client) HealthCheck() error {
	return c.Get(fmt.Sprintf("%shealthcheck", baseUrl), nil, nil, nil)
}
