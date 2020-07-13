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


func magic(products []product) {
	
	var wg sync.WaitGroup
	wg.Add(len(products))

	var total float64
	for _, product := range products {
		total += product.Price()
		wg.Done()
	}
	fmt.Println(total)

	wg.Wait()
}

func main() {

	markets := map[string][]product{
		"coto": []product{10.3, 56.77, 9.04},
		"dia": []product{10.3, 5.17, 39.04},
		"chino": []product{16.73, 6.20},
	}

	for market, products := range markets {
		fmt.Println(market)
		go magic(products)
	}




	/*
	var wg sync.WaitGroup
	wg.Add() // nro de rutinas

	// bla blabla
	go func() {
		defer wg.Done()
		// yad ayada yada
	}()

	*/
}