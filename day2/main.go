package day2

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func Solve1() {
	data, err := os.ReadFile("day2/in1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	reports := strings.Split(string(data), "\n")
	safeCount := 0
	for _, report := range reports {
		entries := strings.Split(report, " ")
		levels := []int{}
		for _, entry := range entries {
			level, _ := strconv.Atoi(entry)
			levels = append(levels, level)
		}
		if isSafe(levels) {
			safeCount++
		}
	}
	fmt.Println(safeCount)
}

func isSafe(levels []int) bool {
	diffs := calcDiffs(levels)
	posCount, negCount, _ := countSigns(diffs)
	return (posCount == len(diffs) || negCount == len(diffs)) && isInRange(diffs)
}

func Solve2BruteForce() {
	log.SetOutput(io.Discard) // io.Discard or os.Stdout

	data, err := os.ReadFile("day2/in1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	reports := strings.Split(string(data), "\n")
	safeCount := 0
	for _, report := range reports {
		entries := strings.Split(report, " ")
		levels := []int{}
		for _, entry := range entries {
			level, _ := strconv.Atoi(entry)
			levels = append(levels, level)
		}
		log.Println("")
		log.Printf("Checking: %v\n", levels)
		if isSafe(levels) {
			log.Println("Safe")
			safeCount++
			continue
		}
		for i := 0; i < len(levels); i++ {
			levelsWithSkip := make([]int, len(levels)-1)
			copy(levelsWithSkip, levels[:i])
			copy(levelsWithSkip[i:], levels[i+1:])
			if isSafe(levelsWithSkip) {
				log.Printf("Safe with skip at index %d\n", i)
				safeCount++
				break
			}
		}
	}
	fmt.Println(safeCount)
}

func calcDiffs(levels []int) []int {
	diffs := make([]int, 0, len(levels)-1)
	for i := 0; i < len(levels)-1; i++ {
		diffs = append(diffs, levels[i]-levels[i+1])
	}
	return diffs
}

func countSigns(diffs []int) (int, int, int) {
	posCount, negCount, zeroCount := 0, 0, 0
	for _, diff := range diffs {
		if diff > 0 {
			posCount++
		} else if diff < 0 {
			negCount++
		} else {
			zeroCount++
		}
	}
	return posCount, negCount, zeroCount
}

func isInRange(diffs []int) bool {
	for _, diff := range diffs {
		absDiff := int(math.Abs(float64(diff)))
		if !(0 < absDiff && absDiff < 4) {
			return false
		}
	}
	return true
}

func Solve2() {
	log.SetOutput(io.Discard) // io.Discard or os.Stdout

	data, err := os.ReadFile("day2/in1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	reports := strings.Split(string(data), "\n")
	safeCount := 0
	for _, report := range reports {
		entries := strings.Split(report, " ")
		levels := []int{}
		for _, entry := range entries {
			level, _ := strconv.Atoi(entry)
			levels = append(levels, level)
		}
		log.Println("")
		log.Printf("Checking: %v\n", levels)

		diffs := calcDiffs(levels)
		log.Printf("Diffs: %v\n", diffs)
		posCount, negCount, _ := countSigns(diffs)
		if posCount == len(diffs) || negCount == len(diffs) {
			log.Println("All pos or all neg")
			if isInRange(diffs) {
				log.Printf("Safe: %v\n", levels)
				safeCount++
			}
			continue
		}
		if posCount == len(diffs)-1 {
			log.Println("All but one pos")
			skipIndex := indexOf(diffs, func(diff int) bool { return diff == 0 || diff < 0 })
			if isSafeWithSkips(levels, skipIndex) {
				safeCount++
			}
			continue
		}
		if negCount == len(diffs)-1 {
			log.Println("All but one neg")
			skipIndex := indexOf(diffs, func(diff int) bool { return diff == 0 || diff > 0 })
			if isSafeWithSkips(levels, skipIndex) {
				safeCount++
			}
			continue
		}
	}
	fmt.Println(safeCount)
}

func indexOf[T any](slice []T, predicate func(T) bool) int {
	for i, v := range slice {
		if predicate(v) {
			return i
		}
	}
	return -1
}

func isSafeWithSkips(levels []int, skipIndex int) bool {
	return isSafeWithSkip(levels, skipIndex) || isSafeWithSkip(levels, skipIndex+1)
}

func isSafeWithSkip(levels []int, skipIndex int) bool {
	levelsWithSkip := make([]int, len(levels)-1)
	copy(levelsWithSkip, levels[:skipIndex])
	copy(levelsWithSkip[skipIndex:], levels[skipIndex+1:])
	diffs1 := calcDiffs(levelsWithSkip)
	return isInRange(diffs1)
}
