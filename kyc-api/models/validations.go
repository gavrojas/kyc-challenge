package models

import (
	"time"

	"gorm.io/gorm"
)

// Enums
type TypeValidation string
type Country string
type DocumentType string

const (
	DocumentValidation TypeValidation = "document-type"

	Colombia   Country = "CO"
	Chile      Country = "CL"
	Mexico     Country = "MX"
	Peru       Country = "PE"
	BRASIL     Country = "BR"
	CostaRica  Country = "CR"
	CountryAll Country = "ALL"

	National      DocumentType = "national-id"
	Foreign       DocumentType = "foreign-id"
	DriverLicense DocumentType = "driver-license"
	Passport      DocumentType = "passport"
	IdentityCard  DocumentType = "identity-card"
	RUT           DocumentType = "rut"
	PPT           DocumentType = "ppt"
	Invoice       DocumentType = "invoice"
	PictureID     DocumentType = "picture-id"
	Record        DocumentType = "record"
	CNH           DocumentType = "cnh"
)

type ValidationParams struct {
	Type           TypeValidation `json:"type"`
	UserAuthorized bool           `json:"user_authorized"`
	Country        Country        `json:"country"`
	DocumentType   DocumentType   `json:"document_type"`
	Timeout        int            `json:"timeout"`
}

type CreateValidationResult struct {
	ValidationID     string    `json:"validation_id"`
	IPAddress        string    `json:"ip_address"`
	AccountID        string    `json:"account_id"`
	Type             string    `json:"type"`
	ValidationStatus string    `json:"validation_status"`
	CreationDate     time.Time `json:"creation_date"`

	Instructions struct {
		FrontURL   string `json:"front_url"`
		ReverseURL string `json:"reverse_url"`
	} `json:"instructions"`
}

type ValidationData struct {
	gorm.Model   `json:"-"`
	ValidationID string    `json:"validation_id"`
	AccountID    string    `json:"account_id"`
	CreationDate time.Time `json:"creation_date"`
	FrontURL     string    `json:"front_url"`
	ReverseURL   string    `json:"reverse_url"`
}

type SessionInfo struct {
	AccountID    string    `json:"account_id"`
	CreationDate time.Time `json:"creation_date"`
}

// Respuesta final
type DocumentValidationResponse struct {
	Type                   string           `json:"type"`
	ValidationStatus       string           `json:"validation_status"`
	FailureStatus          string           `json:"failure_status"`
	DeclinedReason         string           `json:"declined_reason"`
	ExpectedExpirationDate time.Time        `json:"expected_expiration_date"`
	CreationDate           time.Time        `json:"creation_date"`
	ProcessingStartDate    time.Time        `json:"processing_start_date"`
	ProcessingFinishDate   time.Time        `json:"processing_finish_date"`
	Details                Details          `json:"details"`
	ValidationInputs       ValidationInputs `json:"validation_inputs"`
	DocumentExpectedPages  int              `json:"document_expected_pages"`
	UserResponse           UserResponse     `json:"user_response"`
	ValidationID           string           `json:"validation_id"`
	IPAddress              string           `json:"ip_address"`
	AccountID              string           `json:"account_id"`
	AttachmentStatus       string           `json:"attachment_status"`
}

type Details struct {
	DocumentDetails DocumentDetails `json:"document_details"`
}

type DocumentDetails struct {
	BirthPlace        string    `json:"birth_place"`
	ClientID          string    `json:"client_id"`
	Country           string    `json:"country"`
	CreationDate      time.Time `json:"creation_date"`
	DateOfBirth       time.Time `json:"date_of_birth"`
	DocID             string    `json:"doc_id"`
	DocumentType      string    `json:"document_type"`
	DocumentVersion   string    `json:"document_version"`
	ExpeditionPlace   string    `json:"expedition_place"`
	Gender            string    `json:"gender"`
	Height            string    `json:"height"`
	IssueDate         time.Time `json:"issue_date"`
	MimeType          string    `json:"mime_type"`
	NationalRegistrar string    `json:"national_registrar"`
	ProductionData    string    `json:"production_data"`
	RH                string    `json:"rh"`
	UpdateDate        time.Time `json:"update_date"`
}

type ValidationInputs struct {
	Country      string `json:"country"`
	DocumentType string `json:"document_type"`
}

type UserResponse struct {
	InputFiles []string `json:"input_files"`
}

type DocumentInfo struct {
	ValidationStatus     string       `json:"validation_status"`
	FailureStatus        string       `json:"failure_status"`
	DeclinedReason       string       `json:"declined_reason"`
	ProcessingStartDate  time.Time    `json:"processing_start_date"`
	ProcessingFinishDate time.Time    `json:"processing_finish_date"`
	Details              Details      `json:"details"`
	UserResponse         UserResponse `json:"user_response"`
}
