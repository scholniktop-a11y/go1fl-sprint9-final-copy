package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 10_000_000
	CHUNKS = 8
)

func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}
	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(1_000_000) + 1
	}
	return slice
}

func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	maxVal := data[0]
	for _, v := range data {
		if v > maxVal {
			maxVal = v
		}
	}
	return maxVal
}

func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}
	if len(data) < CHUNKS {
		return maximum(data)
	}

	chunkSize := len(data) / CHUNKS
	maxes := make([]int, CHUNKS)

	var wg sync.WaitGroup
	wg.Add(CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize

		if i == CHUNKS-1 {
			end = len(data)
		}

		go func(idx, s, e int) {
			defer wg.Done()
			maxes[idx] = maximum(data[s:e])
		}(i, start, end)
	}

	wg.Wait()
	return maximum(maxes)
}

func main() {
	data := generateRandomElements(SIZE)
	
	start := time.Now()
	maxSingle := maximum(data)
	durationSingle := time.Since(start)
	
	start = time.Now()
	maxMulti := maxChunks(data)
	durationMulti := time.Since(start)

	fmt.Printf("1 поток: %d (время: %v)\n", maxSingle, durationSingle.Microseconds())
	fmt.Printf("8 потоков: %d (время: %v)\n", maxMulti, durationMulti.Microseconds())
}