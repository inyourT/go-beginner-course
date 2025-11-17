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

func (p Product) FinalPrice() float64 {
	if p.Discount <= 0 {
		return p.Price
	}
	return p.Price - (p.Price * p.Discount / 100)
}

func (p Product) IsAvailable() bool {
	return p.Stock > 0
}

func (p *Product) ReduceStock(qty int) bool {
	if qty <= 0 {
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
		return false
	}

	p.Stock += qty
	return true
}

type CartItem struct {
	Product  *Product
	Quantity int
}

type Cart struct {
	Items []CartItem
}

func (c *Cart) AddProduct(product *Product, qty int) bool {
	if qty <= 0 {
		return false
	}

	if product.Stock < qty {
		return false
	}

	// cek apakah product sudah ada di cart
	for i, item := range c.Items {
		if item.Product.ID == product.ID {
			c.Items[i].Quantity += qty

			product.ReduceStock(qty)
			return true
		}
	}

	newItem := CartItem{
		Product:  product,
		Quantity: qty,
	}

	c.Items = append(c.Items, newItem)

	product.ReduceStock(qty)
	return true
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

func (c *Cart) TotalPrice() float64 {
	var total float64

	for _, item := range c.Items {
		total += float64(item.Quantity) * item.Product.Price
	}

	return total
}

func (c *Cart) TotalQuantity() int {
	var total int

	for _, item := range c.Items {
		total += 	item.Quantity
	}

	return total
}

func (c *Cart) ViewCart(){
	fmt.Println("===Cart Content===")
	if len(c.Items) == 0 {
		fmt.Println("Cart masih kosong")
		return
	}

	for _, item := range c.Items{
		fmt.Printf("- %s (%d pcs) Harga Satuan: %.0f Harga Akhir: %.0f\n",
			item.Product.Name,
			item.Quantity,
			item.Product.Price,
			item.Product.FinalPrice(),
	)
	}

	fmt.Println("-----------------------")
	fmt.Printf("Total Quantity: %d\n", c.TotalQuantity())
	fmt.Printf("Total Price: Rp %.0f\n", c.TotalPrice())
	fmt.Println()
}

// checkout

func (c *Cart) checkout() {
	if len(c.Items) == 0 {
		fmt.Println("Tidak ada checkout, cart kosong")
		return
	}

	fmt.Println("Checkout berhasil")
	c.Items = []CartItem{}
}

// Cancel Checkout
func (c *Cart) CancelCheckout() {
	for _, item := range c.Items {
		item.Product.AddStock(item.Quantity)
	}

	c.Items = []CartItem{}
	fmt.Println("Checkout dibatalkan! stock telah dikembalikan.")
}


func main() {
	p1 := &Product{"P001", "Keyboard", 150000, 10, "Elektronik", 10}
	p2 := &Product{"P002", "Mouse", 80000, 20, "Elektronik", 0}
	p3 := &Product{"P003", "Headset", 200000, 5, "Elektronik", 5}

	cart := Cart{}

	cart.AddProduct(p1, 2)
	cart.AddProduct(p2, 3)
	cart.AddProduct(p2, 1)
	cart.AddProduct(p3, 1)

	cart.ViewCart()

	cart.RemoveProduct("P002")
	cart.ViewCart()

	cart.checkout()
	cart.CancelCheckout()
}
