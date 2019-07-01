package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// handleAPIRequest returns an APIGatewayProxyResponse with the value of APIGatewayProxyRequest
func handleAPIRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		Body:       request.Body,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(handleAPIRequest)
}
