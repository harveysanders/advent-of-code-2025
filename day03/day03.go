package day03

import (
	"bytes"
	"io"
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

func (b Bank) FindHighestJoltage() int {
	var max1stDigit int
	var max2ndDigit int
	var maxPos int

	// Skip the last digit because we need at least 2 batteries
	for i := 0; i < len(b.Batteries)-1; i++ {
		curJ := b.Batteries[i]
		if curJ > max1stDigit {
			max1stDigit = curJ
			maxPos = i
		}
	}
	// slice the batteries list from the max1stDigit to end
	for _, v := range b.Batteries[maxPos+1:] {
		if v > max2ndDigit {
			max2ndDigit = v
		}
	}

	return (max1stDigit * 10) + max2ndDigit
}

func (b Banks) Total() int {
	var total int
	for _, bank := range b {
		total += bank.FindHighestJoltage()
	}
	return total
}
