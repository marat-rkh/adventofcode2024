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
		if isSafe(levels) {
			log.Printf("Safe: %v\n", levels)
			safeCount++
			continue
		}

		prevDiff := 0
		signChangeIndex := -1
		for i := 0; i < len(levels)-1; i++ {
			diff := levels[i] - levels[i+1]
			if prevDiff != 0 && (diff == 0 || diff > 0 && prevDiff < 0 || diff < 0 && prevDiff > 0) {
				signChangeIndex = i
				break
			}
			prevDiff = diff
		}
		log.Printf("Trying skips for: %v\n", levels)
		if signChangeIndex != -1 {
			log.Printf("Sign change at %d\n", signChangeIndex)
			levelsWithSkip1 := append(levels[:signChangeIndex], levels[signChangeIndex+1:]...)
			if isSafe(levelsWithSkip1) {
				log.Printf("Safe: %v\n", levelsWithSkip1)
				safeCount++
			} else if signChangeIndex+2 < len(levels) {
				levelsWithSkip2 := append(levels[:signChangeIndex+1], levels[signChangeIndex+2:]...)
				if isSafe(levelsWithSkip2) {
					log.Printf("Safe: %v\n", levelsWithSkip2)
					safeCount++
				}
			}
		}
	}
	fmt.Println(safeCount)
}
