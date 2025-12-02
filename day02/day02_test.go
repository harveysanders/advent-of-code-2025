package day02_test

import (
	"embed"
	"io"
	"strings"
	"testing"

	"github.com/harveysanders/aoc2025/day02"
	"github.com/stretchr/testify/require"
)

//go:embed input/*.txt
var inputFS embed.FS

func TestPart1(t *testing.T) {
	testCases := []struct {
		desc     string
		getInput func() (io.Reader, error)
		wantSum  int
	}{
		{
			desc: "tiny example",
			getInput: func() (io.Reader, error) {
				in := strings.NewReader(`
95-115
`[1:])
				return in, nil
			},
			wantSum: 99,
		},
		{
			desc: "tiny example: 998-1012",
			getInput: func() (io.Reader, error) {
				in := strings.NewReader(`
998-1012
`[1:])
				return in, nil
			},
			wantSum: 1010,
		},
		{
			desc: "example",
			getInput: func() (io.Reader, error) {
				in := strings.NewReader(`
11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,46443-446449,38593856-38593862,565653-565659,824824821-824824827,121212118-2121212124
`[1:])
				return in, nil
			},
			wantSum: 1227775554,
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
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			input, err := tc.getInput()
			require.NoError(t, err)

			idRanges, err := day02.ParseIDS(input)
			require.NoError(t, err)

			err = day02.FindAllInvalidIDs(idRanges)
			require.NoError(t, err)
			gotSum := day02.Part1Sum(idRanges)
			require.Equal(t, tc.wantSum, gotSum)
		})
	}
}
