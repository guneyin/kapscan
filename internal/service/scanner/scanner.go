package scanner

import "gorm.io/gorm"

type Service struct {
	db *gorm.DB
}

func NewScannerService(db *gorm.DB) *Service {
	return &Service{
		db: db,
	}
}
