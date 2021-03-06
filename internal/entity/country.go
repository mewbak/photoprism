package entity

import (
	"github.com/gosimple/slug"
	"github.com/jinzhu/gorm"
)

var altCountryNames = map[string]string{
	"United States of America": "USA",
	"United States":            "USA",
	"":                         "Unknown",
}

type Country struct {
	ID                 string `gorm:"primary_key"`
	CountrySlug        string
	CountryName        string
	CountryDescription string `gorm:"type:text;"`
	CountryNotes       string `gorm:"type:text;"`
	CountryPhoto       *Photo
	CountryPhotoID     uint
	New                bool `gorm:"-"`
}

// Create a new country
func NewCountry(countryCode string, countryName string) *Country {
	if countryCode == "" {
		countryCode = "zz"
	}

	if altName, ok := altCountryNames[countryName]; ok {
		countryName = altName
	}

	countrySlug := slug.MakeLang(countryName, "en")

	result := &Country{
		ID:          countryCode,
		CountryName: countryName,
		CountrySlug: countrySlug,
	}

	return result
}

func (m *Country) FirstOrCreate(db *gorm.DB) *Country {
	if err := db.FirstOrCreate(m, "id = ?", m.ID).Error; err != nil {
		log.Errorf("country: %s", err)
	}

	return m
}

func (m *Country) AfterCreate(scope *gorm.Scope) error {
	return scope.SetColumn("New", true)
}
