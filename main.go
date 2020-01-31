package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

func main() {
	mux := runtime.NewServeMux()

	// Apply the CORS middleware to our top-level router, with the defaults.
	corsMux := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
	)(mux)

	logMux := LoggingMiddleware(corsMux)

	http.ListenAndServe(":9999", logMux)
}

func userEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("### Request ###\nMethod:%v\nHeader:%v\nBody:%v", r.Method, r.Header, r.Body)
		// corsAllowOriginHeader := "Access-Control-Allow-Origin"
		// w.Header().Set(corsAllowOriginHeader, "http://localhost:3000")
		next.ServeHTTP(w, r)
	})
}
