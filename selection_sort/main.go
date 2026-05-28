package main

import "fmt"

func selectionSearch(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		minIdx := i

		for j := i + 1; j < len(arr); j++ {
			if arr[j] <= arr[minIdx] {
				minIdx = j
			}
		}

		arr[i], arr[minIdx] = arr[minIdx], arr[i]
	}

	return arr
}

func main() {
	fmt.Println(selectionSearch([]int{54, 23, 1, 4351, 5668})) // [1 23 54 4351 5668]
}
