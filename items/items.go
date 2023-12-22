package items

import (
	"errors"
	"strconv"
)

type Product struct {
	ID              int     `json:"id"`
	Name            string  `json:"name"`
	Description     string  `json:"description"`
	LongDescription string  `json:"longdescription"`
	Price           float64 `json:"price"`
	Image           string  `json:"image"`
}

var products = &[]Product{
	{ID: 1, Name: "Generic Item 1", Description: "A generic item we sell", LongDescription: "A longer description of the generic item we sell", Price: 56.99},
	{ID: 2, Name: "Generic Item 2", Description: "A generic item we sell", LongDescription: "A longer description of the generic item we sell", Price: 57.99},
	{ID: 3, Name: "Generic Item 3", Description: "A generic item we sell", LongDescription: "A longer description of the generic item we sell", Price: 58.99},
	{ID: 4, Name: "Generic Item 4", Description: "A generic item we sell", LongDescription: "A longer description of the generic item we sell", Price: 59.99},
}

var nextId int = 5

func GetProducts() *[]Product {
	return products
}

func AddProduct(newProduct Product) (Product, error) {
	newProduct.ID = nextId
	*products = append(*products, newProduct)
	nextId++
	return newProduct, nil
}

func GetProduct(id string) (Product, error) {

	intID, err := strconv.Atoi(id)
	if err == nil {
		for _, product := range *products {
			if product.ID == intID {
				return product, nil
			}
		}
	}

	return Product{}, errors.New("Could not find product with that id.")

}
