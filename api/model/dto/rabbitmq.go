package dto

type UserMail struct {
	To string `json:"to"`
	Subject string `json:"subject"`
	Body string `json:"body"`
}
