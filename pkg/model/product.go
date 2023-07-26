package model

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string
	Description string
	Price       float64
}

// Create inserts the product into the database.
func (p *Product) Create(db *gorm.DB) error {
	return db.Create(p).Error
}

// Read populates the product from the database based on its ID.
func (p *Product) Read(db *gorm.DB, id uint) error {
	return db.First(p, id).Error
}

// Update updates the product in the database.
func (p *Product) Update(db *gorm.DB) error {
	return db.Save(p).Error
}

// Delete removes the product from the database.
func (p *Product) Delete(db *gorm.DB) error {
	return db.Delete(p).Error
}
