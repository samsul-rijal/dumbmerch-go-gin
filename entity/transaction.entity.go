package entity

import "time"

type Transaction struct {
	ID        int64 `gorm:"primary_key:auto_increment" json:"-"`
	ProductID int64 `json:"-"`
	Product   Product
	BuyerID   int64 `gorm:"type:bigint" json:"-"`
	Buyer     User
	SellerID  int64 `gorm:"type:bigint" json:"-"`
	Seller    User
	Price     int64  `gorm:"type:bigint" json:"-"`
	Status    string `gorm:"type:varchar(25)" json:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
