package main

import (
	"fmt"
	"math"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var CONCURRECY int = 10
var totalPrimes int32 = 0
var NUMBER int = 100000000
var currentNum int32 = 2
var batchSize = int(NUMBER / int(CONCURRECY))

func main() {
	numstart := 3
	start := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < int(CONCURRECY); i++ {
		wg.Add(1)
		go work(strconv.Itoa(i), &wg)
		numstart += batchSize
	}

	wg.Wait()
	fmt.Println("Total primes found: ", totalPrimes, "Time taken: ", time.Since(start))
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
	atomic.AddInt32(&totalPrimes, 1)
}

func work(name string, wg *sync.WaitGroup) {
	t1 := time.Now()
	defer wg.Done()
	for {
		x := atomic.AddInt32(&currentNum, 1)
		if x > int32(NUMBER) {
			break
		}
		checkPrime(int(x))
	}

	fmt.Printf("thread %s, took %v\n", name, time.Since(t1))
}
