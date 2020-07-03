package main

import (
	"fmt"
)

type Cart struct {
	List map[Product]float32
}

func (cart Cart) Checkout() {
	fmt.Println("== CHECKOUT ==")
	var total float32 
	for product, quantity := range cart.List {
		fmt.Println(product.Name, ":", quantity)
		total += product.Total(quantity)
	}
	fmt.Println("Total :", total)
}

type Product struct {
	Name string
	Price float32
}

func (product Product) Total(quantity float32) float32 {
	return product.Price * quantity
}

func main() {
	cart := Cart{ 
		List: map[Product]float32{
			Product{"arroz", 45.00}: 1.50,
			Product{"leche", 46.50}: 2.00,
			Product{"palta", 699.00}: 0.25,
			Product{"yerba", 120.33}: 3.00,
		}}
	cart.Checkout()
}