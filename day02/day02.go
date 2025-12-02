package day02

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
)

type IDRange struct {
	First      int
	Last       int
	invalidIDs []int
}

func ParseIDS(in io.Reader) ([]*IDRange, error) {
	data, err := io.ReadAll(in)
	if err != nil {
		return nil, err
	}
	data = bytes.TrimSuffix(data, []byte("\n"))
	ranges := strings.Split(string(data), ",")
	idRanges := make([]*IDRange, 0, len(ranges))
	for _, rawRange := range ranges {
		ids := strings.Split(rawRange, "-")
		if len(ids) != 2 {
			continue
		}
		first, err := strconv.Atoi(ids[0])
		if err != nil {
			return nil, fmt.Errorf("parse low: %w", err)
		}

		last, err := strconv.Atoi(ids[1])
		if err != nil {
			return nil, fmt.Errorf("parse low: %w", err)
		}

		idRange := &IDRange{
			First:      first,
			Last:       last,
			invalidIDs: make([]int, 0, 8),
		}
		idRanges = append(idRanges, idRange)
	}
	return idRanges, nil
}

func FindAllInvalidIDs(ranges []*IDRange) error {
	for _, v := range ranges {
		invalidIDs, err := findInvalidIDs(*v)
		if err != nil {
			return err
		}
		v.invalidIDs = invalidIDs
	}
	return nil
}

func findInvalidIDs(idr IDRange) ([]int, error) {
	res := make([]int, 0, 16)
	leastDigits := int(math.Ceil(math.Log10(float64(idr.First))))
	mostDigits := int(math.Ceil(math.Log10(float64(idr.Last))))

	var sb strings.Builder
	for nDigits := leastDigits; nDigits <= mostDigits; nDigits++ {
		if nDigits%2 != 0 {
			continue
		}
		for i := range 10 {
			sb.Reset()
			if i == 0 && sb.Len() < 2 {
				continue
			}
			for range nDigits {
				sb.WriteString(strconv.Itoa(i))
			}
			val, err := strconv.Atoi(sb.String())
			if err != nil {
				return nil, err
			}
			if idr.First <= val && val <= idr.Last {
				res = append(res, val)
			}
		}
	}
	// highestPow := math.Ceil(math.Log10(float64(idr.Last)))
	// get the first digit from first ID
	// create sequences of length up to half the amount of digits
	// check if each seqence is in range
	return res, nil
}

func Part1Sum(ranges []*IDRange) int {
	var sum int
	for _, r := range ranges {
		x := make([]string, 0, len(r.invalidIDs))
		for _, v := range r.invalidIDs {
			x = append(x, strconv.Itoa(v))
		}
		println(strings.Join(x, ", "))
		for _, v := range r.invalidIDs {
			sum += v
		}
	}
	return sum
}
