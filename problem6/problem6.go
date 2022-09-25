package main

import "fmt"

func Fibo2(n int) int {
	if n < 2 {
		return n
	}
	return Fibo2(n-2) + Fibo2(n-1)

}

func main() {
	fmt.Println(Fibo2(0))
	fmt.Println(Fibo2(1))
	fmt.Println(Fibo2(2))
	fmt.Println(Fibo2(3))
	fmt.Println(Fibo2(5))
	fmt.Println(Fibo2(6))
	fmt.Println(Fibo2(7))
	fmt.Println(Fibo2(9))
	fmt.Println(Fibo2(10))
}
