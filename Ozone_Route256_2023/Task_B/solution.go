// (B) Сумма к оплате (10 баллов), Частичное решение: 5 баллов.

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var queries, prices, price, sum int

	fmt.Scanln(&queries)
	answers := make([]string, 0, queries)

	for ; queries > 0; queries-- {
		basket := make([]int, 10_001, 10_001)
		fmt.Scan(&prices)
		sum = 0

		for ; prices > 0; prices-- {
			fmt.Scan(&price)
			basket[price]++
		}

		for price, amount := range basket {
			sum += (amount - amount/3) * price
		}

		answers = append(answers, strconv.Itoa(sum))
	}
	fmt.Println(strings.Join(answers, "\n"))
}
