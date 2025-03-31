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
	Title           string
	Name            string
	Role            string
	Bio             string
	Skills          []Skill
	WorkExperiences []WorkExperience
	Education       []Education
	PersonalInfo    PersonalInfo
	Projects        []Project
}

type Skill struct {
	Category string
	Items    []string
}

type WorkExperience struct {
	Company     string
	Position    string
	Duration    string
	Description []string
}

type Education struct {
	Institution string
	Degree      string
	Duration    string
}

type PersonalInfo struct {
	Phone         string
	Email         string
	Languages     string
	DOB           string
	LinkedIn      string
	Github        string
	MSLearn       string
	GCPProfile    string
	MaritalStatus string
}

type Project struct {
	Title       string
	Description string
	ImagePath   string
	TechStack   []string
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

	// Data from Michael Brian Muthee's CV
	data := PageData{
		Title: "Michael Brian Muthee - DevOps/System Admin",
		Name:  "Michael Brian Muthee",
		Role:  "DevOps/System Admin",
		Bio:   "To leverage my technical expertise in backend development, systems administration, and network engineering to create innovative, efficient, and scalable solutions. I seek a position that will challenge my skills and offer professional growth while contributing to the success and goals of the organization.",
		Skills: []Skill{
			{
				Category: "Programming Languages and Frameworks",
				Items:    []string{"Go (Golang)", "Python (Fast API)", "PHP", "Linux/UNIX Systems", "Bash Scripting"},
			},
			{
				Category: "Cloud and DevOps",
				Items:    []string{"AWS", "Azure", "Google Cloud Platform (GCP)", "Server Pronto", "CI/CD Pipelines", "Infrastructure as Code", "Terraform", "Ansible", "Containerization", "Docker", "Kubernetes"},
			},
			{
				Category: "Network Engineering and Cybersecurity",
				Items:    []string{"Network Segmentation", "Firewall Configuration", "VPN (Cisco ASA)", "SIEM", "Penetration Testing", "ISO 27001", "PCI DSS"},
			},
			{
				Category: "API Development and Integration",
				Items:    []string{"RESTful APIs", "Mobile Money Platforms (M-Pesa, MoMo, Airtel Money)", "Payment Gateways", "Financial Integration", "Token-based Authentication"},
			},
		},
		WorkExperiences: []WorkExperience{
			{
				Company:  "Yet Kenya Limited",
				Position: "DevOps/System Admin",
				Duration: "2022 - Present",
				Description: []string{
					"Develop and maintain Infrastructure as Code (IaC) scripts using Terraform and Ansible",
					"Perform system audits on external companies to ensure compliance",
					"Design and implement CI/CD pipelines to automate build, test, and deployment processes",
					"Implement monitoring systems using Prometheus and Grafana",
					"Manage cloud infrastructure on AWS, Azure, and Google Cloud Platform",
					"Develop and optimize Docker containers with Kubernetes orchestration",
					"Implement security best practices for infrastructure and networks",
					"Establish backup and disaster recovery policies",
					"Conduct regular security audits and penetration tests",
				},
			},
			{
				Company:  "ANIMAL HEALTH AND INDUSTRY TRAINING INSTITUTE (AHITI KABETE)",
				Position: "IT Support",
				Duration: "2020 - 2022",
				Description: []string{
					"Provided technical support and IT services",
					"Maintained computer systems and network infrastructure",
				},
			},
			{
				Company:  "Comet Designers",
				Position: "Technical Support/Office Assistant",
				Duration: "2018 - 2020",
				Description: []string{
					"Data Entry",
					"Computer Troubleshooting",
					"KVB Indexing",
					"Customer Service",
					"Printing and Design",
					"Record Keeping",
					"Office Messenger",
				},
			},
		},
		Education: []Education{
			{
				Institution: "Zetech University",
				Degree:      "Diploma in Information Technology",
				Duration:    "September 2019 - Present",
			},
			{
				Institution: "Githiga Boys' High School",
				Degree:      "Secondary Education",
				Duration:    "January 2014 - November 2017",
			},
			{
				Institution: "Westlands Primary School",
				Degree:      "Primary Education",
				Duration:    "January 2005 - October 2013",
			},
		},
		PersonalInfo: PersonalInfo{
			Phone:         "+254-758-930-908",
			Email:         "mikkohbrayoh@gmail.com",
			Languages:     "English, Swahili",
			DOB:           "21st July 1998",
			MaritalStatus: "Married",
			LinkedIn:      "www.linkedin.com/in/micheal-brian-456041215",
			Github:        "https://github.com/mikecurious",
			MSLearn:       "MichaelBrian-3822 | Microsoft Learn",
			GCPProfile:    "https://www.cloudskillsboost.google/public_profiles/f9cccded867a41758c75790939f8137e",
		},
		Projects: []Project{
			{
				Title:       "Cloud Migration Project",
				Description: "Migrated on-premises infrastructure to AWS cloud, implementing IaC with Terraform and setting up CI/CD pipelines.",
				ImagePath:   "/static/images/project1.jpg",
				TechStack:   []string{"AWS", "Terraform", "Jenkins", "Docker"},
			},
			{
				Title:       "Payment Gateway Integration",
				Description: "Developed API integrations for various payment systems including M-Pesa and credit card processors using Go and FastAPI.",
				ImagePath:   "/static/images/project2.jpg",
				TechStack:   []string{"Go", "Python", "FastAPI", "RESTful APIs"},
			},
			{
				Title:       "Network Security Implementation",
				Description: "Configured secure network architecture with VPNs, firewalls, and implemented ISO 27001 compliance measures.",
				ImagePath:   "/static/images/project3.jpg",
				TechStack:   []string{"Cisco ASA", "VPN", "Firewall", "ISO 27001"},
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
	from := "mikkohbrayoh@gmail.com"
	password := "your-password" // Consider using environment variables for this
	to := "mikkohbrayoh@gmail.com"
	smtpHost := "smtp.gmail.com"
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
