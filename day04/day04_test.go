package day04_test

import (
	"context"
	"embed"
	"io"
	"strings"
	"testing"
	"time"

	"github.com/harveysanders/aoc2025/day04"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//go:embed input/*.txt
var inputFS embed.FS

func TestRolls(t *testing.T) {
	testCases := []struct {
		desc      string
		getInput  func() (io.Reader, error)
		wantRolls int
	}{
		{
			desc: "example - part 1",
			getInput: func() (io.Reader, error) {
				in := strings.NewReader(`
..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.
`[1:])
				return in, nil
			},
			wantRolls: 13,
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
			wantRolls: 13,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			input, err := tc.getInput()
			require.NoError(t, err)

			grid, err := day04.ParseGrid(input)
			require.NoError(t, err)

			gotRolls := grid.CountAccessibleRolls()
			require.Equal(t, tc.wantRolls, gotRolls)
		})
	}
}

func TestIsAccessible(t *testing.T) {
	testCases := []struct {
		desc             string
		cellLoc          day04.Loc
		wantIsAccessible bool
	}{
		{
			desc:             "top row, 3rd from left",
			cellLoc:          day04.Loc{X: 2, Y: 0},
			wantIsAccessible: true,
		},
		{
			desc:             "left edge, 5th row",
			cellLoc:          day04.Loc{X: 0, Y: 4},
			wantIsAccessible: true,
		},
		{
			desc:             "bottom row",
			cellLoc:          day04.Loc{X: 2, Y: 9},
			wantIsAccessible: true,
		},
	}

	in := strings.NewReader(`
..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.
`[1:])

	grid, err := day04.ParseGrid(in)
	require.NoError(t, err)

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			val := grid[tc.cellLoc.Y][tc.cellLoc.X]
			require.Equal(t, "@", val, "expected cell to be a roll: '@")

			gotIsAccessible := grid.IsAccessible(tc.cellLoc)
			require.Equal(t, tc.wantIsAccessible, gotIsAccessible)
		})
	}

}

func TestAccessMap(t *testing.T) {
	testCases := []struct {
		desc     string
		getGrid  func(t *testing.T) day04.Grid
		wantGrid string
	}{
		{
			getGrid: func(t *testing.T) day04.Grid {
				in := strings.NewReader(`
..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.
`[1:])
				grid, err := day04.ParseGrid(in)
				require.NoError(t, err)
				return grid
			},
			wantGrid: `
..xx.xx@x.
x@@.@.@.@@
@@@@@.x.@@
@.@@@@..@.
x@.@@@@.@x
.@@@@@@@.@
.@.@.@.@@@
x.@@@.@@@@
.@@@@@@@@.
x.x.@@@.x.
`[1:],
		},
	}

	for _, tc := range testCases {
		grid := tc.getGrid(t)
		diagram, err := grid.AccessMap()
		require.NoError(t, err)

		require.Equal(t, tc.wantGrid, diagram)
	}
}

func TestRemoveRolls(t *testing.T) {
	testCases := []struct {
		desc        string
		getInput    func() (io.Reader, error)
		wantRemoved int
	}{
		{
			desc: "example - part 2",
			getInput: func() (io.Reader, error) {
				in := strings.NewReader(`
..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.
`[1:])
				return in, nil
			},
			wantRemoved: 43,
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
			wantRemoved: 8557,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			input, err := tc.getInput()
			require.NoError(t, err)

			grid, err := day04.ParseGrid(input)
			require.NoError(t, err)

			ctx, cancel := context.WithTimeout(t.Context(), 10*time.Second)
			defer cancel()

			gotRolls, err := grid.RemoveRolls(ctx)
			assert.NoError(t, err)
			require.Equal(t, tc.wantRemoved, gotRolls)
		})
	}
}
