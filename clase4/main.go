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

func sumPrice(prod product, total *float64, wg *sync.WaitGroup, mu *sync.Mutex) {

	defer func(mu *sync.Mutex, wg *sync.WaitGroup) {
		wg.Done()
		mu.Unlock()
	}(mu, wg)

	price := prod.Price()
	mu.Lock()
	*total += price
}

func calculate(market string, products []product, wg *sync.WaitGroup, mu *sync.Mutex, result map[string]*float64) {

	for _, product := range products {
		mu.Lock()
		total := result[market]
		mu.Unlock()
		go sumPrice(product, total, wg, mu)
	}
}

func main() {

	var wg sync.WaitGroup
	var mu sync.Mutex

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
		go calculate(market, products, &wg, &mu, result)
		wg.Wait()
		fmt.Println(market, *result[market])
	}

}