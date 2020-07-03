package main

import (
	"fmt"
)

type Cart struct {
	List []Item
}

func (cart Cart) Checkout() {
	fmt.Println("== CHECKOUT ==")
	var total float32 
	for _, item := range cart.List {
		fmt.Println(item.Product.Name, ":", item.Quantity)
		total += item.Total()
	}
	fmt.Println("Total :", total)
}

type Utility struct {
	Name string
	Price float32
}

type Item struct {
	Product Utility
	Quantity float32
}

func (item Item) Total() float32 {
	return item.Product.Price * item.Quantity
}

func main() {
	cart := Cart{ 
		List: []Item{
			Item{Utility{"arroz", 45.00}, 1.50},
			Item{Utility{"leche", 46.50}, 2.00},
			Item{Utility{"palta", 699.00}, 0.25},
			Item{Utility{"yerba", 120.33}, 3.00},
		}}
	cart.Checkout()
}