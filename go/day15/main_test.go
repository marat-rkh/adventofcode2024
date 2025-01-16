package day15_test

import (
	"testing"

	"github.com/marat-rkh/adventofcode2024/day15"
)

func TestSolve1In0(t *testing.T) {
	res := day15.Solve1("in0.txt")
	expected := 10092
	if res != expected {
		t.Errorf("Expected %d, got %d", expected, res)
	}
}

func TestSolve1InDebug(t *testing.T) {
	res := day15.Solve1("in-debug.txt")
	expected := 2028
	if res != expected {
		t.Errorf("Expected %d, got %d", expected, res)
	}
}

func TestSolve1In1(t *testing.T) {
	res := day15.Solve1("in1.txt")
	expected := 1476771
	if res != expected {
		t.Errorf("Expected %d, got %d", expected, res)
	}
}
