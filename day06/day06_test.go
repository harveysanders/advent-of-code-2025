package day06

import (
	"embed"
	"io"
	"strings"
	"testing"

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
			desc: "example - part 1",
			getInput: func() (io.Reader, error) {
				in := strings.NewReader(`
123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  
`[1:])
				return in, nil
			},
			wantTotal: 4277556,
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
			wantTotal: 5667835681547,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {

			input, err := tc.getInput()
			require.NoError(t, err)

			worksheet, err := ParseWorksheet(input)
			require.NoError(t, err)

			total := worksheet.Do()
			require.Equal(t, tc.wantTotal, total)
		})
	}
}

func TestPart2(t *testing.T) {
	testCases := []struct {
		desc      string
		getInput  func() (io.Reader, error)
		wantTotal int
	}{
		{
			desc: "example - part 2",
			getInput: func() (io.Reader, error) {
				in := strings.NewReader(`
123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  
`[1:])
				return in, nil
			},
			wantTotal: 3263827,
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
			wantTotal: 9434900032651,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {

			input, err := tc.getInput()
			require.NoError(t, err)

			worksheet, err := ParseWorksheet(input)
			require.NoError(t, err)

			total := worksheet.Do2()
			require.Equal(t, tc.wantTotal, total)
		})
	}
}
