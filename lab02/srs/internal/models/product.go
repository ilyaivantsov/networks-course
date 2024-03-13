package models

type Product struct {
	ID          int    `json:"id"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type Products []Product

type ProductCreate struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type ProductUpdate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func New(id int, name, description string) *Product {
	return &Product{id, name, description}
}
