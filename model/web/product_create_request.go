package web

type ProductCreateRequest struct {
	Name  string `validate:"required,max=100,min=1" json:"name"`
	Price int    `validate:"required" json:"price"`
	Stock int    `validate:"required" json:"stock"`
}
