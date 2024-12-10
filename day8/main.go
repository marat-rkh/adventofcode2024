package day8

import (
	"fmt"

	"github.com/marat-rkh/adventofcode2024/util"
	"github.com/zyedidia/generic/mapset"
)

func Solve1() {
	lines := util.ReadInput("day8/in1.txt")
	res := DoSolve1(lines)
	fmt.Println(res)
}

func DoSolve1(lines []string) int {
	antennasByType := make(map[rune]*mapset.Set[[2]int])
	for i, line := range lines {
		for j, antenna := range line {
			if antenna != '.' {
				if _, ok := antennasByType[antenna]; !ok {
					antennaSet := mapset.New[[2]int]()
					antennasByType[antenna] = &antennaSet
				}
				antennasByType[antenna].Put([2]int{i, j})
			}
		}
	}
	antinodeSet := mapset.New[[2]int]()
	for _, antennaSet := range antennasByType {
		antennas := make([][2]int, 0, antennaSet.Size())
		antennaSet.Each(func(antenna [2]int) {
			antennas = append(antennas, antenna)
		})
		for i := 0; i < len(antennas); i++ {
			for j := i + 1; j < len(antennas); j++ {
				xDiff := antennas[j][0] - antennas[i][0]
				absXDiff := xDiff
				if absXDiff < 0 {
					absXDiff = -absXDiff
				}
				yDiff := antennas[j][1] - antennas[i][1]
				absYDiff := yDiff
				if absYDiff < 0 {
					absYDiff = -absYDiff
				}
				antinode1 := [2]int{-1, -1}
				antinode2 := [2]int{-1, -1}
				if xDiff >= 0 {
					if yDiff >= 0 {
						antinode1 = [2]int{antennas[i][0] - absXDiff, antennas[i][1] - absYDiff}
						antinode2 = [2]int{antennas[j][0] + absXDiff, antennas[j][1] + absYDiff}
					} else {
						antinode1 = [2]int{antennas[i][0] - absXDiff, antennas[i][1] + absYDiff}
						antinode2 = [2]int{antennas[j][0] + absXDiff, antennas[j][1] - absYDiff}
					}
				} else {
					if yDiff >= 0 {
						antinode1 = [2]int{antennas[i][0] + absXDiff, antennas[i][1] - absYDiff}
						antinode2 = [2]int{antennas[j][0] - absXDiff, antennas[j][1] + absYDiff}
					} else {
						antinode1 = [2]int{antennas[i][0] + absXDiff, antennas[i][1] + absYDiff}
						antinode2 = [2]int{antennas[j][0] - absXDiff, antennas[j][1] - absYDiff}
					}
				}
				if antinode1[0] >= 0 && antinode1[0] < len(lines) && antinode1[1] >= 0 && antinode1[1] < len(lines[0]) {
					antinodeSet.Put(antinode1)
				}
				if antinode2[0] >= 0 && antinode2[0] < len(lines) && antinode2[1] >= 0 && antinode2[1] < len(lines[0]) {
					antinodeSet.Put(antinode2)
				}
			}
		}
	}
	return antinodeSet.Size()
}

func Solve2() {
	lines := util.ReadInput("day8/in1.txt")
	res := DoSolve2(lines)
	fmt.Println(res)
}

func DoSolve2(lines []string) int {
	antennasByType := make(map[rune]*mapset.Set[[2]int])
	for i, line := range lines {
		for j, antenna := range line {
			if antenna != '.' {
				if _, ok := antennasByType[antenna]; !ok {
					antennaSet := mapset.New[[2]int]()
					antennasByType[antenna] = &antennaSet
				}
				antennasByType[antenna].Put([2]int{i, j})
			}
		}
	}
	antinodeSet := mapset.New[[2]int]()
	for _, antennaSet := range antennasByType {
		antennas := make([][2]int, 0, antennaSet.Size())
		antennaSet.Each(func(antenna [2]int) {
			antennas = append(antennas, antenna)
		})
		for i := 0; i < len(antennas); i++ {
			for j := i + 1; j < len(antennas); j++ {
				antinodeSet.Put(antennas[i])
				antinodeSet.Put(antennas[j])
				xDiff := antennas[j][0] - antennas[i][0]
				absXDiff := xDiff
				if absXDiff < 0 {
					absXDiff = -absXDiff
				}
				yDiff := antennas[j][1] - antennas[i][1]
				absYDiff := yDiff
				if absYDiff < 0 {
					absYDiff = -absYDiff
				}
				for k := 1; ; k++ {
					antinode1 := [2]int{-1, -1}
					antinode2 := [2]int{-1, -1}
					absXDiff := absXDiff * k
					absYDiff := absYDiff * k
					if xDiff >= 0 {
						if yDiff >= 0 {
							antinode1 = [2]int{antennas[i][0] - absXDiff, antennas[i][1] - absYDiff}
							antinode2 = [2]int{antennas[j][0] + absXDiff, antennas[j][1] + absYDiff}
						} else {
							antinode1 = [2]int{antennas[i][0] - absXDiff, antennas[i][1] + absYDiff}
							antinode2 = [2]int{antennas[j][0] + absXDiff, antennas[j][1] - absYDiff}
						}
					} else {
						if yDiff >= 0 {
							antinode1 = [2]int{antennas[i][0] + absXDiff, antennas[i][1] - absYDiff}
							antinode2 = [2]int{antennas[j][0] - absXDiff, antennas[j][1] + absYDiff}
						} else {
							antinode1 = [2]int{antennas[i][0] + absXDiff, antennas[i][1] + absYDiff}
							antinode2 = [2]int{antennas[j][0] - absXDiff, antennas[j][1] - absYDiff}
						}
					}
					isAdded := false
					if antinode1[0] >= 0 && antinode1[0] < len(lines) && antinode1[1] >= 0 && antinode1[1] < len(lines[0]) {
						antinodeSet.Put(antinode1)
						isAdded = true
					}
					if antinode2[0] >= 0 && antinode2[0] < len(lines) && antinode2[1] >= 0 && antinode2[1] < len(lines[0]) {
						antinodeSet.Put(antinode2)
						isAdded = true
					}
					if !isAdded {
						break
					}
				}
			}
		}
	}
	return antinodeSet.Size()
}
