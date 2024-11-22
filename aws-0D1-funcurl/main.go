package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/lambdaurl"
)

// HelloWorldHandler is a simple struct that conforms to the http.Handler interface.
type HelloWorldHandler struct{}

// ServeHTTP is a method that conforms to the http.Handler interface.
// It responds to HTTP requests with "Hello, World!".
func (h HelloWorldHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Hello", "world1")
	w.Header().Add("Hello", "world2")
	http.SetCookie(w, &http.Cookie{Name: "yummy", Value: "cookie"})
	http.SetCookie(w, &http.Cookie{Name: "yummy", Value: "cake"})
	http.SetCookie(w, &http.Cookie{Name: "fruit", Value: "banana", Expires: time.Date(2000, time.January, 0, 0, 0, 0, 0, time.UTC)})
	for _, c := range r.Cookies() {
		http.SetCookie(w, c)
	}

	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	_ = encoder.Encode(struct{ RequestQueryParams, Method any }{r.URL.Query(), r.Method})
	// w.Write([]byte("Hello World!"))
	// fmt.Fprintln(w, "Hello, Wonderful World!")
}

func main() {
	// Create an instance of HelloWorldHandler.
	handler := HelloWorldHandler{}

	// Create a new HTTP server and set the handler for the "/" route.
	http.Handle("/", handler)

	// Start the server on port 8080.
	log.Println("Server started at http://localhost:8080")

	lambdaurl.Start(handler)
}
