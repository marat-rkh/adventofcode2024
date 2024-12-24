package day1

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Solve1() {
	data, err := os.ReadFile("day1/in1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	var list1, list2 []int
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		pairs := strings.Split(line, "   ")
		num1, _ := strconv.Atoi(pairs[0])
		num2, _ := strconv.Atoi(pairs[1])
		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}
	sort.Ints(list1)
	sort.Ints(list2)
	totalDist := 0
	for i := 0; i < len(list1); i++ {
		totalDist += int(math.Abs(float64(list1[i] - list2[i])))
	}
	fmt.Println(totalDist)
}

func Solve2() {
	data, err := os.ReadFile("day1/in1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	var list1 []int
	dict2 := make(map[int]int)
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		pairs := strings.Split(line, "   ")
		num1, _ := strconv.Atoi(pairs[0])
		num2, _ := strconv.Atoi(pairs[1])
		list1 = append(list1, num1)
		if count2, ok := dict2[num2]; ok {
			dict2[num2] = count2 + 1
		} else {
			dict2[num2] = 1
		}
	}
	sim := 0
	for i := 0; i < len(list1); i++ {
		num1 := list1[i]
		if count2, ok := dict2[num1]; ok {
			sim += num1 * count2
		}
	}
	fmt.Println(sim)
}
