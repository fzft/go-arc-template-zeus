package pkg

import "gorm.io/gorm"

// CRUD defines methods for Create, Read, Update, and Delete operations.
type CRUD interface {
	Create(db *gorm.DB) error
	Read(db *gorm.DB, id uint) error
	Update(db *gorm.DB) error
	Delete(db *gorm.DB) error
}
