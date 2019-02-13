package zoop

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var client *Client

func init() {
	client = NewClient("", "", "")
}

func TestHealthCheck(t *testing.T) {
	err := client.HealthCheck()
	assert.NoError(t, err)
}
