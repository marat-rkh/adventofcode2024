package day7_test

import (
	"testing"

	"github.com/marat-rkh/adventofcode2024/day7"
	"github.com/marat-rkh/adventofcode2024/util"
)

func BenchmarkDoSolve1(b *testing.B) {
	lines := util.ReadInput("in1.txt")
	if lines == nil {
		b.Fatal("Failed to read input")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		day7.DoSolve1(lines)
	}
}

func BenchmarkDoSolve2(b *testing.B) {
	lines := util.ReadInput("in1.txt")
	if lines == nil {
		b.Fatal("Failed to read input")
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		day7.DoSolve2(lines)
	}
}
