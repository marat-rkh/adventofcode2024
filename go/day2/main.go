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
			if isSafeWithSkip(levels, i) {
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

func isSafeWithSkip(levels []int, skipIndex int) bool {
	levelsWithSkip := make([]int, len(levels)-1)
	copy(levelsWithSkip, levels[:skipIndex])
	copy(levelsWithSkip[skipIndex:], levels[skipIndex+1:])
	return isSafe(levelsWithSkip)
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
		posCount, negCount, _ := countSigns(diffs)
		if posCount == len(diffs) || negCount == len(diffs) {
			if isInRange(diffs) {
				log.Println("Safe")
				safeCount++
			} else if isSafeWithSkip(levels, 0) {
				log.Println("Safe with skip at index 0")
				safeCount++
			} else if isSafeWithSkip(levels, len(levels)-1) {
				log.Printf("Safe with skip at index %d\n", len(levels)-1)
				safeCount++
			}
		} else if posCount == len(diffs)-1 || negCount == len(diffs)-1 {
			isMostlyPos := posCount == len(diffs)-1
			skipIndex := indexOf(diffs, func(diff int) bool { return diff == 0 || (isMostlyPos && diff < 0 || !isMostlyPos && diff > 0) })
			if isSafeWithSkip(levels, skipIndex) {
				log.Printf("Safe with skip at index %d\n", skipIndex)
				safeCount++
			} else if isSafeWithSkip(levels, skipIndex+1) {
				log.Printf("Safe with skip at index %d\n", skipIndex+1)
				safeCount++
			}
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
