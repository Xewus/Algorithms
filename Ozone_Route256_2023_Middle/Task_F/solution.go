// (F) Печать документа (20 баллов)
// Полное решение: 20 баллов

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var amountTests, amountPages uint8
	var printed string
	input := bufio.NewReader(os.Stdin)

	fmt.Fscanln(input, &amountTests)

	for ; amountTests > 0; amountTests-- {
		fmt.Fscanln(input, &amountPages)
		fmt.Fscanln(input, &printed)
		printed := isPrinted(amountPages, printed)
		fmt.Println(getToPrint(printed))
	}

}

func isPrinted(amountPages uint8, data string) *[]bool {
	var idx, end int
	pages := make([]bool, amountPages, amountPages)
	printed := strings.Split(data, ",")

	for _, print := range printed {
		pointers := strings.Split(print, "-")
		idx, _ = strconv.Atoi(pointers[0])
		idx--
		if len(pointers) == 1 {
			pages[idx] = true
		} else {
			end, _ = strconv.Atoi(pointers[1])
			for ; idx < end; idx++ {
				pages[idx] = true
			}
		}
	}
	return &pages
}

func getToPrint(pPages *[]bool) string {
	var many bool
	builder := strings.Builder{}
	pages := *pPages
	if !pages[0] {
		builder.WriteString("1")
	}
	for i := 1; i < len(pages); i++ {
		switch {
		case pages[i] && !pages[i-1]:
			if many {
				many = false
				builder.WriteString(strconv.Itoa(i))
			}
			builder.WriteString(",")
		case !pages[i]:
			if pages[i-1] {
				builder.WriteString(strconv.Itoa(i + 1))
			} else if !many {
				many = true
				builder.WriteString("-")
			}
		}
	}
	if many {
		builder.WriteString(strconv.Itoa(len(pages)))
	}
	ans := builder.String()
	if ans[len(ans)-1] == ',' {
		ans = ans[:len(ans)-1]
	}
	return ans
}
