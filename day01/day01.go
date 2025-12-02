package day01

import (
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
)

type direction string

const (
	dirLeft  direction = "L"
	dirRight direction = "R"
)

type rotation struct {
	// number of clicks to rotate the dial in direction
	n int
	// the dial is rotated clockwise, so the dial's position is decremented for a left rotation,
	// incremented for right.
	dir direction
}

type Dial struct {
	// series of rotation instructions
	Sequence []rotation
	// number from 0-99 indicating where the dial currently pointing.
	Position int
}

func (d *Dial) Password() int {
	var nAt0 int
	for _, rot := range d.Sequence {
		var nextPos int
		switch rot.dir {
		case dirLeft:
			nextPos = d.Position - rot.n
			if nextPos < 0 {
				nextPos += 100
			}
		case dirRight:
			nextPos = d.Position + rot.n
		default:
			continue
		}

		d.Position = nextPos % 100
		if d.Position == 0 {
			nAt0 += 1
		}
	}
	return nAt0
}

func (d *Dial) PasswordV2() int {
	var nPasses0 int
	for _, rot := range d.Sequence {
		var nextPos int
		diff := math.Abs(float64(rot.n - d.Position))
		fullRotations := math.Floor(diff / 100)
		nPasses0 += int(fullRotations)
		switch rot.dir {
		case dirLeft:
			nextPos = d.Position - rot.n
			if nextPos < 0 {
				nextPos += 100
				// nPasses0 += 1
			}
		case dirRight:
			nextPos = d.Position + rot.n
			if nextPos > 99 {
				// nPasses0 += 1
			}
		default:
			continue
		}

		d.Position = nextPos % 100
		if d.Position == 0 {
			nPasses0 += 1
		}
	}
	return nPasses0
}

func (d *Dial) rotate(rot rotation) (nextPos int, n0Passes int) {
	switch rot.dir {
	case dirLeft:
		distToZero := d.Position
		if rot.n < distToZero {
			nextPos = d.Position - rot.n
			return nextPos, n0Passes
		}
		n0Passes += 1
		// (rot.n - distToZero)
	case dirRight:
		distToZero := 100 - d.Position
		if rot.n < distToZero {
			nextPos = d.Position + rot.n
			return nextPos, n0Passes
		}
		n0Passes += 1
	default:
	}
	return 0, 0
}

func (d *Dial) ReadFrom(in io.Reader) (int64, error) {
	if d == nil {
		d = &Dial{}
	}
	data, err := io.ReadAll(in)
	if err != nil {
		return int64(len(data)), err
	}
	lines := strings.Split(string(data), "\n")
	d.Sequence = make([]rotation, 0, len(lines))

	for _, line := range lines {
		if line == "" {
			continue
		}
		rotation := rotation{}
		err := parseRotation(&rotation, line)
		if err != nil {
			return int64(len(data)), (err)
		}
		d.Sequence = append(d.Sequence, rotation)
	}
	return 0, nil
}

func parseRotation(r *rotation, line string) error {
	if len(line) < 2 {
		return fmt.Errorf("expected line to be at least 2 characters, got: %q", line)
	}
	r.dir = direction(line[0])
	n, err := strconv.Atoi(line[1:])
	if err != nil {
		return err
	}
	r.n = n
	return nil
}
