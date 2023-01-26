// (I) Планировщик задач (30 баллов)
// Частичное решение: 20 баллов

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var GlobalTime int
var Work int

type Procces struct {
	Power int
	End   int
}

type Task struct {
	Start int
	Dur   int
}

type FreeQueue struct {
	Proc *Procces
	Next *FreeQueue
	Tail *FreeQueue
}

func (fq *FreeQueue) add(pr *Procces, parent *FreeQueue) {
	switch {
	case fq.Proc == nil:
		fq.Proc = pr
	case fq.Proc.Power >= pr.Power:
		next := FreeQueue{Proc: fq.Proc, Next: fq.Next}
		fq.Proc = pr
		fq.Next = &next
		if parent != nil {
			parent.Next = fq
		}
	case fq.Next == nil:
		fq.Next = &FreeQueue{Proc: pr}
	default:
		fq.Next.add(pr, fq)
	}
}

// fill queue from sorted slice
func (fq *FreeQueue) push(pr *Procces) {
	if fq.Proc == nil {
		fq.Proc = pr
		fq.Tail = fq
	} else {
		tail := FreeQueue{Proc: pr}
		fq.Tail.Next = &tail
		fq.Tail = &tail
	}
}

// get first element from queue
func (q *FreeQueue) get() (error, *Procces) {
	if q.Proc == nil {
		err := errors.New("empty queue")
		return err, nil
	}
	p := q.Proc
	if q.Next == nil {
		q.Proc = nil
	} else {
		q.Proc = q.Next.Proc
		q.Next = q.Next.Next
	}
	return nil, p
}

type WorkQueue struct {
	Proc *Procces
	Next *WorkQueue
}

func (wq *WorkQueue) add(pr *Procces, parent *WorkQueue) {
	switch {
	case wq.Proc == nil:
		wq.Proc = pr
	case wq.Proc.End >= pr.End:
		next := WorkQueue{Proc: wq.Proc, Next: wq.Next}
		wq.Proc = pr
		wq.Next = &next
		if parent != nil {
			parent.Next = wq
		}
	case wq.Next == nil:
		wq.Next = &WorkQueue{Proc: pr}
	default:
		wq.Next.add(pr, wq)
	}
}

func (wq *WorkQueue) clean(fq *FreeQueue) {
	if wq.Proc == nil || wq.Proc.End > GlobalTime {
		return
	}
	fq.add(wq.Proc, nil)
	if wq.Next == nil {
		wq.Proc = nil
		return
	}
	wq.Proc = wq.Next.Proc
	wq.Next = wq.Next.Next
	wq.clean(fq)
}

func (wq *WorkQueue) addTask(t *Task, fq *FreeQueue) {
	wq.clean(fq)
	err, pr := fq.get()
	if err != nil {
		return
	}
	pr.End = GlobalTime + t.Dur
	Work += pr.Power * t.Dur
	wq.add(pr, nil)
}

func main() {
	var fq FreeQueue
	var wq WorkQueue

	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanWords(data, atEOF)
		if err == nil && token != nil {
			_, err = strconv.ParseInt(string(token), 10, 32)
		}
		return
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(split)

	scanner.Scan()
	amountProcess, _ := strconv.Atoi(scanner.Text())
	ps := make([]*Procces, amountProcess, amountProcess)

	scanner.Scan()
	amountTasks, _ := strconv.Atoi(scanner.Text())

	for i := range ps {
		scanner.Scan()
		power, _ := strconv.Atoi(scanner.Text())
		ps[i] = &Procces{Power: power}
	}

	sort.Slice(ps, func(i, j int) bool { return ps[i].Power < ps[j].Power })

	for _, p := range ps {
		fq.push(p)
	}

	for ; amountTasks > 0; amountTasks-- {
		task := Task{}
		scanner.Scan()
		task.Start, _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		task.Dur, _ = strconv.Atoi(scanner.Text())
		GlobalTime = task.Start
		wq.addTask(&task, &fq)
	}
	fmt.Println(Work)
}
