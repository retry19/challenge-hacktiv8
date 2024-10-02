package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "selamat datang"
	slice := strings.Split(text, "")

	mapSlice := map[string]int{}

	for _, letter := range slice {
		fmt.Println(letter)

		if mapSlice[letter] < 1 {
			mapSlice[letter] = 1
		} else {
			mapSlice[letter] += 1
		}
	}

	fmt.Printf("%v\n", mapSlice)
}
