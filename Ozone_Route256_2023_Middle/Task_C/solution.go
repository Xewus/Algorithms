// (C) Автомобильные номера (15 баллов)
// Полное решение: 15 баллов

package main
 
import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)
 
func main() {
	var amountTests uint16
	var numbers string
	f := os.Stdin
	input := bufio.NewReader(f)
 
	fmt.Fscan(input, &amountTests)
 
	for ; amountTests > 0; amountTests-- {
		fmt.Fscan(input, &numbers)
		fmt.Println(isNumbers(numbers))
	}
 
}
 
func isNumbers(numbers string) string {
	builder := strings.Builder{}
	for len(numbers) > 0 {
		matched, _ := regexp.MatchString(`^[A-Z]\d[A-Z]{2}`, numbers)
		if matched {
			builder.WriteString(numbers[:4])
			builder.WriteString(" ")
			numbers = numbers[4:]
			continue
		}
		matched, _ = regexp.MatchString(`^[A-Z]\d{2}[A-Z]{2}`, numbers)
		if matched {
			builder.WriteString(numbers[:5])
			builder.WriteString(" ")
			numbers = numbers[5:]
			continue
		}
		return "-"
	}
	return builder.String()
}
