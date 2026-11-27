package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements генерирует слайс случайных целых чисел заданного размера.
func generateRandomElements(size int) []int {
	if size <= 0 {
		return nil
	}

	slice := make([]int, size)
	for i := 0; i < size; i++ {
		slice[i] = rand.Int()
	}
	return slice
}

// maximum находит максимальное значение в слайсе.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}
	maxVal := data[0]
	for _, val := range data {
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}

// maxChunks находит максимальное значение в слайсе, разделяя его на CHUNKS частей
func maxChunks(data []int) int {
	if len(data) < CHUNKS {
		return maximum(data)
	}

	maxValues := make([]int, CHUNKS)
	var wg sync.WaitGroup

	chunkSize := len(data) / CHUNKS
	wg.Add(CHUNKS)

	for i := 0; i < CHUNKS; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if i == CHUNKS-1 {
			end = len(data)
		}

		go func(index int, slice []int) {
			defer wg.Done()
			maxValues[index] = maximum(slice)
		}(i, data[start:end])
	}
	wg.Wait()

	return maximum(maxValues)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	res1 := maximum(data)
	elapsed1 := time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс (µs)\n", res1, elapsed1)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	res2 := maxChunks(data)
	elapsed2 := time.Since(start).Microseconds()
	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс (µs)\n", res2, elapsed2)
	// fmt.Println("res1 == res2", res1 == res2) // true
}
