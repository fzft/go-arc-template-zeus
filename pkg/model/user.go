package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string
	Email    string
	Password string // Note: this should be hashed, never store plaintext passwords
}

// Create inserts the user into the database.
func (u *User) Create(db *gorm.DB) error {
	return db.Create(u).Error
}

// Read populates the user from the database based on its ID.
func (u *User) Read(db *gorm.DB, id uint) error {
	return db.First(u, id).Error
}

// Update updates the user in the database.
func (u *User) Update(db *gorm.DB) error {
	return db.Save(u).Error
}

// Delete removes the user from the database.
func (u *User) Delete(db *gorm.DB) error {
	return db.Delete(u).Error
}
