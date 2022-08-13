package entity

import "time"

type Profile struct {
	ID        int64  `gorm:"primary_key:auto_increment" json:"-"`
	Phone     string `gorm:"type: varchar(255)" json:"-"`
	Gender    string `gorm:"type: varchar(255)" json:"-"`
	Address   string `gorm:"type: text" json:"-"`
	UserID    int64  `gorm:"not null" json:"-"`
	User      User   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
