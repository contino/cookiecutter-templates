package main

import (
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestHandleAPIRequest(t *testing.T) {
		// given
		evt := events.APIGatewayProxyRequest{
			Body: "hello",
		}

		// when
		response, err := handleAPIRequest(nil, evt)

		// then
		assert.NoError(t, err)
		assert.Equal(t, "hello", response.Body)
}
