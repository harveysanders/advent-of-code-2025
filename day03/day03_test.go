package day03_test

import (
	"embed"
	"io"
	"strings"
	"testing"

	"github.com/harveysanders/aoc2025/day03"
	"github.com/stretchr/testify/require"
)

//go:embed input/*.txt
var inputFS embed.FS

func TestTotal(t *testing.T) {
	testCases := []struct {
		desc      string
		getInput  func() (io.Reader, error)
		nDigits   int
		wantTotal int
	}{
		{
			desc: "example - part 1",
			getInput: func() (io.Reader, error) {
				in := strings.NewReader(`
987654321111111
811111111111119
234234234234278
818181911112111
`[1:])
				return in, nil
			},
			nDigits:   2,
			wantTotal: 357,
		},
		{
			desc: "example - part 2",
			getInput: func() (io.Reader, error) {
				in := strings.NewReader(`
987654321111111
811111111111119
234234234234278
818181911112111
`[1:])
				return in, nil
			},
			nDigits:   12,
			wantTotal: 3121910778619,
		},
		{
			desc: "real input - part 1",
			getInput: func() (io.Reader, error) {
				f, err := inputFS.Open("input/input.txt")
				if err != nil {
					return nil, err
				}
				return f, nil
			},
			nDigits:   2,
			wantTotal: 17095,
		},
		{
			desc: "real input - part 2",
			getInput: func() (io.Reader, error) {
				f, err := inputFS.Open("input/input.txt")
				if err != nil {
					return nil, err
				}
				return f, nil
			},
			nDigits:   12,
			wantTotal: 17095,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			input, err := tc.getInput()
			require.NoError(t, err)

			banks, err := day03.ParseBanks(input)
			require.NoError(t, err)

			gotTotal := banks.Total(tc.nDigits)
			require.Equal(t, tc.wantTotal, gotTotal)
		})
	}
}

func TestFindHighestJoltage(t *testing.T) {
	testCases := []struct {
		desc        string
		bank        day03.Bank
		nDigits     int
		wantJoltage int
	}{
		{
			desc: "part 1: 987654321111111",
			bank: day03.Bank{
				Batteries: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1},
			},
			nDigits:     2,
			wantJoltage: 98,
		},
		{
			desc: "part 1: 811111111111119",
			bank: day03.Bank{
				Batteries: []int{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9},
			},
			nDigits:     2,
			wantJoltage: 89,
		},
		{
			desc: "part 1: 234234234234278",
			bank: day03.Bank{
				Batteries: []int{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8},
			},
			nDigits:     2,
			wantJoltage: 78,
		},
		{
			desc: "part 1: 818181911112111",
			bank: day03.Bank{
				Batteries: []int{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1},
			},
			nDigits:     2,
			wantJoltage: 92,
		},
		{
			desc: "part 2: 987654321111111",
			bank: day03.Bank{
				Batteries: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1},
			},
			nDigits:     12,
			wantJoltage: 987654321111,
		},
		{
			desc: "part 2: 811111111111119",
			bank: day03.Bank{
				Batteries: []int{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9},
			},
			nDigits:     12,
			wantJoltage: 811111111119,
		},
		{
			desc: "part 2: 234234234234278",
			bank: day03.Bank{
				Batteries: []int{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8},
			},
			nDigits:     12,
			wantJoltage: 434234234278,
		},
		{
			desc: "part 2: 818181911112111",
			bank: day03.Bank{
				Batteries: []int{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1},
			},
			nDigits:     12,
			wantJoltage: 888911112111,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			got := tc.bank.FindHighestJoltage(tc.nDigits)
			require.Equal(t, tc.wantJoltage, got)
		})
	}
}
