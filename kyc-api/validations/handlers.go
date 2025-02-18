package validations

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"kyc-challenge/models"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

func GetConfigurations(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var countries []models.Countries

	// Cargar países con sus documentos
	err := db.Preload("DocumentTypes").Find(&countries).Error
	if err != nil {
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}

	// Construcción del JSON con el formato requerido
	result := make(map[string]map[string][]string)
	for _, country := range countries {
		documentTypes := []string{}
		for _, doc := range country.DocumentTypes {
			documentTypes = append(documentTypes, doc.Name)
		}
		result[country.Name] = map[string][]string{
			"documentTypes": documentTypes,
		}
	}

	// Responder con JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func CreateValidation(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	var params models.ValidationParams
	apiKey := os.Getenv("API_SECRET_KEY")
	apiKeyHeader := os.Getenv("API_KEY_HEADER")

	if apiKey == "" {
		http.Error(w, "API key not set", http.StatusInternalServerError)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, "Invalid headers request", http.StatusBadRequest)
		return
	}

	// Construir el body en formato x-www-form-urlencoded
	formData := url.Values{}
	formData.Set("type", string(params.Type))
	formData.Set("country", string(params.Country))
	formData.Set("document_type", string(params.DocumentType))
	formData.Set("timeout", strconv.Itoa(params.Timeout))
	formData.Set("user_authorized", strconv.FormatBool(params.UserAuthorized))
	// formData.Set("account_id", os.Getenv("ACCOUNT_ID"))

	req, err := http.NewRequest("POST", os.Getenv("URL_VALIDATIONS"), strings.NewReader(formData.Encode()))
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set(apiKeyHeader, apiKey)

	// Hacer la petición a la API de POST Validation
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to call external API", http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	var apiResponse models.CreateValidationResult
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		http.Error(w, "Failed to parse API response", http.StatusInternalServerError)
		return
	}

	dataToSave := models.ValidationData{
		ValidationID: apiResponse.ValidationID,
		AccountID:    apiResponse.AccountID,
		FrontURL:     apiResponse.Instructions.FrontURL,
		ReverseURL:   apiResponse.Instructions.ReverseURL,
		CreationDate: apiResponse.CreationDate,
	}

	response := models.SessionInfo{
		AccountID:    apiResponse.AccountID,
		CreationDate: apiResponse.CreationDate,
	}

	if err := db.Create(&dataToSave).Error; err != nil {
		http.Error(w, "Failed to store in database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func UploadImage(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	accountId := r.URL.Query().Get("account_id")
	imageType := r.URL.Query().Get("image_type")

	if accountId == "" {
		http.Error(w, "Missing accountId", http.StatusBadRequest)
		return
	}

	if imageType != "front_url" && imageType != "reverse_url" {
		http.Error(w, "Invalid image type", http.StatusBadRequest)
		return
	}

	// Obtener imagen front

	err := r.ParseMultipartForm(10 << 20) // Limitar a 10 MB
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to get file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	imageData, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusBadRequest)
		return
	}

	// Obtener imagen postman
	// imageFile, err := io.ReadAll(r.Body)
	// if err != nil {
	// 	http.Error(w, "Failed to get file", http.StatusBadRequest)
	// 	return
	// }
	// defer r.Body.Close()

	var dbData models.ValidationData
	var columnName string
	if imageType == "front_url" {
		columnName = "front_url"
	} else {
		columnName = "reverse_url"
	}

	err = db.Where("account_id = ?", accountId).First(&dbData).Error
	if err != nil {
		http.Error(w, "Data not found for given accountId", http.StatusNotFound)
		return
	}

	var urlToUse string
	if columnName == "front_url" {
		urlToUse = dbData.FrontURL
	} else {
		urlToUse = dbData.ReverseURL
	}

	// req, err := http.NewRequest("PUT", urlToUse, bytes.NewReader(imageFile)) --> postman
	req, err := http.NewRequest("PUT", urlToUse, bytes.NewReader(imageData))
	if err != nil {
		http.Error(w, "Failed to create PUT request", http.StatusInternalServerError)
		return
	}

	req.Header.Set("Content-Type", "text/plain")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to send PUT request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(resp.StatusCode)
	w.Write(respBody)
}

func GetValidation(w http.ResponseWriter, r *http.Request, db *gorm.DB) {
	apiKey := os.Getenv("API_SECRET_KEY")
	apiKeyHeader := os.Getenv("API_KEY_HEADER")

	if apiKey == "" {
		http.Error(w, "API key not set", http.StatusInternalServerError)
		return
	}

	accountId := r.URL.Query().Get("account_id")
	showDetailsStr := r.URL.Query().Get("show_details")

	if accountId == "" {
		http.Error(w, "Missing accountId", http.StatusBadRequest)
		return
	}

	if showDetailsStr == "" {
		showDetailsStr = "false"
	}

	var dbData models.ValidationData
	err := db.Where("account_id = ?", accountId).First(&dbData).Error
	if err != nil {
		http.Error(w, "Data not found for given accountId", http.StatusNotFound)
		return
	}

	var urlToUse = os.Getenv("URL_VALIDATIONS") + "/" + dbData.ValidationID + "?show_details=" + showDetailsStr

	fmt.Println("url: ", urlToUse)
	req, err := http.NewRequest("GET", urlToUse, nil)
	if err != nil {
		http.Error(w, "Failed to create GET request", http.StatusInternalServerError)
		return
	}

	req.Header.Set(apiKeyHeader, apiKey)

	// Enviar la solicitud
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to send GET request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var apiResponse models.DocumentValidationResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		http.Error(w, "Failed to parse API response", http.StatusInternalServerError)
		return
	}

	response := models.DocumentInfo{
		ValidationStatus:     apiResponse.ValidationStatus,
		FailureStatus:        apiResponse.FailureStatus,
		DeclinedReason:       apiResponse.DeclinedReason,
		ProcessingStartDate:  apiResponse.ProcessingStartDate,
		ProcessingFinishDate: apiResponse.ProcessingFinishDate,
		Details:              apiResponse.Details,
		UserResponse:         apiResponse.UserResponse,
	}

	fmt.Printf("UserResponse: %+v\n", response.UserResponse)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
