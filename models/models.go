package model

import "time"

type User struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	CreatedAt time.Time     `json:"created_at"`
    UpdatedAt time.Time     `json:"updated_at"`
    CreditCards []*CreditCard `json:"credit_cards,omitempty"`

}

type CreditCard struct {
    ID      int    `json:"id"`
    Number  string `json:"number"`
	CreatedAt time.Time     `json:"created_at"`
    UpdatedAt time.Time     `json:"updated_at"`
}