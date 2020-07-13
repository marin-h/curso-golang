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

func magic(products []product, total *float64, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, product := range products {
		*total += product.Price()
	}
}

func main() {

	markets := map[string][]product{
		"coto": []product{1.3, 6.3, 9.2},
		"dia": []product{10.3, 5.17, 39.04},
		"chino": []product{16.73, 6.20},
	}

	var total float64

	var wg sync.WaitGroup
	wg.Add(len(markets))

	for market, products := range markets {
		total = 0
		magic(products, &total, &wg)
		fmt.Println(market, total)
	}
	
	wg.Wait()
}