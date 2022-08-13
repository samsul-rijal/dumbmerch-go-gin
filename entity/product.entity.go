package entity

import "time"

type Product struct {
	ID     int64  `gorm:"primary_key:auto_increment" json:"-" `
	Name   string `gorm:"type: varchar(255)" json:"-"`
	Desc   string `gorm:"type:text" form:"desc"`
	Price  uint64 `gorm:"type:bigint" json:"-"`
	Image  string `gorm:"type: varchar(255)"`
	Qty    uint64 `gorm:"type:bigint" json:"-"`
	UserID int64  `gorm:"not null" json:"-"`
	User   User   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"-"`
	// CategoryID []int64    `gorm:"not null" json:"-"`
	// Category   []Category `gorm:"many2many:product_categories"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
