package main

import (
	"fmt"
)

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
