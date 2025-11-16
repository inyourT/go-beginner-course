package main

import (
	"fmt"
)

type Product struct {
	ID       string
	Name     string
	Price    float64
	Stock    int
	Category string
	Discount float64
}

func (p Product) Info() {
	fmt.Printf("ID: %s\n Name: %s\n Category: %s\n Price: %.2f\nDiscount: %.2f\nStock: %d\n",
		p.ID, p.Name, p.Category, p.Price, p.Discount, p.Stock)
}

func (p Product) FinalPrice() float64{
	if p.Discount <= 0 {
		return p.Price
	}
	return p.Price - (p.Price * p.Discount / 100)
}

func (p Product) IsAvailable() bool{
	return p.Stock > 0
}

func (p *Product) ReduceStock(qty int) bool {
	if qty <= 0{
		return false
	}

	if p.Stock < qty {
		return false
	}

	p.Stock -= qty 
		return true
}

func (p *Product) AddStock(qty int) bool {
	if qty <= 0 {
		return  false
	}

	p.Stock += qty
	return true
}

type CartItem struct {
	Product *Product
	Quantity int
}

type Cart struct{
	Items []CartItem
}

func (c *Cart) AddProduct(product *Product, qty int) bool {
	if qty <= 0 {
		return  false
	}

	if product.Stock < qty {
		return  false
	}

	// cek apakah product sudah ada di cart
	for i, item := range c.Items {
		if item.Product.ID == product.ID {
			c.Items[i].Quantity += qty

			product.ReduceStock(qty)
			return  true
		}
	}

	newItem := CartItem {
		Product: product,
		Quantity: qty,
	}

	c.Items = append(c.Items, newItem)

	product.ReduceStock(qty)
	return  true
}

func (c *Cart) RemoveProduct(productID string) {
	for i, item := range c.Items {
		if item.Product.ID == productID {

			// kembali ke stock produk
			item.Product.AddStock(item.Quantity)

			c.Items = append(c.Items[:i], c.Items[i+1:]...)
			return
		} 
	}
}

func (c*Cart) TotalPrice() float64 {
	var total float64

	for _, item := range c.Items{
		total += float64(item.Quantity) * item.Product.Price
	}

	return total
}

func main() {
	
}