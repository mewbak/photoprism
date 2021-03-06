package entity

import (
	"github.com/gosimple/slug"
	"github.com/jinzhu/gorm"
)

// Camera lens (as extracted from UpdateExif metadata)
type Lens struct {
	Model
	LensSlug        string
	LensModel       string
	LensMake        string
	LensType        string
	LensOwner       string
	LensDescription string `gorm:"type:text;"`
	LensNotes       string `gorm:"type:text;"`
}

func (Lens) TableName() string {
	return "lenses"
}

func NewLens(modelName string, makeName string) *Lens {
	if modelName == "" {
		modelName = "Unknown"
	}

	lensSlug := slug.MakeLang(modelName, "en")

	result := &Lens{
		LensModel: modelName,
		LensMake:  makeName,
		LensSlug:  lensSlug,
	}

	return result
}

func (m *Lens) FirstOrCreate(db *gorm.DB) *Lens {
	if err := db.FirstOrCreate(m, "lens_model = ? AND lens_make = ?", m.LensModel, m.LensMake).Error; err != nil {
		log.Errorf("lens: %s", err)
	}

	return m
}
