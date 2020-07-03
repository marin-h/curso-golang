package main

import (
	"fmt"
	"strings"
	"time"
)

type Cart struct {
	List []Item
}

func (cart Cart) Checkout() {
	fmt.Println("============== CHECKOUT ==============", "\n")
	var total float32 
	for _, item := range cart.List {
		totalItem := item.Total()
		fmt.Println("*", item.Product.GetName(), "\n", 
			"	Unit price :", item.Product.GetPrice(), "\n", 
			"	Quantity :", item.Quantity, "\n", 
			"	Discount:", 1 - item.Product.GetDiscount(), "%", "\n", 
			"	Price with discount:", totalItem, "\n") 
		total += totalItem
	}
	fmt.Println("Total :", total)
}

type Vegetable struct {
	Name string
	Price float32
}

type Food struct {
	Name string
	Price float32
}

type Utility interface {
	GetName() string
	GetPrice() float32
	GetDiscount() float32
}

func (food Food) GetName() string {
	return food.Name
}

func (vegetable Vegetable) GetName() string {
	return vegetable.Name
}

func (food Food) GetPrice() float32 {
	return food.Price
}

func (vegetable Vegetable) GetPrice() float32 {
	return vegetable.Price
}

func (food Food) GetDiscount() float32 {
	if strings.HasPrefix(food.Name, "a") {
		return 0.7
	} else {
		return 1
	}
}

func (vegetable Vegetable) GetDiscount() float32 {
	if time.Now().Weekday().String() == "Friday" {
		return 0.85
	} else {
		return 1
	}
}

type Item struct {
	Product Utility
	Quantity float32
}

func (item Item) Total() float32 {
	return item.Product.GetPrice() * item.Product.GetDiscount() * item.Quantity
}

func main() {
	cart := Cart{ 
		List: []Item{
			Item{Food{"arroz", 45.00}, 1.50},
			Item{Food{"leche", 46.50}, 2.00},
			Item{Vegetable{"palta", 699.00}, 0.25},
			Item{Food{"yerba", 120.33}, 3.00},
		}}
	cart.Checkout()
}