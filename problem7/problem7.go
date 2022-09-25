package main

import "fmt"

func Frog(jumps []int) int {
	var res int
	prev := 0
	prev2 := 0
	for i := 1; i < len(jumps); i++ {
		jump2 := 1000
		jump1 := prev + (jumps[i] - jumps[i-1])
		if i > 1 {
			jump2 = prev2 + (jumps[i] - jumps[i-2])
		}
		if jump1 > jump2 {
			res = jump1
		} else if jump1 < jump2 {
			res = jump2
		}

	}
	return res
}

func main() {
	fmt.Println(Frog([]int{10, 30, 40, 20}))
	fmt.Println(Frog([]int{30, 10, 60, 10, 60, 50}))
}
