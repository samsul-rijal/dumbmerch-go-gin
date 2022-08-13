package entity

import "time"

type Category struct {
	ID        int64  `gorm:"primary_key:auto_increment" json:"-"`
	Name      string `gorm:"type: varchar(255)" json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
