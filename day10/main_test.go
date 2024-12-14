package day10_test

import (
	"testing"

	"github.com/marat-rkh/adventofcode2024/day10"
)

func TestSolve1In0(t *testing.T) {
	res := day10.Solve1("in0.txt")
	expected := 36
	if res != expected {
		t.Errorf("Expected %d, got %d", expected, res)
	}
}

func TestSolve1In1(t *testing.T) {
	res := day10.Solve1("in1.txt")
	expected := 688
	if res != expected {
		t.Errorf("Expected %d, got %d", expected, res)
	}
}
