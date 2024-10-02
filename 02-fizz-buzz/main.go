package main

import "fmt"

func main() {
	var count int

	fmt.Print("Enter a number: ")
	fmt.Scan(&count)

	for i := 1; i <= count; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}
