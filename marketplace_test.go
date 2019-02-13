package zoop

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetMarketplace(t *testing.T) {
	marketplace, err := client.GetMarketplace()
	assert.NoError(t, err)
	assert.Equal(t, client.MarketplaceId, marketplace.Id)
}
