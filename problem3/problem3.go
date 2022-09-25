package main

import "fmt"

func DragonOfLoowater(dragonHead, knightHeight []int) {
	var res int
	//var save []int
	min := 1000
	for _, height := range knightHeight {
		if height > dragonHead[0] || height > dragonHead[1] {
			res = res + height
			if res < min {
				min = res
			} else {
				height = 0
			}
			fmt.Println(res)
		}
		if height < dragonHead[0] || height < dragonHead[1] {
			fmt.Println(" Faill ")
		}
	}

}

func main() {

	DragonOfLoowater([]int{5, 4}, []int{7, 8, 4})
	// DragonOfLoowater([]int{5, 10}, []int{5})
	// DragonOfLoowater([]int{7, 2}, []int{4, 3, 1, 2})

}
