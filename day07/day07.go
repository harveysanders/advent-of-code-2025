package day07

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"sync"
	"sync/atomic"
)

type Diagram struct {
	grid []string
}

func ParseDiagram(r io.Reader) (*Diagram, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}
	data = bytes.TrimSuffix(data, []byte{'\n'})
	lines := strings.Split(string(data), "\n")
	d := Diagram{grid: lines}
	return &d, nil
}

func tachPositions(row string) []int {
	tacyhons := make([]int, 0, len(row))
	for i, char := range row {
		if char == 'S' || char == '|' {
			tacyhons = append(tacyhons, i)
		}
	}
	return tacyhons
}

func splitBeams(curRow string, beamPositions []int) (int, string) {
	var splits int
	newRow := make([]byte, len(curRow))
	_ = copy(newRow, curRow)
	for _, pos := range beamPositions {
		char := newRow[pos]
		switch char {
		case '|':
			continue
		case '.':
			newRow[pos] = '|'
		case '^':
			splits += 1
			left := pos - 1
			right := pos + 1
			if left >= 0 {
				newRow[left] = '|'
			}
			if right <= len(newRow)-1 {
				newRow[right] = '|'
			}
		}
	}
	return splits, string(newRow)
}

// CountBeams simulates tachyon beam propagation through the diagram, spliting the beam
// whenever it encounters a splitter "^".
// Returns the total number of splits, the final grid state as a string, and any error.
func (d Diagram) CountBeams() (int, string) {
	var nSplits int
	if len(d.grid) == 0 {
		return 0, ""
	}

	for y, row := range d.grid {
		if y == 0 {
			continue
		}
		// find cur tachs
		prevRow := d.grid[y-1]
		tacyhons := tachPositions(prevRow)
		splits, newRow := splitBeams(row, tacyhons)
		nSplits += splits
		d.grid[y] = newRow
	}
	return nSplits, d.String()
}

func (d Diagram) String() string {
	return strings.Join(d.grid, "\n")
}

func (d Diagram) CountTimelines() int {
	var wg sync.WaitGroup
	var traverse func(d Diagram, beamPosition int, y int, paths *uint64)

	traverse = func(d Diagram, x int, y int, paths *uint64) {
		if y == len(d.grid) {
			atomic.AddUint64(paths, 1)
			if *paths%1_000_000 == 0 {
				println(*paths)
			}
			return
		}
		row := d.grid[y]
		// newRow := make([]byte, len(row))
		// _ = copy(newRow, row)
		switch row[x] {
		case 'S':
			traverse(d, x, y+1, paths)
			return
		case '.', '|':
			// newRow[x] = '|'
			// d.grid[y] = string(newRow)
			traverse(d, x, y+1, paths)
			return
		case '^':
			// iterate 1 pos. left and right of x

			for i := -1; i < 2; i += 2 {
				nextX := x + i
				if nextX < 0 || nextX > len(row)-1 {
					continue
				}

				// newRow[nextX] = '|'
				// d.grid[y] = string(newRow)
				wg.Go(func() { traverse(d, nextX, y+1, paths) })
			}
		default:
			panic("we should have hit a '|' or '.' ")
		}
	}
	start := strings.IndexByte(d.grid[0], 'S')
	var timeslines uint64
	traverse(d, start, 0, &timeslines)
	wg.Wait()
	fmt.Printf("****** done! ->%d ******* \n", timeslines)
	return int(timeslines)
}
