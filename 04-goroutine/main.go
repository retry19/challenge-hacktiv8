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
		go print("bisa", v.(string), &wg, &mu)
	}

	for _, v := range secondData {
		wg.Add(1)
		mu.Lock()
		go print("coba", v.(string), &wg, &mu)
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
		go print("bisa", v.(string), &wg, nil)
	}

	for _, v := range secondData {
		wg.Add(1)
		go print("coba", v.(string), &wg, nil)
	}

	wg.Wait()
}

func print(prefix string, index string, wg *sync.WaitGroup, mu *sync.Mutex) {
	data := []interface{}{fmt.Sprintf("%s1", prefix), fmt.Sprintf("%s2", prefix), fmt.Sprintf("%s3", prefix)}

	fmt.Println(data, index)

	if mu != nil {
		mu.Unlock()
	}

	wg.Done()
}
