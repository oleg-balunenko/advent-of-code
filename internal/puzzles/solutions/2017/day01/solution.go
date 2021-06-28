package day01

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/obalunenko/advent-of-code/internal/puzzles"
)

const (
	puzzleName = "day01"
	year       = "2017"
)

type solution struct {
	year string
	name string
}

func (s solution) Name() string {
	return s.name
}

func (s solution) Year() string {
	return s.year
}

func init() {
	puzzles.Register(solution{
		year: year,
		name: puzzleName,
	})
}

func (s solution) Part1(in io.Reader) (string, error) {
	reader := bufio.NewReader(in)

	var (
		idx   int
		first int
		cur   int
		prev  int
		sum   int
	)

	for {
		if idx != 0 {
			prev = cur
		}

		r, _, err := reader.ReadRune()
		if err != nil {
			if errors.Is(err, io.EOF) {
				if prev == first {
					sum += first
				}

				break
			}

			return "", fmt.Errorf("read rune: %w", err)
		}

		cur, err = strconv.Atoi(string(r))
		if err != nil {
			return "", fmt.Errorf("strconv atoi: %w", err)
		}

		if idx == 0 {
			first = cur
		}

		if cur == prev {
			sum += cur
		}

		idx++
	}

	return strconv.Itoa(sum), nil
}

func (s solution) Part2(in io.Reader) (string, error) {
	_ = bufio.NewReader(in)

	return "", puzzles.ErrNotImplemented
}
