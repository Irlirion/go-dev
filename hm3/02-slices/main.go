package main

import "fmt"

func Sum(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	fmt.Println(Sum(numbers))                // 30
	fmt.Println(Sum([]int{}))                // 0
	fmt.Println(Sum(nil))                    // 0
	fmt.Println(Sum([]int{-2, -1, 0, 1, 2})) // 0

	numbers = append(numbers, 12)
	fmt.Println(len(numbers), Sum(numbers)) // 6 42

	shared := numbers
	shared[0] = 100
	fmt.Println(numbers[0], shared[0]) // 100 100
}
