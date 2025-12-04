package day06

import (
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
	"sync"
)

type operator string

const (
	OpAdd      operator = "+"
	OpMultiply operator = "*"
)

type Problem struct {
	op          operator
	rawOperands chan string
	// The following fields are only used for vertial numDir.

	// For vertically aligned numbers, the offset from the start
	// of the row to the first column of the problem's
	// column group.
	// Ex:
	//   123 328  51 64
	//    45 64  387 23
	//     6 98  215 314
	//   *   +   *   +
	// in the worksheet above, the leftmost problem's offset is 0
	// determined by the operator position.
	// The next problem, from left to right, has an offset of 4
	offset int
	// Number of colums the problem uses.
	colWidth int
	rawRows  []string
}

type Worksheet struct {
	probs   []Problem
	rawRows []string
}

func ParseWorksheet(r io.Reader) (Worksheet, error) {
	raw, err := io.ReadAll(r)
	if err != nil {
		return Worksheet{}, err
	}
	raw = bytes.TrimSuffix(raw, []byte("\n"))
	lines := strings.Split(string(raw), "\n")

	w := Worksheet{
		probs:   make([]Problem, 0, 200),
		rawRows: lines[:len(lines)-1],
	}

	operatorLine := lines[len(lines)-1]
	var prob Problem
	for col, v := range strings.Split(operatorLine, "") {
		switch v {
		case string(OpAdd), string(OpMultiply):
			// stop parsing the prev problem
			// and add to list
			if col > 0 {
				// decrement the col width for
				// the separator column
				prob.colWidth -= 1
				w.probs = append(w.probs, prob)
			}
			// --------------------
			// start a new problem
			prob = Problem{
				rawOperands: make(chan string, 1),
				offset:      col,
				rawRows:     w.rawRows,
				colWidth:    1,
			}
			switch operator(v) {
			case OpAdd:
				prob.op = OpAdd
			case OpMultiply:
				prob.op = OpMultiply
			}
			continue
		case " ":
			prob.colWidth += 1
		}
	}
	// add the last problem
	w.probs = append(w.probs, prob)

	return w, nil
}

func (w Worksheet) Do() int {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var total int

	for _, prob := range w.probs {
		wg.Go(func() {
			res := prob.calc()
			mu.Lock()
			total += res
			mu.Unlock()
		})
	}

	for _, row := range w.rawRows {
		rawNums := strings.Fields(row)
		for i, v := range rawNums {
			w.probs[i].rawOperands <- v
		}
	}
	for _, p := range w.probs {
		close(p.rawOperands)
	}
	wg.Wait()
	return total
}

func (p Problem) calc() int {
	var total int
	if p.op == OpMultiply {
		// so we don't multiply by 0
		total = 1
	}
	for rawNum := range p.rawOperands {
		v, err := strconv.Atoi(rawNum)
		if err != nil {
			// TODO: Handle this error
			panic(err)
		}
		switch p.op {
		case OpAdd:
			total += v
		case OpMultiply:
			total *= v
		default:
			panic(fmt.Sprintf("invalid operator: %q", string(p.op)))
		}
	}
	return total
}

func (w Worksheet) Do2() int {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var total int

	for _, prob := range w.probs {
		wg.Go(func() {
			res, err := prob.calc2()
			if err != nil {
				panic(err)
			}
			mu.Lock()
			total += res
			mu.Unlock()
		})
	}

	wg.Wait()
	return total
}

func (p Problem) calc2() (int, error) {
	var total int
	if p.op == OpMultiply {
		// so we don't multiply by 0
		total = 1
	}
	operands := make([]int, 0, len(p.rawRows))
	var rawNum strings.Builder
	for x := p.offset; x < p.colWidth+p.offset; x++ {
		for y := range p.rawRows {
			char := p.rawRows[y][x]
			if char == ' ' {
				continue
			}
			err := rawNum.WriteByte(char)
			if err != nil {
				return 0, err
			}
		}
		str := rawNum.String()
		num, err := strconv.Atoi(str)
		if err != nil {
			return 0, err
		}
		operands = append(operands, num)
		rawNum.Reset()
	}

	for _, v := range operands {
		switch p.op {
		case OpAdd:
			total += v
		case OpMultiply:
			total *= v
		}
	}
	return total, nil
}
