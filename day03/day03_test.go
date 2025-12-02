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

func TestPart1(t *testing.T) {
	testCases := []struct {
		desc      string
		getInput  func() (io.Reader, error)
		wantTotal int
	}{

		{
			desc: "example",
			getInput: func() (io.Reader, error) {
				in := strings.NewReader(`
987654321111111
811111111111119
234234234234278
818181911112111
`[1:])
				return in, nil
			},
			wantTotal: 357,
		},
		{
			desc: "real input",
			getInput: func() (io.Reader, error) {
				f, err := inputFS.Open("input/input.txt")
				if err != nil {
					return nil, err
				}
				return f, nil
			},
			wantTotal: 17095,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			input, err := tc.getInput()
			require.NoError(t, err)

			banks, err := day03.ParseBanks(input)
			require.NoError(t, err)

			gotTotal := banks.Total()
			require.Equal(t, tc.wantTotal, gotTotal)
		})
	}
}

func TestFindHighestJoltage(t *testing.T) {
	testCases := []struct {
		desc        string
		bank        day03.Bank
		wantJoltage int
	}{
		{
			desc: "987654321111111",
			bank: day03.Bank{
				Batteries: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 1, 1, 1, 1, 1},
			},
			wantJoltage: 98,
		},
		{
			desc: "811111111111119",
			bank: day03.Bank{
				Batteries: []int{8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 9},
			},
			wantJoltage: 89,
		},
		{
			desc: "234234234234278",
			bank: day03.Bank{
				Batteries: []int{2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 3, 4, 2, 7, 8},
			},
			wantJoltage: 78,
		},
		{
			desc: "818181911112111",
			bank: day03.Bank{
				Batteries: []int{8, 1, 8, 1, 8, 1, 9, 1, 1, 1, 1, 2, 1, 1, 1},
			},
			wantJoltage: 92,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			got := tc.bank.FindHighestJoltage()
			require.Equal(t, tc.wantJoltage, got)
		})
	}
}
