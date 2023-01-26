// (B) Сумма к оплате (10 баллов)
// Полное решение: 10 баллов

package main
 
import (
	"bufio"
	"fmt"
	"os"
)
 
func main() {
	var a, b, t, sum int
	input := bufio.NewReader(os.Stdin)
	output := bufio.NewWriter(os.Stdout)
	defer output.Flush()
 
	fmt.Fscan(input, &a)
 
	for ; a > 0; a-- {
		basket := make(map[int]int)
		fmt.Fscan(input, &b)
		sum = 0
		// Read data from StdIn
		for ; b > 0; b-- {
			fmt.Fscan(input, &t)
			basket[t]++
		}
		// Calculate
		for price, amount := range basket {
			sum += (amount - amount/3) * price
		}
 
		fmt.Fprint(output, sum, "\n")
	