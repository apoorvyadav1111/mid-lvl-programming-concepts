package main

import (
	"fmt"
	"math"
	"strconv"
	"sync"
	"sync/atomic"
	"time"
)

var primeCnt int32 = 1
var CONCURRENCY = 10
var MAX_INT = 100000000
var curr int32 = 3

func checkPrime(x int) {
	if x%2 == 0 {
		return
	}
	for i := 2; i <= int(math.Sqrt(float64(x))); i++ {
		if x%i == 0 {
			return
		}
	}
	atomic.AddInt32(&primeCnt, 1)
}

func doSomeWork(name string, wg *sync.WaitGroup) {

	start := time.Now()
	defer wg.Done()

	for {
		x := atomic.AddInt32(&curr, 1)
		if x > int32(MAX_INT) {
			break
		}
		checkPrime(int(x))
	}
	fmt.Printf("thread %s completed in %s\n", name, time.Since(start))
}

func main() {
	start := time.Now()
	var wg sync.WaitGroup

	for i := 0; i < CONCURRENCY; i++ {
		wg.Add(1)
		go doSomeWork(strconv.Itoa(i), &wg)
	}
	wg.Wait()

	fmt.Println(
		"Sequential Checking till", MAX_INT, "found", primeCnt, "time taken", time.Since(start))
}
