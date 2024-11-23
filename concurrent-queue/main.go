package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type ConcurrentQueue struct {
	queue []int32
}

func (q *ConcurrentQueue) Enqueue(item int32) {
	mu.Lock()
	defer mu.Unlock()
	q.queue = append(q.queue, item)
}

func (q *ConcurrentQueue) Dequeue() int32 {
	mu.Lock()
	defer mu.Unlock()
	if len(q.queue) == 0 {
		panic("Queue Empty")
	}
	item := q.queue[0]
	q.queue = q.queue[1:]
	return item
}

func (q *ConcurrentQueue) Size() int {
	return len(q.queue)
}

var wg sync.WaitGroup

var mu sync.Mutex

func main() {
	q := ConcurrentQueue{
		queue: make([]int32, 0),
	}

	for i := 0; i < 1_000_000; i++ {
		wg.Add(1)
		go func() {
			q.Enqueue(rand.Int31())
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(q.Size())
}
