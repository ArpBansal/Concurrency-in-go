package main

import (
	"fmt"
	"math"
	"time"
)

var NUMBER int = 100000000

var totalPrimes uint32 = 0

func main() {
	start := time.Now()
	for i := 3; i < NUMBER; i++ {
		checkPrime(i)
	}
	fmt.Println("Total primes: ", totalPrimes, "Time taken: ", time.Since(start))
	var a int = int(math.Sqrt(100))
	fmt.Printf("%d", a)
}

func checkPrime(x int) {
	if x&1 == 0 {
		return
	}
	for i := 2; i <= int(math.Sqrt(float64(x))); i++ {
		if x%i == 0 {
			return
		}
	}
	totalPrimes++
}
