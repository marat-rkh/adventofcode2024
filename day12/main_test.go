package day12_test

import (
	"testing"

	"github.com/marat-rkh/adventofcode2024/day12"
)

func TestSolve1In0(t *testing.T) {
	res := day12.Solve1("in0.txt")
	expected := 140
	if res != expected {
		t.Errorf("Expected %d, got %d", expected, res)
	}
}

func TestSolve1InDebug(t *testing.T) {
	res := day12.Solve1("in-debug.txt")
	expected := 772
	if res != expected {
		t.Errorf("Expected %d, got %d", expected, res)
	}
}

func TestSolve1In1(t *testing.T) {
	res := day12.Solve1("in1.txt")
	expected := 1522850
	if res != expected {
		t.Errorf("Expected %d, got %d", expected, res)
	}
}

func TestSolve2In0(t *testing.T) {
	res := day12.Solve2("in0.txt")
	expected := 80
	if res != expected {
		t.Errorf("Expected %d, got %d", expected, res)
	}
}

func TestSolve2In1(t *testing.T) {
	res := day12.Solve2("in1.txt")
	expected := 953738
	if res != expected {
		t.Errorf("Expected %d, got %d", expected, res)
	}
}
