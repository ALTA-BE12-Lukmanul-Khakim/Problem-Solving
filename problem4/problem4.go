package main

import (
	"fmt"
	"sort"
)

func BinarySearch(array []int, x int) int {
	sort.Ints(array)
	mid := len(array) / 2
	left := 0
	right := len(array) - 1
	ctr := 0

	for right >= left {
		ctr++
		if array[mid] == x {
			return mid
		} else if array[mid] > x {
			right = mid - 1
			mid = (left + right) / 2
		} else if array[mid] < x {
			left = mid + 1
			mid = (left + right) / 2
		}
	}
	return -1

}

func main() {
	fmt.Println(BinarySearch([]int{1, 1, 3, 5, 5, 6, 7}, 3))
	fmt.Println(BinarySearch([]int{1, 2, 3, 5, 6, 8, 10}, 5))
	fmt.Println(BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 53))
	fmt.Println(BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 100))
}
