package main

import (
	"fmt";
	"math/rand";
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(101);
	loops := Loop(n);
	fmt.Println(fmt.Sprintln("You tried %s times!", loops));
}

func Loop(n int) int {
	var guess int
	i := 0
	fmt.Println("Guess the number")
	for n != guess {
		i++
		fmt.Scanln(&guess)
		if guess > n {
			fmt.Println("Guess is too high")
		} else if guess < n {
			fmt.Println("Guess is too low")
		}
	}
	return i
}