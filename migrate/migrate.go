package main

import (
	"github.com/lghartmann/go-crud/initializers"
	"github.com/lghartmann/go-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
