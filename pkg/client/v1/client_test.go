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

	clientPort := &Port{
		ID:      123,
		Name:    "Ajman",
		City:    "Ajman",
		Country: "United Arab Emirates",
		Alias:   nil,
		Regions: nil,
		Coordinates: []float32{
			float32(55.5136433),
			float32(25.4052165),
		},
		Province: "Ajman",
		Timezone: "Asia/Dubai",
		Unlocs:   []string{"AEAJM"},
		Code:     "52000",
	}
	response, err := client.UpsertPort(clientPort)

	assert.Nil(t, err)
	assert.Equal(t, clientPort, response)
}
