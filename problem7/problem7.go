package main

import "fmt"

func Frog(jumps []int) int {
	var res int
	var j, i int
	var minDistance int = 1000
	for i = 0; i < len(jumps); i++ {
		for j = i + 1; j < len(jumps); j++ {
			if jumps[i] == jumps[j] && j-i < minDistance {
				minDistance = j - i
			}
			if minDistance == 1000 {
				res = -1
			} else {
				res = minDistance
			}
		}
	}
	return res
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50}))
}
