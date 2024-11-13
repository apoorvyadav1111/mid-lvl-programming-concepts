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

func HandleBatch(name string, wg *sync.WaitGroup, nstart int, nend int) {
	defer wg.Done()

	start := time.Now()
	for i := nstart; i < nend; i++ {
		checkPrime(i)
	}
	fmt.Printf("thread %s [%d, %d) completed in %s\n", name, nstart,
		nend, time.Since(start))
}

func main() {
	start := time.Now()
	var wg sync.WaitGroup
	nstart := 3
	batchSize := int(float64(MAX_INT) / float64(CONCURRENCY))

	for i := 0; i < CONCURRENCY-1; i++ {
		wg.Add(1)
		go HandleBatch(strconv.Itoa(i), &wg, nstart, nstart+batchSize)
		nstart += batchSize
	}

	wg.Add(1)
	go HandleBatch(strconv.Itoa(CONCURRENCY-1), &wg, nstart, nstart+batchSize)
	wg.Wait()

	fmt.Println(
		"Sequential Checking till", MAX_INT, "found", primeCnt, "time taken", time.Since(start))
}
