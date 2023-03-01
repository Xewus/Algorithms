// B Golang 1.20.1
// ID 83185436
// OK  4 ms  776.00 Kb

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type SeatsMap struct {
	Rows [][7]rune
}

func (sm *SeatsMap) outSeats(idx int, line byte, builder *strings.Builder) string {
	var temp [7]byte
	right := 0
	builder.Reset()
	builder.WriteString("Passengers can take seats:")
	seats := fmt.Sprintf("%07b", line)
	for i, v := range seats {
		if i == 3 {
			temp[i] = '_'
			right++
			continue
		}
		if v == '1' && sm.Rows[idx][i] == '.' {
			sm.Rows[idx][i] = '#'
			temp[i] = 'X'
			builder.WriteString(fmt.Sprintf(" %d%c", idx+1, i+'A'-right))
		} else {
			temp[i] = byte(sm.Rows[idx][i])
		}

	}
	builder.WriteString("\n")

	for i, row := range sm.Rows {
		if i == idx {
			builder.WriteString(string(temp[:]))
		} else {
			builder.WriteString(string(row[:]))
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

func inputSeats(input *bufio.Reader) (*[]byte, *SeatsMap) {
	var amountRows int
	var line string
	var row byte

	fmt.Fscan(input, &amountRows)
	plane := make([]byte, amountRows, amountRows)
	seatsMap := SeatsMap{Rows: make([][7]rune, amountRows, amountRows)}
	for i := 0; i < amountRows; i++ {
		row = 0
		fmt.Fscan(input, &line)
		for j, v := range line {
			seatsMap.Rows[i][j] = v
			row <<= 1
			if v == '#' {
				row++
			}
		}
		plane[i] = row
	}
	return &plane, &seatsMap
}

func main() {
	var amountGroups, amountPassengers uint8
	var side, window string
	errMsg := "Cannot fulfill passengers requirements\n"
	leftToWindow := map[uint8]byte{
		1: 64,
		2: 96,
		3: 112,
	}
	leftNotWindow := map[uint8]byte{
		1: 16,
		2: 48,
		3: 112,
	}
	rightToWindow := map[uint8]byte{
		1: 1,
		2: 3,
		3: 7,
	}
	rightNotWindow := map[uint8]byte{
		1: 4,
		2: 6,
		3: 7,
	}
	var group byte
	builder := strings.Builder{}
	input := bufio.NewReader(os.Stdin)
	p, sm := inputSeats(input)
	plane, seatsMap := *p, *sm

	fmt.Fscan(input, &amountGroups)

	for ; amountGroups > 0; amountGroups-- {
		fmt.Fscan(input, &amountPassengers, &side, &window)
		for i, row := range plane {
			switch {
			case side == "left":
				if window == "window" {
					group = leftToWindow[amountPassengers]
				} else {
					group = leftNotWindow[amountPassengers]
				}
			default:
				if window == "window" {
					group = rightToWindow[amountPassengers]
				} else {
					group = rightNotWindow[amountPassengers]

				}
			}
			if row&group == 0 {
				side = "ok"
				plane[i] = row | group
				fmt.Print(seatsMap.outSeats(i, plane[i], &builder))
				break
			}
		}
		if side != "ok" {
			fmt.Print(errMsg)
		}
	}
}
