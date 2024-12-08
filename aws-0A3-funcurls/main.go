package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

// HandleRequest processes Lambda Function URL requests
func HandleRequest(ctx context.Context, request events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	// https://pkg.go.dev/github.com/aws/aws-lambda-go@v1.47.0/events#LambdaFunctionURLRequest

	// Mux
	switch request.RequestContext.HTTP.Path {
	case "/", "/hello":
		return helloHandler(request)
	case "/echo":
		return echoHandler(request)
	default:
		return notFoundHandler(request)
	}
}

// helloHandler handles Hello requests
func helloHandler(request events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	// Return a hello message
	response := events.LambdaFunctionURLResponse{
		StatusCode:      200,
		Body:            "Hello from Lambda Function URL!",
		IsBase64Encoded: false,
	}
	return response, nil
}

// echoHandler handles Echo requests
func echoHandler(request events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	// Echo back the received req RawQueryString
	response := events.LambdaFunctionURLResponse{
		StatusCode:      200,
		Body:            "Echo: " + request.RawQueryString,
		IsBase64Encoded: false,
	}
	return response, nil
}

// notFoundHandler handles requests for paths that are not found
func notFoundHandler(request events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	// Return a 404 not found response
	response := events.LambdaFunctionURLResponse{
		StatusCode:      404,
		Body:            "Not Found",
		IsBase64Encoded: false,
	}
	return response, nil
}

func main() {
	// Start the Lambda function
	lambda.Start(HandleRequest)
}
