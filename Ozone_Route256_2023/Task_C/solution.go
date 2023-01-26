// (C) Парное программирование (10 баллов)
// Полное решение: 10 баллов

package main

import "fmt"

func main() {
	var amnt uint8
	fmt.Scanln(&amnt)
	for ; amnt > 0; amnt-- {
		solution()
	}
}

func solution() {
	var devs, first, second, dif, tmp int8
	fmt.Scan(&devs)
	skills := make([]int8, devs)

	for i := 0; i < int(devs); i++ {
		fmt.Scan(&skills[i])
	}

	if devs == 2 {
		fmt.Println(1, 2, "\n")
		return
	}
	second = 1

	for u := devs / 2; u > 1; u-- {
		dif = skills[first] - skills[second]

		if dif < 0 {
			dif = -dif
		}

		for i := first + 2; i < devs; i++ {
			if skills[i] == 0 {
				continue
			}
			tmp = skills[first] - skills[i]
			if tmp < 0 {
				tmp = -tmp
			}
			if tmp < dif {
				dif = tmp
				second = i
			}
			if dif == 0 {
				break
			}
		}
		skills[first] = 0
		skills[second] = 0
		first++
		fmt.Println(first, second+1)

		for ; first < devs; first++ {
			if skills[first] != 0 {
				break
			}
		}
		second = first + 1
		for ; second < devs; second++ {
			if skills[second] != 0 {
				break
			}
		}
	}
	fmt.Println(first+1, second+1, "\n")
}
