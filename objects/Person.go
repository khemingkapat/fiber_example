package object

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Age       int    `json:"age"`
	Role      string `json:"role"`
}
