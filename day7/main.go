package day7

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"

	"github.com/marat-rkh/adventofcode2024/util"
	"github.com/samber/lo"
)

func Solve1() {
	lines := util.ReadInput("day7/in1.txt")
	res := DoSolve1(lines)
	fmt.Println(res)
}

func DoSolve1(lines []string) int64 {
	res := int64(0)
	for _, line := range lines {
		pair := strings.Split(line, ": ")
		testValue, _ := strconv.ParseInt(pair[0], 10, 64)
		operands := lo.Map(strings.Split(pair[1], " "), func(s string, _ int) int64 {
			num, _ := strconv.ParseInt(s, 10, 64)
			return num
		})
		if canCombine(testValue, operands) {
			res += testValue
		}
	}
	return res
}

func canCombine(testValue int64, operands []int64) bool {
	combinationsNum := int64(math.Pow(2, float64(len(operands)-1)))
	for combination := int64(0); combination < combinationsNum; combination++ {
		acc := operands[0]
		for i := 0; i < len(operands)-1; i++ {
			if combination&(1<<i) != 0 {
				acc += operands[i+1]
			} else {
				acc *= operands[i+1]
			}
			if acc > testValue {
				break
			}
		}
		if acc == testValue {
			return true
		}
	}
	return false
}

func Solve2() {
	lines := util.ReadInput("day7/in1.txt")
	res := DoSolve2(lines)
	fmt.Println(res)
}

func DoSolve2(lines []string) int64 {
	res := int64(0)
	for _, line := range lines {
		pair := strings.Split(line, ": ")
		testValue, _ := strconv.ParseInt(pair[0], 10, 64)
		operands := lo.Map(strings.Split(pair[1], " "), func(s string, _ int) int64 {
			num, _ := strconv.ParseInt(s, 10, 64)
			return num
		})
		if canCombine2(testValue, operands) {
			res += testValue
		}
	}
	return res
}

func canCombine2(testValue int64, operands []int64) bool {
	combinationsNum := int64(math.Pow(3, float64(len(operands)-1)))
	for combination := int64(0); combination < combinationsNum; combination++ {
		combinationText := lo.Reverse([]rune(big.NewInt(combination).Text(3)))
		acc := operands[0]
		for i := 0; i < len(operands)-1; i++ {
			if acc > testValue {
				break
			}
			if i >= len(combinationText) || combinationText[i] == '0' {
				acc += operands[i+1]
			} else if combinationText[i] == '1' {
				acc *= operands[i+1]
			} else {
				concatVal, _ := strconv.ParseInt(strconv.FormatInt(acc, 10)+strconv.FormatInt(operands[i+1], 10), 10, 64)
				acc = concatVal
			}
		}
		if acc == testValue {
			return true
		}
	}
	return false
}
