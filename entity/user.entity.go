package entity

import "time"

type User struct {
	ID       int64  `gorm:"primary_key:auto_increment" json:"-"`
	Name     string `gorm:"type: varchar(255)" json:"-"`
	Email    string `gorm:"type: varchar(255)" json:"-"`
	Password string `gorm:"type: varchar(255)" json:"-"`
	// Profile   Profile
	Products  []Product
	CreatedAt time.Time
	UpdatedAt time.Time
}
