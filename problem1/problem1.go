package main

import "fmt"

func SimpleEquation(a, b, c int) string {
	var x, y, z int

	for x = 0; x < a; x++ {
		for y = 0; y < a; y++ {
			for z = 0; z < a; z++ {
				if x+y+z == a && x*y*z == b {
					return (fmt.Sprintf("%d %d %d", x, y, z))
				}
			}
		}
	}
	return ("no solution")
}
func main() {
	fmt.Println(SimpleEquation(1, 2, 3))
	fmt.Println(SimpleEquation(6, 6, 14))

}
