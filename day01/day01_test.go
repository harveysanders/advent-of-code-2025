package day01_test

import (
	"embed"
	"io"
	"strings"
	"testing"

	"github.com/harveysanders/aoc2025/day01"
	"github.com/stretchr/testify/require"
)

//go:embed input/*.txt
var inputFS embed.FS

func TestDay01Part1(t *testing.T) {
	testCases := []struct {
		desc         string
		getInput     func() (io.Reader, error)
		wantPassword int
	}{
		{
			desc: "example",
			getInput: func() (io.Reader, error) {
				in := strings.NewReader(`
L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
`[1:])
				return in, nil
			},
			wantPassword: 3,
		},
		{
			desc: "real input",
			getInput: func() (io.Reader, error) {
				dir, err := inputFS.ReadDir("input")
				if err != nil {
					return nil, err
				}
				for _, e := range dir {
					println(e.Name())
				}
				f, err := inputFS.Open("input/input.txt")
				if err != nil {
					return nil, err
				}
				return f, nil
			},
			wantPassword: 1026,
		},
	}

	for _, tc := range testCases {
		input, err := tc.getInput()
		require.NoError(t, err)

		dial := day01.Dial{Position: 50}

		_, err = dial.ReadFrom(input)
		require.NoError(t, err)

		gotPW := dial.Password()
		require.Equal(t, tc.wantPassword, gotPW)
	}
}

func TestDay01Part2(t *testing.T) {
	testCases := []struct {
		desc         string
		getInput     func() (io.Reader, error)
		wantPassword int
	}{
		{
			desc: "example",
			getInput: func() (io.Reader, error) {
				in := strings.NewReader(`
L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
`[1:])
				return in, nil
			},
			wantPassword: 6,
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
			wantPassword: 1026,
		},
	}

	for _, tc := range testCases {
		input, err := tc.getInput()
		require.NoError(t, err)

		dial := day01.Dial{Position: 50}

		_, err = dial.ReadFrom(input)
		require.NoError(t, err)

		gotPW := dial.PasswordV2()
		require.Equal(t, tc.wantPassword, gotPW)
	}
}
