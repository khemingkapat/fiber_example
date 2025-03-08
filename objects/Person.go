package object

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	Firstname string
	Lastname  string
	Age       int
	Role      string
}
