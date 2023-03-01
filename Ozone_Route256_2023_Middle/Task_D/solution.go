// (D) Результаты соревнования (20 баллов)
// Полное решение: 20 баллов

package main
 
import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)
 
type Player struct {
	res   int
	index int
	place int
}
 
func main() {
	var a uint16
	var amount, res int
	builder := strings.Builder{}
	input := bufio.NewReader(os.Stdin)
	output := bufio.NewWriter(os.Stdout)
	defer output.Flush()
 
	fmt.Fscan(input, &a)
 
	for ; a > 0; a-- {
		fmt.Fscan(input, &amount)
		players := make([]*Player, amount, amount)
 
		for i := 0; i < amount; i++ {
			fmt.Fscan(input, &res)
			players[i] = &Player{res: res, index: i}
		}
 
		sort.Slice(players, func(i, j int) bool { return players[i].res < players[j].res })
		players[0].place = 1
 
		for i := 1; i < amount; i++ {
			if players[i].res <= players[i-1].res+1 {
				players[i].place = players[i-1].place
			} else {
				players[i].place = i + 1
			}
		}
 
		out := make([]int, amount, amount)
		for _, p := range players {
			out[p.index] = p.place
		}
 
		for _, v := range out {
			builder.WriteString(strconv.Itoa(v))
			builder.WriteString(" ")
		}
		ans := builder.String()[:builder.Len()-1] + "\n"
		fmt.Fprint(output, ans)
		builder.Reset()
	}
}
