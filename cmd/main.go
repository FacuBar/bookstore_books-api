package main

import (
	"github.com/FacuBar/bookstore_books-api/pkg/infraestructure/clients"
	"github.com/FacuBar/bookstore_books-api/pkg/infraestructure/http/rest"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	db := clients.ConnectDB()

	server := rest.NewServer(db)

	server.Start(":8082")
}
