package dto

type CreateProductRequest struct {
	Name  string `json:"name" form:"name" binding:"required,min=1"`
	Price uint64 `json:"price" form:"email" binding:"required"`
	Desc  string `json:"desc" form:"desc" binding:"required"`
	Image string `json:"image" form:"image" binding:"required"`
	Qty   uint64 `json:"qty" form:"qty" binding:"required"`
}

type UpdateProductRequest struct {
	ID    int64  `json:"id" form:"id"`
	Name  string `json:"name" form:"name" binding:"required,min=1"`
	Price uint64 `json:"price" form:"email" binding:"required"`
	Desc  string `json:"desc" form:"desc" binding:"required"`
	Image string `json:"image" form:"image" binding:"required"`
	Qty   uint64 `json:"qty" form:"qty" binding:"required"`
}
