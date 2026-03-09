package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"strings"
)

type ContactRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8083"
	}

	mux := http.NewServeMux()

	// Serve React SPA static assets
	dist := http.Dir("./dist")
	fileServer := http.FileServer(dist)

	mux.HandleFunc("/contact", corsMiddleware(handleContact))

	// All other routes: serve from dist/, fall back to index.html for SPA routing
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		// Try to serve the file directly
		f, err := dist.Open(path)
		if err == nil {
			f.Close()
			fileServer.ServeHTTP(w, r)
			return
		}

		// Fall back to index.html for client-side routing
		http.ServeFile(w, r, "./dist/index.html")
	})

	fmt.Printf("Server starting on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		allowed := []string{
			"https://michael.brian.dominicatechnologies.com",
			"https://a6c27ba7-da19-4a1b-a30a-66b313a19446.lovableproject.com",
			"https://id-preview--a6c27ba7-da19-4a1b-a30a-66b313a19446.lovable.app",
		}
		for _, o := range allowed {
			if origin == o {
				w.Header().Set("Access-Control-Allow-Origin", origin)
				break
			}
		}
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next(w, r)
	}
}

func handleContact(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ContactRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		jsonError(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	req.Name = strings.TrimSpace(req.Name)
	req.Email = strings.TrimSpace(req.Email)
	req.Message = strings.TrimSpace(req.Message)

	if req.Name == "" || req.Email == "" || req.Message == "" {
		jsonError(w, "All fields are required", http.StatusBadRequest)
		return
	}
	if len(req.Name) > 100 || len(req.Email) > 255 || len(req.Message) > 1000 {
		jsonError(w, "Input exceeds maximum length", http.StatusBadRequest)
		return
	}

	if err := sendEmail(req); err != nil {
		log.Printf("Error sending email: %v", err)
		jsonError(w, "Failed to send message", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Message sent successfully"}) //nolint:errcheck
}

func sendEmail(req ContactRequest) error {
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASSWORD")
	to := os.Getenv("CONTACT_EMAIL")

	if smtpHost == "" {
		smtpHost = "smtp.gmail.com"
	}
	if smtpPort == "" {
		smtpPort = "587"
	}
	if to == "" {
		to = smtpUser
	}

	subject := "Portfolio contact from " + req.Name
	body := fmt.Sprintf("Name: %s\nEmail: %s\n\n%s", req.Name, req.Email, req.Message)
	msg := []byte("To: " + to + "\r\n" +
		"From: " + smtpUser + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" + body)

	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)
	return smtp.SendMail(smtpHost+":"+smtpPort, auth, smtpUser, []string{to}, msg)
}

func jsonError(w http.ResponseWriter, msg string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": msg}) //nolint:errcheck
}
