// Полное решение: 5 баллов.

package main

import "fmt"

func main() {
	var a, x, y int
	fmt.Scanln(&a)
	for ; a > 0; a-- {
		fmt.Scanln(&x, &y)
		fmt.Println(x + y)
	}
}
