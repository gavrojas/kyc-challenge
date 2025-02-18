package models

type Countries struct {
	ID            uint      `gorm:"primaryKey"`
	Name          string    `gorm:"unique;not null"`
	DocumentTypes []DocType `gorm:"foreignKey:CountryID"`
}

type DocType struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	CountryID uint   `gorm:"index"`
}
