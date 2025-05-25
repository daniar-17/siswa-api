package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Student struct {
	ID    uint           `gorm:"primaryKey"`
	UUID  uuid.UUID      `gorm:"type:char(36);uniqueIndex"`
	Nama  string
	NIS   string `gorm:"unique"`
	Foto  string
}

func (s *Student) BeforeCreate(tx *gorm.DB) (err error) {
	s.UUID = uuid.New()
	return
}
