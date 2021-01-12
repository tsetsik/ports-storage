package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	target := ":9000"
	client, err := NewClient(target)

	assert.Nil(t, err)
	assert.NotNil(t, client)
}

func TestUpsertPort(t *testing.T) {
	target := ":9000"
	client, err := NewClient(target)
	assert.Nil(t, err)

	data := map[string]interface{}{"foo": "bar"}
	response, err := client.UpsertPort(12, data)

	assert.Nil(t, err)
	assert.Equal(t, int32(12), response.Id)
}
