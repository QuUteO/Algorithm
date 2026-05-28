package main

import "fmt"

func binary_search(array []int, value int) int {
	first := 0
	last := len(array) - 1

	for first <= last {
		mid := (first + last) / 2

		if array[mid] == value {
			return mid
		} else if array[mid] > value {
			last = mid - 1
		} else {
			first = mid + 1
		}
	}

	return -1
}

func main() {
	fmt.Println(binary_search([]int{1, 2, 3, 4, 5}, 1)) // 0
}
