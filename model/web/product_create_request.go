package web

type ProductCreateRequest struct {
	Name  string `validate:"required,min=3,max=100" json:"name"`
	Price int    `validate:"required,gt=0,valid_price" json:"price"`
	Stock int    `validate:"required,min=0" json:"stock"`
}
