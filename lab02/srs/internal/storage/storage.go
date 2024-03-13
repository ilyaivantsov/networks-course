package storage

import (
	"errors"
	"srs/internal/models"
)

var (
	ProductStorage          = make(map[int]models.Product)
	ErrProductAlreadyExists = errors.New("product already exists")
	ErrProductNotFound      = errors.New("product not found")
)

func CreateProduct(name, description string) (*models.Product, error) {
	product := models.New(len(ProductStorage)+1, name, description)
	ProductStorage[product.ID] = *product
	return product, nil
}

func UpdateProduct(id int, name, description string) (*models.Product, error) {
	product, ok := ProductStorage[id]
	if !ok {
		return nil, ErrProductNotFound
	}
	product.Name = name
	product.Description = description
	ProductStorage[id] = product
	return &product, nil
}

func GetProductByID(id int) (*models.Product, error) {
	product, ok := ProductStorage[id]
	if !ok {
		return nil, ErrProductNotFound
	}
	return &product, nil
}

func GetProducts() models.Products {
	products := make(models.Products, 0, len(ProductStorage))
	for _, product := range ProductStorage {
		products = append(products, product)
	}
	return products
}

func DeleteProduct(id int) (models.Product, error) {
	product, ok := ProductStorage[id]
	if !ok {
		return models.Product{}, ErrProductNotFound
	}
	delete(ProductStorage, id)
	return product, nil
}
