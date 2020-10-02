package main

import (
	"fmt"
	"sort"
	"strconv"
)

const sliceSize = 3

func main() {
	var number string
	numbers := make([]int, sliceSize)

	// Press "X" to exit
	for number != "X" {
		fmt.Print("Please enter a number ('X' to exit): ")
		_, err := fmt.Scanf("%s", &number)
		if err == nil {
			numberI, _ := strconv.Atoi(number)
			numbers = append(numbers, numberI)
		}
	}
	numbers = numbers[sliceSize : len(numbers)-1]
	sort.Ints(numbers)
	fmt.Printf("Slice: %d", numbers)
}
