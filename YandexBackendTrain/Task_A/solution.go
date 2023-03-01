// Golang 1.20.1	
// OK  106 ms  2.68 Mb

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var amount_tanks, last, maximum int
	var err error
	input := bufio.NewReader(os.Stdin)

	fmt.Fscan(input, &amount_tanks)
	tanks := make([]int, amount_tanks, amount_tanks)

	for i := 0; i < amount_tanks; i++ {
		fmt.Fscan(input, &tanks[i])
		switch {
		case err != nil:
			continue
		case tanks[i] > maximum:
			maximum = tanks[i]
		case tanks[i] < maximum || tanks[i] < last:
			err = errors.New("-1/n")
		default:
			last = tanks[i]
		}
	}
	if err != nil {
		fmt.Println("-1")
	} else {
		fmt.Println(strconv.Itoa(maximum - tanks[0]))
	}
}
