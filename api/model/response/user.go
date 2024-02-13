package response

import (
	"database/sql"
	"encoding/json"
)

// NullString represents a string that may be null.
type NullString sql.NullString

// MarshalJSON marshals the NullString to JSON.
func (ns NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ns.String)
}

type User struct {
	ID       int64      `json:"id"`
	Name     string     `json:"name"`
	Bio      string     `json:"bio"`
	Email    string     `json:"email"`
	Password string     `json:"password,omitempty"`
	Image    *string 	`json:"image,omitempty"`
	IsAdmin  bool       `json:"isadmin,omitempty"`
}

type UserResponse struct {
	User         User   `json:"user"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
