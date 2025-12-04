package day07

import (
	"bytes"
	"io"
	"strings"
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
func (d Diagram) CountBeams() (int, string, error) {
	var nSplits int
	if len(d.grid) == 0 {
		return 0, "", nil
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
	return nSplits, d.String(), nil
}

func (d Diagram) String() string {
	return strings.Join(d.grid, "\n")
}
