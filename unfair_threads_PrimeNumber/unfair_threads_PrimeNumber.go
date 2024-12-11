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

var batchSize = int(NUMBER / int(CONCURRECY))

func main() {
	numstart := 3
	start := time.Now()
	var wg sync.WaitGroup
	for i := 0; i < int(CONCURRECY-1); i++ {
		wg.Add(1)
		go batch(strconv.Itoa(i), &wg, numstart, numstart+batchSize)
		numstart += batchSize
	}
	wg.Add(1)
	go batch(strconv.Itoa(CONCURRECY-1), &wg, numstart, NUMBER)

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

func batch(name string, wg *sync.WaitGroup, start int, end int) {
	defer wg.Done()
	t1 := time.Now()
	for i := start; i < end; i++ {
		checkPrime(i)
	}
	fmt.Printf("thread %s [%d, %d], took %v\n", name, start, end, time.Since(t1))
}
