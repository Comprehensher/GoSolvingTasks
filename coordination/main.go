package main

import (
	"math"
	"math/rand"
	"sync"
	"time"
)

var waitGroup = sync.WaitGroup{}
var rwmutex = sync.RWMutex{}

var readyCond = sync.NewCond(rwmutex.RLocker())

// var readyCond = sync.NewCond(&rwmutex)

var squares = map[int]int{}

func generateSquares(max int) {
	rwmutex.Lock()
	Printfln("Generating data...")
	for val := 0; val < max; val++ {
		squares[val] = int(math.Pow(float64(val), 2))
	}
	rwmutex.Unlock()
	Printfln("Broadcasting condition")
	readyCond.Broadcast()
	waitGroup.Done()
}

func readSquares(id, max, iterations int) {
	readyCond.L.Lock()
	Printfln("Goroutine %v has got lock", id)
	for len(squares) == 0 {
		Printfln("Goroutine %v starts to wait...", id)
		readyCond.Wait()
		Printfln("Goroutine %v finish to wait...", id)
	}
	for i := 0; i < iterations; i++ {
		key := rand.Intn(max)
		Printfln("#%v Read value: %v = %v", id, key, squares[key])
		time.Sleep(time.Millisecond * 100)
	}
	readyCond.L.Unlock()
	Printfln("Goroutine %v released lock", id)
	waitGroup.Done()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	numRoutines := 2
	waitGroup.Add(numRoutines)
	for i := 0; i < numRoutines; i++ {
		go readSquares(i, 10, 5)
	}
	time.Sleep(time.Millisecond * 200)
	waitGroup.Add(1)
	go generateSquares(10)
	waitGroup.Wait()
	Printfln("Cached values: %v", len(squares))
}
