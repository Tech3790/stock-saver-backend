package models

import "gorm.io/gorm"

// User class
type User struct {
	gorm.Model
	FirstName string `json:"firstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"Email"`
}
