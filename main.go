package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const apiKey = "662ea841f1a740f7a638cc0f82f16860" // Replace with your actual API key
const apiUrl = "https://emailvalidation.abstractapi.com/v1/"

// Response 
type ApiResponse struct {
	Email             string               `json:"email"`
	Deliverability    string               `json:"deliverability"`
	IsValidFormat     struct{ Value bool } `json:"is_valid_format"`
	IsFreeEmail       struct{ Value bool } `json:"is_free_email"`
	IsDisposableEmail struct{ Value bool } `json:"is_disposable_email"`
	MxFound           struct{ Value bool } `json:"mx_found"`
	SmtpCheck         struct{ Value bool } `json:"smtp_check"`
}

func main() {
	fmt.Print("Enter the email address to validate: ")
	var email string
	fmt.Scanln(&email)
	if !isValidEmail(email) {
		fmt.Println("Invalid email format.")
		return
	}
	MakeRequest(email)
}

func MakeRequest(email string) {
	// URL with the email and API key
	url := fmt.Sprintf("%s?api_key=%s&email=%s", apiUrl, apiKey, email)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var apiResp ApiResponse
	err = json.Unmarshal(body, &apiResp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\nVerification results for %s:\n", email)
	fmt.Printf("Deliverability: %s\n", apiResp.Deliverability)
	fmt.Printf("Valid Format: %t\n", apiResp.IsValidFormat.Value)
	fmt.Printf("Free Email Provider: %t\n", apiResp.IsFreeEmail.Value)
	fmt.Printf("Disposable Email: %t\n", apiResp.IsDisposableEmail.Value)
	fmt.Printf("MX Records Found: %t\n", apiResp.MxFound.Value)
	fmt.Printf("SMTP Check: %t\n", apiResp.SmtpCheck.Value)
}

func isValidEmail(email string) bool {
	return strings.Contains(email, "@") && strings.Contains(email, ".")
}
