package day4

import (
	"fmt"
	"os"
	"strings"
)

func Solve1() {
	data, err := os.ReadFile("day4/in1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	lines := strings.Split(string(data), "\n")
	xmasCount := 0
	lenLines := len(lines)
	for i := 0; i < lenLines; i++ {
		lenCols := len(lines[i])
		for j := 0; j < lenCols; j++ {
			if lines[i][j] == 'X' {
				if j+3 < lenCols && lines[i][j+1] == 'M' && lines[i][j+2] == 'A' && lines[i][j+3] == 'S' {
					xmasCount++
				}
				if j-3 >= 0 && lines[i][j-1] == 'M' && lines[i][j-2] == 'A' && lines[i][j-3] == 'S' {
					xmasCount++
				}
				if i+3 < lenLines && lines[i+1][j] == 'M' && lines[i+2][j] == 'A' && lines[i+3][j] == 'S' {
					xmasCount++
				}
				if i-3 >= 0 && lines[i-1][j] == 'M' && lines[i-2][j] == 'A' && lines[i-3][j] == 'S' {
					xmasCount++
				}
				if i+3 < lenLines && j+3 < lenCols && lines[i+1][j+1] == 'M' && lines[i+2][j+2] == 'A' && lines[i+3][j+3] == 'S' {
					xmasCount++
				}
				if i-3 >= 0 && j-3 >= 0 && lines[i-1][j-1] == 'M' && lines[i-2][j-2] == 'A' && lines[i-3][j-3] == 'S' {
					xmasCount++
				}
				if i+3 < lenLines && j-3 >= 0 && lines[i+1][j-1] == 'M' && lines[i+2][j-2] == 'A' && lines[i+3][j-3] == 'S' {
					xmasCount++
				}
				if i-3 >= 0 && j+3 < lenCols && lines[i-1][j+1] == 'M' && lines[i-2][j+2] == 'A' && lines[i-3][j+3] == 'S' {
					xmasCount++
				}
			}
		}
	}
	fmt.Println(xmasCount)
}

func Solve2() {
	data, err := os.ReadFile("day4/in1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	lines := strings.Split(string(data), "\n")
	xmasCount := 0
	lenLines := len(lines)
	for i := 0; i < lenLines; i++ {
		lenCols := len(lines[i])
		for j := 0; j < lenCols; j++ {
			if lines[i][j] == 'A' && j-1 >= 0 && i-1 >= 0 && j+1 < lenCols && i+1 < lenLines &&
				(lines[i-1][j-1] == 'M' && lines[i+1][j+1] == 'S' || lines[i-1][j-1] == 'S' && lines[i+1][j+1] == 'M') &&
				(lines[i-1][j+1] == 'M' && lines[i+1][j-1] == 'S' || lines[i-1][j+1] == 'S' && lines[i+1][j-1] == 'M') {
				xmasCount++
			}
		}
	}
	fmt.Println(xmasCount)
}
