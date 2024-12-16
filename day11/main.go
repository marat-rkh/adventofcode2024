package day11

import (
	"math"
	"strconv"
	"strings"

	"github.com/marat-rkh/adventofcode2024/util"
	"github.com/samber/lo"
)

func Solve1(file string) int {
	return solve(file, 25)
}

func Solve2(file string) int {
	return solve(file, 75)
}

func solve(file string, iterations int) int {
	curNums := ReadInput(file)
	nextNums := []int64{}
	for i := 0; i < iterations; i++ {
		for _, num := range curNums {
			if num == 0 {
				nextNums = append(nextNums, 1)
			} else if digitsCount := countDigits(num); digitsCount%2 == 0 {
				divisor := int64(math.Pow10(digitsCount / 2))
				num1, num2 := num/divisor, num%divisor
				nextNums = append(nextNums, num1, num2)
			} else {
				nextNums = append(nextNums, num*2024)
			}
		}
		curNums = nextNums
		nextNums = []int64{}
	}
	return len(curNums)
}

func ReadInput(file string) []int64 {
	line := util.ReadInput(file)[0]
	parts := strings.Split(line, " ")
	return lo.Map(parts, func(part string, _ int) int64 {
		num, _ := strconv.ParseInt(part, 10, 64)
		return num
	})
}

func countDigits(num int64) int {
	count := 0
	for num != 0 {
		num /= 10
		count++
	}
	return count
}
