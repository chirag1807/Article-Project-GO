package dto

type User struct{
	ID       int64      `json:"id,string"`
	Name     string     `json:"name"`
	Bio      string     `json:"bio"`
	Email    string     `json:"email"`
	Image    *string 	`json:"image,omitempty"`
	IsAdmin  bool       `json:"isadmin,omitempty"`
}