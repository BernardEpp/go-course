package main

import "fmt"

func main() {
	numbers := []int{}

	i := 0

	for i < 10 {
		numbers = append(numbers, i)
		if i%2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}

		i++
	}

	// fmt.Println(numbers)
}
