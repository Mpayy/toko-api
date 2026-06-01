package web

type ProductUpdateRequest struct {
	Id    int    `validate:"required,numeric,min=1" json:"id"`
	Name  string `validate:"required,min=3,max=100" json:"name"`
	Price int    `validate:"required,gt=0,valid_price" json:"price"`
	Stock int    `validate:"required,min=0" json:"stock"`
}
