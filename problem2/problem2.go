package main

import "fmt"

func MoneyChange(money int) []int {
	listMoney := []int{10000, 5000, 2000, 1000, 500, 200, 100, 50, 20, 10, 1}
	var res []int

	if money < 2 {
		res = append(res, money)
	}

	i := 0
	for money >= 0 && i < len(listMoney) {
		if money-listMoney[i] >= 0 {
			res = append(res, listMoney[i])
			money -= listMoney[i]
		} else {
			i++
		}
	}
	return res
}

func main() {
	fmt.Println(MoneyChange(123))
	fmt.Println(MoneyChange(432))
	fmt.Println(MoneyChange(543))
	fmt.Println(MoneyChange(7752))
	fmt.Println(MoneyChange(15321))
}
