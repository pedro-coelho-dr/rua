package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/pedro-coelho-dr/rua/backend/internal/configs"
	"github.com/pedro-coelho-dr/rua/backend/internal/database"
	"github.com/pedro-coelho-dr/rua/backend/internal/routes"
)

func main() {
	config.LoadConfig()

	dsn := "host=" + os.Getenv("DB_HOST") + " user=" + os.Getenv("DB_USER") + " password=" + os.Getenv("DB_PASSWORD") + " dbname=" + os.Getenv("DB_NAME") + " port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	database.DB = db

	r := chi.NewRouter()
	routes.SetupRoutes(r)

	log.Println("Server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
