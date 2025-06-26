package models

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type ContentStatus string
type Language string

const (
	StatusInProgress ContentStatus = "in_progress"
	StatusFinish     ContentStatus = "finish"
	StatusFail       ContentStatus = "fail"

	LanguageArabic  Language = "ar"
	LanguageEnglish Language = "en"
)

type Content struct {
	ID              string        `json:"id" gorm:"primary_key;size:26"`
	Title           string        `json:"title" gorm:"size:36;not null;unique"`
	Description     string        `json:"description"`
	Language        Language      `json:"language"`
	Duration        float64       `json:"duration"`
	PublicationDate time.Time     `json:"publication_date"`
	Status          ContentStatus `json:"status"`
	S3Key           string        `json:"s3_key" gorm:"not null;unique"`
}

// BeforeCreate is a GORM hook that sets a ULID before inserting a new record
func (c *Content) BeforeCreate(tx *gorm.DB) (err error) {
	if c.ID == "" {
		t := time.Now().UTC()
		entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
		c.ID = ulid.MustNew(ulid.Timestamp(t), entropy).String()
	}
	return nil
}
