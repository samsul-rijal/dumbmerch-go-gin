package _product

import (
	"dumbmerch-go/entity"
	_user "dumbmerch-go/service/user"
)

type ProductResponse struct {
	ID    int64              `json:"id"`
	Name  string             `json:"name"`
	Price uint64             `json:"price"`
	Desc  string             `json:"desc"`
	Image string             `json:"image"`
	Qty   uint64             `json:"qty" `
	User  _user.UserResponse `json:"user,omitempty"`
}

func NewProductResponse(product entity.Product) ProductResponse {
	return ProductResponse{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
		Desc:  product.Desc,
		Image: product.Image,
		Qty:   product.Qty,
		User:  _user.NewUserResponse(product.User),
	}
}

func NewProductArrayResponse(products []entity.Product) []ProductResponse {
	productRes := []ProductResponse{}
	for _, v := range products {
		p := ProductResponse{
			ID:    v.ID,
			Name:  v.Name,
			Price: v.Price,
			Desc:  v.Desc,
			Image: v.Image,
			Qty:   v.Qty,
			User:  _user.NewUserResponse(v.User),
		}
		productRes = append(productRes, p)
	}
	return productRes
}
