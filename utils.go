package main

import (
	"fmt"
	"log"

	"github.com/khemingkapat/fiber_example/objects"
	"gorm.io/gorm"
)

var people = []object.Person{
	{Firstname: "Alice", Lastname: "Smith", Age: 30, Role: "Tenant"},
	{Firstname: "Bob", Lastname: "Johnson", Age: 25, Role: "Tenant"},
	{Firstname: "Charlie", Lastname: "Williams", Age: 35, Role: "Manager"},
	{Firstname: "David", Lastname: "Brown", Age: 28, Role: "Tenant"},
	{Firstname: "Eve", Lastname: "Davis", Age: 40, Role: "Manager"},
	{Firstname: "Frank", Lastname: "Miller", Age: 32, Role: "Tenant"},
	{Firstname: "Grace", Lastname: "Wilson", Age: 27, Role: "Tenant"},
	{Firstname: "Henry", Lastname: "Moore", Age: 38, Role: "Manager"},
	{Firstname: "Ivy", Lastname: "Taylor", Age: 29, Role: "Tenant"},
	{Firstname: "Jack", Lastname: "Anderson", Age: 33, Role: "Tenant"},
	{Firstname: "Karen", Lastname: "Thomas", Age: 42, Role: "Manager"},
	{Firstname: "Liam", Lastname: "Jackson", Age: 26, Role: "Tenant"},
	{Firstname: "Mia", Lastname: "White", Age: 31, Role: "Tenant"},
	{Firstname: "Noah", Lastname: "Harris", Age: 36, Role: "Manager"},
	{Firstname: "Olivia", Lastname: "Martin", Age: 29, Role: "Tenant"},
}

func populateDB(db *gorm.DB) {
	result := db.Create(&people)
	if err := result.Error; err != nil {
		log.Fatal(err)
	}
	fmt.Println("DB Populated")
}

func connStrSelect(mode string) string {
	var host string
	var port int

	if mode == "postgres" {
		host = "postgres"
		port = 5432
	} else if mode == "localhost" {
		host = "localhost"
		port = 5430
	} else {
		panic("Invalid Mode")
	}

	result := fmt.Sprintf("user=admin password=admin dbname=dorm host=%s port=%d sslmode=disable", host, port)
	return result
}
