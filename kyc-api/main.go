package main

import (
	"fmt"
	"kyc-challenge/database"
	"kyc-challenge/models"
	"kyc-challenge/shared"
	"kyc-challenge/validations"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	loadEnvVars()
	dbConn := database.CreateDbConnection()
	dbConn.AutoMigrate(&models.ValidationData{}, &models.Countries{}, &models.DocType{})

	r := setupRouter(dbConn)
	database.UploadDataCountries(dbConn)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server running"))
	})

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	port := os.Getenv("HOST_PORT")
	fmt.Printf("Server running on %s\n", port)

	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}

	http.ListenAndServe(port, r)
}

// Cargar variables de entorno
func loadEnvVars() {
	if os.Getenv("GIN_MODE") != "release" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Unable to load env vars")
		}
	}
}

func setupRouter(dbConn *gorm.DB) *chi.Mux {
	r := chi.NewRouter()

	// Middleware de logging y CORS
	r.Use(shared.Cors)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)

	validations.AddValidationRoutes(r, dbConn)

	// Ruta para verificar conexión con la base de datos
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		tx := dbConn.Exec("SELECT 1")
		if tx.Error != nil {
			fmt.Printf("Error: %v\n", tx.Error)
			http.Error(w, "Database connection failed", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"Success": true}`))
	})

	return r
}
