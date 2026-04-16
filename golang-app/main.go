package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var s3Client *s3.Client

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Printf("AWS config not available: %v", err)
	} else {
		s3Client = s3.NewFromConfig(cfg)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handleRoot)
	mux.HandleFunc("/health", handleHealth)
	mux.HandleFunc("/aws", handleAWS)

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal(err)
	}
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Hello from Go demo app!",
		"status":  "ok",
	})
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "healthy"})
}

func handleAWS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"sdk":     "aws-sdk-go-v2",
		"service": "s3",
	})
}
