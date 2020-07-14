package main

import (
	"fmt"
	"time"
	"math/rand"
	"sync"
)

type product float64

func (p product) Price() float64 {
	time.Sleep(time.Duration(rand.Intn(50)) * time.Millisecond)
	return float64(p)
}

func sumPrice(prod product, total *float64, wg *sync.WaitGroup) {

	defer func(wg *sync.WaitGroup) {
		wg.Done()
	}(wg)

 	*total += prod.Price()
}

func calculate(market string, products []product, wg *sync.WaitGroup, result map[string]*float64) {
	
	total := result[market]
	for _, product := range products {
		go sumPrice(product, total, wg)
	}
}

func main() {

	var wg sync.WaitGroup

	var totalChino float64
	var totalDia float64
	var totalCoto float64

	result := map[string]*float64 {
		"coto": &totalCoto,
		"dia": &totalDia,
		"chino": &totalChino,
	}

	markets := map[string][]product{
		"coto": []product{1.3, 6.3, 9.2},
		"dia": []product{10.3, 5.17, 39.04},
		"chino": []product{16.73, 6.20},
	}

	for market, products := range markets {
		wg.Add(len(products))
		go calculate(market, products, &wg, result)
		wg.Wait()
		fmt.Println(market, *result[market])
	}

}