package main

import (
	"fmt"
	"sync"
)

func main() {
	// WithLocking()
	WithoutLocking()
}

func WithLocking() {
	firstData := []interface{}{
		"1",
		"2",
		"3",
		"4",
	}

	secondData := []interface{}{
		"1",
		"2",
		"3",
		"4",
	}

	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, v := range firstData {
		wg.Add(1)
		mu.Lock()
		go func(v interface{}) {
			fmt.Println(v)
			mu.Unlock()
			wg.Done()
		}(v)
	}

	for _, v := range secondData {
		wg.Add(1)
		mu.Lock()
		go func(v interface{}) {
			fmt.Println(v)
			mu.Unlock()
			wg.Done()
		}(v)
	}

	wg.Wait()
}

func WithoutLocking() {
	firstData := []interface{}{
		"1",
		"2",
		"3",
		"4",
	}

	secondData := []interface{}{
		"1",
		"2",
		"3",
		"4",
	}

	var wg sync.WaitGroup

	for _, v := range firstData {
		wg.Add(1)
		go func(v interface{}) {
			fmt.Println(v)
			wg.Done()
		}(v)
	}

	for _, v := range secondData {
		wg.Add(1)
		go func(v interface{}) {
			fmt.Println(v)
			wg.Done()
		}(v)
	}

	wg.Wait()
}
