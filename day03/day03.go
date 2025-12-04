package day03

import (
	"bytes"
	"io"
	"math"
	"strconv"
	"strings"
)

type Bank struct {
	Batteries []int
}

type Banks []Bank

func ParseBanks(in io.Reader) (Banks, error) {
	raw, err := io.ReadAll(in)
	if err != nil {
		return nil, err
	}
	raw = bytes.TrimSuffix(raw, []byte("\n"))
	lines := strings.Split(string(raw), "\n")
	banks := make(Banks, 0, len(lines))

	for _, line := range lines {
		b := Bank{Batteries: make([]int, 0, len(line))}
		for j := range strings.SplitSeq(line, "") {
			joltage, err := strconv.Atoi(j)
			if err != nil {
				return nil, err
			}
			b.Batteries = append(b.Batteries, joltage)
		}
		banks = append(banks, b)
	}

	return banks, nil
}

func (b Bank) FindHighestJoltage(n int) int {
	var recurse func([]int, []int, int) []int
	recurse = func(restBatts []int, joltages []int, nLeft int) []int {
		if nLeft == 0 {
			return joltages
		}
		var maxPos int
		var max int
		for i, v := range restBatts {
			// need at least nLeft in the slice we pass to the next recurse call
			if nLeft <= len(restBatts)-i {
				if v > max {
					max = v
					maxPos = i
				}
			}
		}
		next := append(joltages, max)
		return recurse(restBatts[maxPos+1:], next, nLeft-1)
	}

	joltages := make([]int, 0, n)
	joltages = recurse(b.Batteries, joltages, n)
	var total float64
	for i, v := range joltages {
		place := math.Pow(10, float64(n-i-1))
		total += float64(v) * place
	}

	return int(total)
}

func (b Banks) Total(n int) int {
	var total int
	for _, bank := range b {
		total += bank.FindHighestJoltage(n)
	}
	return total
}
