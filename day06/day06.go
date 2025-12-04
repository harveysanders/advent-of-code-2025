package day06

import (
	"bytes"
	"fmt"
	"io"
	"slices"
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

	operators := strings.Fields(lines[len(lines)-1])

	w := Worksheet{
		probs:   make([]Problem, 0, len(operators)),
		rawRows: lines[:len(lines)-1],
	}
	validOps := []operator{OpAdd, OpMultiply}
	for _, op := range operators {
		if !slices.Contains(validOps, operator(op)) {
			return w, fmt.Errorf("invalid operator: %q", op)
		}
		prob := Problem{
			op:          operator(op),
			rawOperands: make(chan string, 1),
		}

		w.probs = append(w.probs, prob)
	}

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
