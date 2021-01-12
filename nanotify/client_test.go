package nanotify_test

import (
	"fmt"
	"testing"

	"github.com/nanotify/nanotify-sdk-go/nanotify"
	"github.com/stretchr/testify/assert"
)

func TestNewClient(t *testing.T) {
	asserts := assert.New(t)

	apiKey := "test-api-key"
	client, err := nanotify.NewClient(apiKey, nanotify.NetworkNano)

	asserts.NoError(err)

	asserts.Equal(apiKey, client.APIKey)
	asserts.Equal(nanotify.NetworkNano, client.Network)
}

func TestErrorNewClient(t *testing.T) {
	asserts := assert.New(t)

	testCases := []struct {
		APIKey        string
		Network       nanotify.Network
		ExpectedError error
	}{
		{
			APIKey:        "",
			ExpectedError: fmt.Errorf("invalid api key"),
		},
		{
			APIKey:        "api-key",
			Network:       "foobar",
			ExpectedError: fmt.Errorf("invalid network"),
		},
		{
			APIKey:  "api-key",
			Network: "nano",
		},
		{
			APIKey:  "api-key",
			Network: "nano-beta",
		},
	}

	for _, test := range testCases {
		_, err := nanotify.NewClient(test.APIKey, test.Network)

		if test.ExpectedError != nil {
			asserts.EqualError(err, test.ExpectedError.Error())
		} else {
			asserts.NoError(err)
		}
	}
}
