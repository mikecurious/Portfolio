package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/smtp"
	"os"
)

type ContactForm struct {
	Name    string
	Email   string
	Message string
}

type PageData struct {
	Title    string
	Name     string
	Bio      string
	Skills   []string
	Projects []Project
}

type Project struct {
	Title       string
	Description string
	ImagePath   string
}

func main() {
	// Configure server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Define file server for static files
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Define routes
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/contact", handleContact)

	// Start server
	fmt.Printf("Server starting on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Sample data - replace with your information
	data := PageData{
		Title: "My Portfolio",
		Name:  "Your Name",
		Bio:   "Web developer and software engineer with expertise in Go, JavaScript, and cloud technologies.",
		Skills: []string{
			"Go", "JavaScript", "HTML/CSS", "Docker", "Cloud Platforms",
		},
		Projects: []Project{
			{
				Title:       "Project 1",
				Description: "Description of project 1",
				ImagePath:   "/static/images/project1.jpg",
			},
			{
				Title:       "Project 2",
				Description: "Description of project 2",
				ImagePath:   "/static/images/project2.jpg",
			},
			{
				Title:       "Project 3",
				Description: "Description of project 3",
				ImagePath:   "/static/images/project3.jpg",
			},
		},
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleContact(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Could not parse form", http.StatusBadRequest)
		return
	}

	contactForm := ContactForm{
		Name:    r.FormValue("name"),
		Email:   r.FormValue("email"),
		Message: r.FormValue("message"),
	}

	// Validate form (basic validation)
	if contactForm.Name == "" || contactForm.Email == "" || contactForm.Message == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	// Send email (you'll need to configure your SMTP settings)
	err = sendEmail(contactForm)
	if err != nil {
		log.Printf("Error sending email: %v", err)
		http.Error(w, "Failed to send message", http.StatusInternalServerError)
		return
	}

	// Redirect back to home page with success message
	http.Redirect(w, r, "/?message=success", http.StatusSeeOther)
}

func sendEmail(form ContactForm) error {
	// Configure these with your actual email settings
	from := "your-email@example.com"
	password := "your-password" // Consider using environment variables for this
	to := "your-email@example.com"
	smtpHost := "smtp.example.com"
	smtpPort := "587"

	// Compose message
	subject := "New contact form submission from " + form.Name
	body := fmt.Sprintf("Name: %s\nEmail: %s\nMessage: %s", form.Name, form.Email, form.Message)
	message := []byte("To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body)

	// Authentication
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Send mail
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, message)
	return err
}
