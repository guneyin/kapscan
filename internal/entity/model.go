package entity

import (
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        ulid.ULID `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (m *Model) BeforeCreate(_ *gorm.DB) error {
	m.ID = ulid.Make()

	return nil
}
