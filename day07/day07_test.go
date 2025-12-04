package day07

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
.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............
`[1:])
				return in, nil
			},
			wantTotal: 21,
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
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {

			input, err := tc.getInput()
			require.NoError(t, err)

			diagram, err := ParseDiagram(input)
			require.NoError(t, err)

			total, _ := diagram.CountBeams()
			require.NoError(t, err)

			require.Equal(t, tc.wantTotal, total)
		})
	}
}

func TestDiagramString(t *testing.T) {
	testCases := []struct {
		desc        string
		getInput    func() (io.Reader, error)
		wantDiagram string
	}{
		{
			desc: "example - part 1",
			getInput: func() (io.Reader, error) {
				in := strings.NewReader(`
.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............
`[1:])
				return in, nil
			},
			wantDiagram: `
.......S.......
.......|.......
......|^|......
......|.|......
.....|^|^|.....
.....|.|.|.....
....|^|^|^|....
....|.|.|.|....
...|^|^|||^|...
...|.|.|||.|...
..|^|^|||^|^|..
..|.|.|||.|.|..
.|^|||^||.||^|.
.|.|||.||.||.|.
|^|^|^|^|^|||^|
|.|.|.|.|.|||.|`[1:],
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			input, err := tc.getInput()
			require.NoError(t, err)

			diagram, err := ParseDiagram(input)
			require.NoError(t, err)

			_, gotDiagram := diagram.CountBeams()
			require.NoError(t, err)

			require.Equal(t, tc.wantDiagram, gotDiagram)
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
.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............
`[1:])
				return in, nil
			},
			wantTotal: 40,
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
		},
	}
	for _, tc := range testCases[1:] {
		t.Run(tc.desc, func(t *testing.T) {

			input, err := tc.getInput()
			require.NoError(t, err)

			diagram, err := ParseDiagram(input)
			require.NoError(t, err)

			total := diagram.CountTimelines()
			require.NoError(t, err)

			require.Equal(t, tc.wantTotal, total)
		})
	}
}
