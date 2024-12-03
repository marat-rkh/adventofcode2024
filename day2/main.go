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
	firstDiff := 0
	for i := 0; i < len(levels)-1; i++ {
		diff := levels[i] - levels[i+1]
		absDiff := int(math.Abs(float64(diff)))
		if firstDiff == 0 {
			firstDiff = diff
		}
		if !(0 < absDiff && absDiff < 4 && (firstDiff > 0 && diff > 0 || firstDiff < 0 && diff < 0)) {
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
			if isSafeWithSkip(levels, skipIndex) {
				safeCount++
			}
			continue
		}
		if negCount == len(diffs)-1 {
			log.Println("All but one neg")
			skipIndex := indexOf(diffs, func(diff int) bool { return diff == 0 || diff > 0 })
			if isSafeWithSkip(levels, skipIndex) {
				safeCount++
			}
			continue
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

func isInRange(diffs []int) bool {
	for _, diff := range diffs {
		absDiff := int(math.Abs(float64(diff)))
		if !(0 < absDiff && absDiff < 4) {
			return false
		}
	}
	return true
}

func indexOf[T any](slice []T, predicate func(T) bool) int {
	for i, v := range slice {
		if predicate(v) {
			return i
		}
	}
	return -1
}

func isSafeWithSkip(levels []int, skipIndex int) bool {
	levels1 := append(levels[:skipIndex], levels[skipIndex+1:]...)
	diffs1 := calcDiffs(levels1)
	log.Printf("Levels with skip at index %d: %v\n", skipIndex, levels1)
	if isInRange(diffs1) {
		log.Printf("Safe with index %d skipped: %v\n", skipIndex, levels)
		return true
	}
	levels2 := append(levels[:skipIndex+1], levels[skipIndex+2:]...)
	diffs2 := calcDiffs(levels2)
	log.Printf("Levels with skip at index %d: %v\n", skipIndex+1, levels2)
	if isInRange(diffs2) {
		log.Printf("Safe with index %d skipped: %v\n", skipIndex+1, levels)
		return true
	}
	return false
}
