package database

import (
	"fmt"
	"kyc-challenge/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func CreateDbConnection() *gorm.DB {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			IgnoreRecordNotFoundError: true, // ignorar errores de "record not found"
		})

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func UploadDataCountries(db *gorm.DB) {
	countries := map[string][]string{
		"Colombia":   {"driver-license", "foreign-id", "identity-card", "national-id", "passport", "ppt", "rut"},
		"Chile":      {"driver-license", "foreign-id", "national-id", "passport"},
		"Mexico":     {"foreign-id", "invoice", "national-id", "passport", "picture-id", "record"},
		"Peru":       {"national-id"},
		"Brasil":     {"cnh"},
		"Costa Rica": {"driver-license", "passport"},
		"Other":      {"passport"},
	}

	for countryName, docTypes := range countries {
		var country models.Countries

		// Asegurarse de obtener correctamente el país y su ID
		result := db.Where("name = ?", countryName).FirstOrCreate(&country, models.Countries{Name: countryName})
		if result.Error != nil {
			log.Printf("Error inserting country %s: %v", countryName, result.Error)
			continue
		}

		// Ahora country.ID debería estar correctamente asignado
		for _, docType := range docTypes {
			var doc models.DocType
			db.Where("name = ? AND country_id = ?", docType, country.ID).
				FirstOrCreate(&doc, models.DocType{Name: docType, CountryID: country.ID})
		}
	}
}
