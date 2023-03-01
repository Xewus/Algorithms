// (B) Морской бой (10 баллов)
// Полное решение: 10 баллов

package main
 
import (
	"fmt"
)
 
func main() {
	var a uint16
	fmt.Scanln(&a)
	for ; a > 0; a-- {
		fmt.Println(checker())
	}
}
 
func checker() string {
	var ship uint8
	ships := make([]uint8, 5, 5)
	for i := 10; i > 0; i-- {
		fmt.Scan(&ship)
		ships[ship]++
	}
 
	switch {
	case ships[1] != 4:
		return "NO"
	case ships[2] != 3:
		return "NO"
	case ships[3] != 2:
		return "NO"
	case ships[4] != 1:
		return "NO"
	default:
		return "YES"
	}
}
