package main

import (
	"fmt"
	"strings"
	"time"
	"math"
)

type Cart struct {
	List []Item
}

func RoundDown (number float64) float64 {
	return math.Floor(number*100)/100
}

func (cart Cart) Checkout() {
	fmt.Println("============== CHECKOUT ==============", "\n")
	var total float64
	for _, item := range cart.List {
		totalItem := RoundDown(item.Total())
		fmt.Println("*", item.Product.GetName(), "\n", 
			"	Unit price :", item.Product.GetPrice(), "\n", 
			"	Quantity :", item.Quantity, "\n", 
			"	Discount:", RoundDown(1 - item.Product.GetDiscount()), "%", "\n", 
			"	Price with discount:", totalItem, "\n") 
		total += totalItem
	}
	fmt.Println("Total :", total)
}

type Vegetable struct {
	Name string
	Price float64
}

type Food struct {
	Name string
	Price float64
}

type Utility interface {
	GetName() string
	GetPrice() float64
	GetDiscount() float64
}

func (food Food) GetName() string {
	return food.Name
}

func (vegetable Vegetable) GetName() string {
	return vegetable.Name
}

func (food Food) GetPrice() float64 {
	return food.Price
}

func (vegetable Vegetable) GetPrice() float64 {
	return vegetable.Price
}

func (food Food) GetDiscount() float64 {
	if strings.HasPrefix(food.Name, "a") {
		return 0.7
	} else {
		return 1
	}
}

func (vegetable Vegetable) GetDiscount() float64 {
	if time.Now().Weekday().String() == "Friday" {
		return 0.85
	} else {
		return 1
	}
}

type Item struct {
	Product Utility
	Quantity float64
}

func (item Item) Total() float64 {
	return item.Product.GetPrice() * float64(item.Product.GetDiscount()) * item.Quantity
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