// (E) Товарные знаки (20 баллов)
// Полное решение: 20 баллов

package main
 
import (
	"bufio"
	"fmt"
	"os"
	"sort"
)
 
func main() {
	var amountTests uint16
	var amountTMs int
	f := os.Stdin
	input := bufio.NewReader(f)
 
	fmt.Fscan(input, &amountTests)
 
	for ; amountTests > 0; amountTests-- {
		fmt.Fscan(input, &amountTMs)
		requests := make([]string, amountTMs, amountTMs)
 
		for i := 0; i < amountTMs; i++ {
			fmt.Fscan(input, &requests[i])
			if len(requests[i]) < 3 {
				continue
			}
			clearReq := make([]byte, 2, len(requests[i]))
			clearReq[0] = requests[i][0]
			clearReq[1] = requests[i][1]
			idx := 1
			for j := 2; j < len(requests[i]); j++ {
				if requests[i][j] == clearReq[idx] && requests[i][j] == clearReq[idx-1] {
					continue
				}
				clearReq = append(clearReq, requests[i][j])
				idx++
			}
			requests[i] = string(clearReq)
		}
		sort.Slice(requests, func(i, j int) bool { return requests[i] < requests[j] })
 
		res := amountTMs
		for i := 1; i < amountTMs; i++ {
			if requests[i] == requests[i-1] {
				res--
			}
		}
		fmt.Println(res)
	}
 
}
