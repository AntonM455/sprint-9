package main

import (
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements and is recorded in the created slides.
func generateRandomElements(size int) ([]int, error) {
	if size <= 0 {
		return nil, errors.New("size must be greater than 0")
	}

	data := make([]int, size)
	for i := range data {
		data[i] = rand.Intn(1_000_000) // 0 - 999 999
	}
	return data, nil
}

// maximum returns the maximum number of elements.
func maximum(data []int) (int, error) {
	if len(data) == 0 {
		return 0, errors.New("empty slice")
	}

	max := data[0]
	for _, v := range data[1:] {
		if v > max {
			max = v
		}
	}
	return max, nil
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) (int, error) {
	if len(data) == 0 {
		return 0, errors.New("empty slice")
	}

	chunkSize := len(data) / CHUNKS
	if chunkSize == 0 {
		return 0, errors.New("slice too small for chunking")
	}

	maxSlice := make([]int, CHUNKS)
	var wg sync.WaitGroup
	var mu sync.Mutex
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
			localMax, _ := maximum(chunk)
			mu.Lock()
			maxSlice[index] = localMax
			mu.Unlock()
		}(i, data[start:end])
	}

	wg.Wait()
	return maximum(maxSlice)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	start := time.Now()
	data, err := generateRandomElements(SIZE)
	if err != nil {
		fmt.Println("Data generation error:", err)
		return
	}
	fmt.Printf("Generation took: %d ms\n", time.Since(start).Milliseconds())

	fmt.Println("Ищем максимальное значение в один поток")
	start = time.Now()
	max1, err := maximum(data)
	if err != nil {
		fmt.Println("Maximum search error:", err)
		return
	}
	elapsed := time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение: %d\nВремя поиска: %d ms\n", max1, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max2, err := maxChunks(data)
	if err != nil {
		fmt.Println("Error finding maximum by parts:", err)
		return
	}
	elapsed = time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d ms\n", max2, elapsed)
}
