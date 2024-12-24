package day11_test

import (
	"testing"

	"github.com/marat-rkh/adventofcode2024/day11"
)

func TestSolve1In0(t *testing.T) {
	res := day11.Solve1("in0.txt")
	expected := 55312
	if res != expected {
		t.Errorf("Expected %d, got %d", expected, res)
	}
}

func TestSolve1In1(t *testing.T) {
	res := day11.Solve1("in1.txt")
	expected := 183620
	if res != expected {
		t.Errorf("Expected %d, got %d", expected, res)
	}
}

func TestSolve2In1(t *testing.T) {
	res := day11.Solve2("in1.txt")
	expected := 220377651399268
	if res != expected {
		t.Errorf("Expected %d, got %d", expected, res)
	}
}
