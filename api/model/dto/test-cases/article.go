package testcase

type Article struct {
	Message    string `json:"message,omitempty"`
	StatusCode *int   `json:"statuscode,omitempty"`
}
