package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/suntong/lambdaurl"
)

func main() {
	// Create a new ServeMux router
	mux := http.NewServeMux()

	// Define routes and handlers
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Hello, World! (GET)")
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method not allowed")
		}
	})

	mux.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			// Read the request body
			var requestBody map[string]string
			err := json.NewDecoder(r.Body).Decode(&requestBody)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "Invalid request body")
				return
			}

			// Respond with a personalized greeting
			name, ok := requestBody["name"]
			if !ok {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, "Name is required")
				return
			}

			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Hello, %s! (POST)", name)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method not allowed")
		}
	})

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Service is healthy!")
	})

	// Wrap the ServeMux with the Lambda handler
	lambdaHandler := lambdaurl.WrapHandler(mux)

	// Start the Lambda handler
	lambda.Start(lambdaHandler)
}
