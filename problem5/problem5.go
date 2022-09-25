package main

import "fmt"

func Fibo(n int) int {
	if n < 2 {
		return n
	}
	return Fibo(n-2) + Fibo(n-1)

}

func main() {
	fmt.Println(Fibo(0))
	fmt.Println(Fibo(1))
	fmt.Println(Fibo(2))
	fmt.Println(Fibo(3))
	fmt.Println(Fibo(5))
	fmt.Println(Fibo(6))
	fmt.Println(Fibo(7))
	fmt.Println(Fibo(9))
	fmt.Println(Fibo(10))
}
