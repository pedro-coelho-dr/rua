package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/pedro-coelho-dr/rua/backend/internal/configs"
	"github.com/pedro-coelho-dr/rua/backend/internal/database"
	"github.com/pedro-coelho-dr/rua/backend/internal/routes"
)

func main() {
	configs.LoadConfig()
	database.Connect()

	r := chi.NewRouter()
	routes.SetupRoutes(r)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
