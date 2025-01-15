package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

// AgeResult represents the calculated age
type AgeResult struct {
	Years  int
	Months int
	Days   int
	Error  string
}

// calculateAge computes the precise age from a given birthdate
func calculateAge(birthdate time.Time) AgeResult {
	now := time.Now()
	
	// Calculate years, months, and days
	years := now.Year() - birthdate.Year()
	months := int(now.Month() - birthdate.Month())
	days := now.Day() - birthdate.Day()

	// Adjust for negative months or days
	if days < 0 {
		months--
		lastMonth := now.AddDate(0, -1, 0)
		days += time.Date(lastMonth.Year(), lastMonth.Month()+1, 0, 0, 0, 0, 0, time.UTC).Day()
	}

	if months < 0 {
		years--
		months += 12
	}

	return AgeResult{
		Years:  years,
		Months: months,
		Days:   days,
	}
}

// ageHandler manages the web interface and age calculation
func ageHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the HTML template
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Handle form submission
	if r.Method == http.MethodPost {
		// Parse the birthdate from the form
		birthdateStr := r.FormValue("birthdate")
		birthdate, err := time.Parse("2006-01-02", birthdateStr)

		var result AgeResult
		if err != nil {
			result.Error = "Invalid birthdate. Please select a valid date."
		} else {
			result = calculateAge(birthdate)
		}

		// Render the template with the result
		tmpl.Execute(w, result)
		return
	}

	// Render the initial template
	tmpl.Execute(w, nil)
}

func main() {
	// Set up the routes
	http.HandleFunc("/", ageHandler)

	// Log server start
	log.Println("Server starting on http://0.0.0.0:8080")
	log.Println("Access via localhost or your local network IP")

	// Start the server on all network interfaces
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
