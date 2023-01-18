// (B) Сумма к оплате (10 баллов)
package main

import "fmt"

func main() {
	var a, b, t, sum int

	fmt.Scanln(&a)

	for ; a > 0; a-- {
		goods := make(map[int]int, 1_000_000)
		fmt.Scan(&b)
		sum = 0
		// Read data from StdIn
		for ; b > 0; b-- {
			fmt.Scan(&t)
			goods[t]++
		}
		// Calculate
		for price, amount := range goods {
			sum += (amount - amount/3) * price
		}

		fmt.Println(sum)
	}
}
