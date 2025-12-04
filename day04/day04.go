package day04

import (
	"bytes"
	"context"
	"io"
	"strings"
)

type Grid [][]string

type Loc struct {
	X int
	Y int
}

func (g Grid) leftOf(c Loc) string {
	if c.X == 0 {
		return ""
	}
	return g[c.Y][c.X-1]
}

func (g Grid) righttOf(c Loc) string {
	if c.X == len(g[c.Y])-1 {
		return ""
	}
	return g[c.Y][c.X+1]
}

func (g Grid) above(c Loc) string {
	if c.Y == 0 {
		return ""
	}
	return g[c.Y-1][c.X]
}

func (g Grid) below(c Loc) string {
	if c.Y == len(g)-1 {
		return ""
	}
	return g[c.Y+1][c.X]
}

func (g Grid) topLeft(c Loc) string {
	if c.X == 0 {
		return ""
	}
	if c.Y == 0 {
		return ""
	}
	return g[c.Y-1][c.X-1]
}

func (g Grid) topRight(c Loc) string {
	if c.Y == 0 {
		return ""
	}
	if c.X == len(g[c.Y])-1 {
		return ""
	}
	return g[c.Y-1][c.X+1]
}

func (g Grid) bottomLeft(c Loc) string {
	if c.Y == len(g)-1 {
		return ""
	}
	if c.X == 0 {
		return ""
	}
	return g[c.Y+1][c.X-1]
}

func (g Grid) bottomRight(c Loc) string {
	if c.Y == len(g)-1 {
		return ""
	}
	if c.X == len(g[c.Y])-1 {
		return ""
	}
	return g[c.Y+1][c.X+1]
}

func ParseGrid(in io.Reader) (Grid, error) {
	data, err := io.ReadAll(in)
	if err != nil {
		return nil, err
	}
	data = bytes.TrimSuffix(data, []byte("\n"))
	lines := strings.Split(string(data), "\n")
	g := make(Grid, 0, len(lines))

	for _, line := range lines {
		g = append(g, strings.Split(line, ""))
	}
	return g, nil
}

func (g Grid) CountAccessibleRolls() int {
	var total int
	for y, row := range g {
		for x, c := range row {
			if c != "@" {
				continue
			}
			if g.IsAccessible(Loc{X: x, Y: y}) {
				total += 1
			}
		}
	}
	return total
}

func (g Grid) IsAccessible(c Loc) bool {
	var adjacentRolls int
	var checkers = []func(Loc) string{
		g.topLeft,
		g.above,
		g.topRight,

		g.leftOf,
		g.righttOf,

		g.bottomLeft,
		g.below,
		g.bottomRight,
	}

	for _, fn := range checkers {
		if fn(c) == "@" {
			adjacentRolls += 1
		}
		if adjacentRolls >= 4 {
			return false
		}
	}
	return true
}

func (g Grid) AccessMap() (string, error) {
	var diagaram strings.Builder
	for y, row := range g {
		for x, c := range row {
			if c != "@" {
				_, err := diagaram.WriteString(c)
				if err != nil {
					return "", err
				}
				continue
			}
			// is a roll
			char := "@"
			if g.IsAccessible(Loc{X: x, Y: y}) {
				char = "x"
			}
			_, err := diagaram.WriteString(char)
			if err != nil {
				return "", err
			}
		}
		_, err := diagaram.WriteString("\n")
		if err != nil {
			return "", err
		}
	}
	return diagaram.String(), nil
}

func (g Grid) RemoveRolls(ctx context.Context) (int, error) {
	var passes int
	removedLastPass := -1
	var totalRemoved int

	for removedLastPass != 0 {
		if ctx.Err() != nil {
			return totalRemoved, ctx.Err()
		}

		removedLastPass = 0
		for y, row := range g {
			for x, c := range row {
				// remove rolls from last pass
				if c == "x" {
					g[y][x] = "."
				}
				if c != "@" {
					continue
				}
				if g.IsAccessible(Loc{X: x, Y: y}) {
					// mark for removal
					g[y][x] = "x"
					removedLastPass += 1
					totalRemoved += 1
					continue
				}
			}
		}
		passes += 1
		// fmt.Println(g.String())
	}
	return totalRemoved, nil
}

func (g Grid) String() string {
	var out strings.Builder
	for _, row := range g {
		_, _ = out.WriteString(strings.Join(row, ""))
		_, _ = out.WriteString("\n")
	}
	return out.String()
}
