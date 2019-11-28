package main

import "fmt"

func bubleSort(numbers []int) []int {
	for i := len(numbers); i > 0; i-- {
		for j := 1; j < i; j++ {
			if numbers[j-1] > numbers[j] {
				swap := numbers[j]
				numbers[j] = numbers[j-1]
				numbers[j-1] = swap
			}
		}
	}
	return numbers
}

func main() {
	sorted := bubleSort([]int{100, 20, 22, 13, 80, 33, 50, 1, 2, 81, 1231, 223, 45, 456, 123, 132})
	fmt.Println(sorted)
}
