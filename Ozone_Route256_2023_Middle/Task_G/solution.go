// (G) Записи ко врачу (25 баллов)
// Полное решение: 25 баллов

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Record struct {
	index  int
	window int
	moved  rune
}

func main() {
	var amountTests uint16
	var amountWindows, amountRecords int
	f := os.Stdin
	// f, err := os.Open("19")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	input := bufio.NewReader(f)

	fmt.Fscanln(input, &amountTests)

	for ; amountTests > 0; amountTests-- {
		fmt.Fscan(input, &amountWindows, &amountRecords)
		windows := make([]bool, amountWindows, amountWindows)
		records := make([]Record, amountRecords, amountRecords)

		for i := range records {
			records[i] = Record{index: i, moved: '0'}
			fmt.Fscan(input, &records[i].window)
			records[i].window--
		}

		if betterAssist(&records, &windows) != nil {
			fmt.Println("x")
		} else {
			outRecords(&records)
		}
	}
}

func betterAssist(pRecords *[]Record, pWindows *[]bool) error {
	records := *pRecords
	windows := *pWindows

	sort.Slice(records, func(i, j int) bool { return records[i].window < records[j].window })

	for i := 0; i < len(records); i++ {
		index := records[i].window
		switch {
		case index > 0 && !windows[index-1]:
			windows[index-1] = true
			records[i].moved = '-'
		case !windows[index]:
			windows[index] = true
		case index < len(windows)-1 && !windows[index+1]:
			windows[index+1] = true
			records[i].moved = '+'
		default:
			return errors.New("x")
		}
	}
	sort.Slice(records, func(i, j int) bool { return records[i].index < records[j].index })
	return nil
}

func outRecords(pRecords *[]Record) {
	builder := strings.Builder{}

	for _, record := range *pRecords {
		builder.WriteByte(byte(record.moved))
	}
	fmt.Println(builder.String())
}
