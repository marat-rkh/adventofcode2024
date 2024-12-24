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
	res := 0
	cache := make(map[[2]int64]int)
	for _, num := range curNums {
		res += processNum(num, iterations, cache)
	}
	return res
}

func processNum(num int64, iterations int, cache map[[2]int64]int) int {
	if iterations == 0 {
		return 1
	}
	if count, ok := cache[[2]int64{num, int64(iterations)}]; ok {
		return count
	}
	count := 0
	if num == 0 {
		count += processNum(1, iterations-1, cache)
	} else if digitsCount := countDigits(num); digitsCount%2 == 0 {
		divisor := int64(math.Pow10(digitsCount / 2))
		num1, num2 := num/divisor, num%divisor
		count += processNum(num1, iterations-1, cache)
		count += processNum(num2, iterations-1, cache)
	} else {
		mul2024 := num * 2024
		count += processNum(mul2024, iterations-1, cache)
	}
	cache[[2]int64{num, int64(iterations)}] = count
	return count
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
