package main

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	// Standard HTTP handler function
	httpHandler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello, World!")
	}

	// Wrap the HTTP handler to work with Lambda Function URL
	lambdaHandler := WrapHandler(httpHandler)

	// Start the Lambda handler
	lambda.Start(lambdaHandler)
}
