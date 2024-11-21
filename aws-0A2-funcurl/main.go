package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// Handler processes the Lambda Function URL requests
func Handler(ctx context.Context, req events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	// Log HTTP method
	log.Printf("HTTP Method: %s", req.RequestContext.HTTP.Method)

	// Log request headers
	log.Printf("Headers: %+v", req.Headers)

	// Log request body
	log.Printf("Body: %s", req.Body)

	// Build response
	responseBody, err := json.Marshal(map[string]interface{}{
		"message": "Echo from Lambda Function URL",
		"method":  req.RequestContext.HTTP.Method,
		"headers": req.Headers,
		"body":    req.Body,
	})
	// Not working, everything reported empty:
	// "body": "{\"body\":\"\",\"headers\":null,\"message\":\"Echo from Lambda Function URL\",\"method\":\"\"}",
	if err != nil {
		log.Printf("Error marshalling response: %v", err)
		return events.LambdaFunctionURLResponse{
			StatusCode: 500,
			Body:       "Internal Server Error",
		}, nil
	}

	return events.LambdaFunctionURLResponse{
		StatusCode: 200,
		Body:       string(responseBody),
	}, nil
}

// StartLambda initializes the Lambda function
func StartLambda() {
	lambda.Start(Handler)
}

func main() {
	// Configure logging to stdout for CloudWatch
	log.SetOutput(os.Stdout)

	// Start the Lambda function
	StartLambda()
}
