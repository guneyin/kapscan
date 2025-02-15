package entity

import (
	"time"

	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type Model struct {
	ID        ulid.ULID      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (m *Model) BeforeCreate(_ *gorm.DB) error {
	m.ID = ulid.Make()

	return nil
}
