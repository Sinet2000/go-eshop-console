package models

import (
	"fmt"

	"github.com/Sinet2000/go-eshop-console/domain/exceptions"
)

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	SKU         string  `json:"sku"`
	Price       float32 `json:"price"`
	Stock       int     `json:"stock"`
	ImageURL    string  `json:"imageUrl"`
}

func CreateProduct(id int, name, sku, description string, price float32, stock int, imageURL string) (*Product, error) {
	if name == "" {
		return nil, &exceptions.DomainException{Message: "Product name cannot be empty"}
	}
	if sku == "" {
		return nil, &exceptions.DomainException{Message: "SKU cannot be empty"}
	}
	if price <= 0 {
		return nil, &exceptions.DomainException{Message: "Price must be a positive value"}
	}
	if stock < 0 {
		return nil, &exceptions.DomainException{Message: "Stock cannot be negative"}
	}

	product := &Product{
		ID:          id,
		Name:        name,
		SKU:         sku,
		Description: description,
		Price:       price,
		Stock:       stock,
		ImageURL:    imageURL,
	}

	return product, nil
}

func (p *Product) UpdateStock(amount int) error {
	if amount < 0 && p.Stock+amount < 0 {
		return &exceptions.DomainException{Message: "Not enough stock to reduce"}
	}

	p.Stock += amount
	return nil
}

// SetPrice allows updating the product's price.
func (p *Product) SetPrice(newPrice float32) error {
	if newPrice <= 0 {
		return &exceptions.DomainException{Message: "Price must be a positive value"}
	}

	p.Price = newPrice
	return nil
}

func (p *Product) GetProductInfo() string {
	return fmt.Sprintf("Product: %s\nDescription: %s\nPrice: %.2f \nStock: %d\n", p.Name, p.Description, p.Price, p.Stock)
}
