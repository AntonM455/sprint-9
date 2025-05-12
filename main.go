package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements and is recorded in the created slides.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return []int{}
	}

	data := make([]int, size)
	for i := range data {
		data[i] = rand.Intn(math.MaxInt64)
	}
	return data
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	max := data[0]
	for _, v := range data[1:] {
		if v > max {
			max = v
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in chunks.
func maxChunks(data []int) int {
	chunkSize := len(data) / CHUNKS
	maxSlice := make([]int, CHUNKS)
	var wg sync.WaitGroup
	wg.Add(CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == CHUNKS-1 {
			end = len(data)
		}
		// launching goroutines
		go func(index int, chunk []int) {
			defer wg.Done()
			localMax := maximum(chunk)
			maxSlice[index] = localMax
		}(i, data[start:end])
	}

	wg.Wait()
	return maximum(maxSlice)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	start := time.Now()
	data := generateRandomElements(SIZE)
	fmt.Printf("Generation took: %d ms\n", time.Since(start).Milliseconds())

	fmt.Println("Ищем максимальное значение в один поток")
	start = time.Now()
	max1 := maximum(data)
	elapsed := time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение: %d\nВремя поиска: %d ms\n", max1, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max2 := maxChunks(data)
	elapsed = time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max2, elapsed)
}
