package day01

import (
	"bufio"
	"io"

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
	_ = bufio.NewReader(in)

	return "", puzzles.ErrNotImplemented
}

func (s solution) Part2(in io.Reader) (string, error) {
	_ = bufio.NewReader(in)

	return "", puzzles.ErrNotImplemented
}
