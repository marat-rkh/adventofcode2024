package day3

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var mulRegex = regexp.MustCompile(`mul\((?P<num1>\d{1,3}),(?P<num2>\d{1,3})\)`)

func Solve1() {
	data, err := os.ReadFile("day3/in1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	memory := string(data)
	sum := calcMulSum(memory)
	fmt.Println(sum)
}

func calcMulSum(memory string) int64 {
	matches := mulRegex.FindAllStringSubmatch(memory, -1)
	sum := int64(0)
	for _, match := range matches {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		sum += int64(num1) * int64(num2)
	}
	return sum
}

func Solve2() {
	data, err := os.ReadFile("day3/in1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	memory := "do()" + string(data) + "don't()"
	sum := int64(0)
	for len(memory) > 0 {
		start := strings.Index(memory, "do()")
		if start == -1 {
			break
		}
		end := start + 4 + strings.Index(memory[start+4:], "don't()")
		sum += calcMulSum(memory[start:end])
		memory = memory[end+7:]
	}
	fmt.Println(sum)
}
