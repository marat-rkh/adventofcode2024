package day12

import (
	"github.com/marat-rkh/adventofcode2024/util"
	"github.com/zyedidia/generic/mapset"
)

func Solve1(file string) int {
	field := util.ReadInput(file)
	plots := calcPlots(field)
	areas := make(map[int]int)
	perimeters := make(map[int]int)
	for i := range plots {
		for j := range plots[i] {
			areas[plots[i][j]]++
			moves := [][2]int{{i, j - 1}, {i, j + 1}, {i - 1, j}, {i + 1, j}}
			for _, move := range moves {
				if !isWithinBounds(move, field) || field[move[0]][move[1]] != field[i][j] {
					perimeters[plots[i][j]]++
				}
			}
		}
	}

	res := 0
	for plotId, area := range areas {
		res += area * perimeters[plotId]
	}
	return res
}

func calcPlots(field []string) [][]int {
	plots := make([][]int, len(field))
	for i := range plots {
		plots[i] = make([]int, len(field[i]))
	}
	plotId := 1
	visited := make(map[[2]int]bool)
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			if visited[[2]int{i, j}] {
				continue
			}
			queue := [][2]int{{i, j}}
			for len(queue) > 0 {
				cur := queue[0]
				queue = queue[1:]
				if visited[cur] {
					continue
				}
				visited[cur] = true
				plots[cur[0]][cur[1]] = plotId
				moves := [][2]int{{cur[0], cur[1] - 1}, {cur[0], cur[1] + 1}, {cur[0] - 1, cur[1]}, {cur[0] + 1, cur[1]}}
				for _, move := range moves {
					if isWithinBounds(move, field) && field[move[0]][move[1]] == field[cur[0]][cur[1]] {
						queue = append(queue, move)
					}
				}
			}
			plotId++
		}
	}
	return plots
}

func isWithinBounds(pos [2]int, graph []string) bool {
	return pos[0] >= 0 && pos[0] < len(graph) && pos[1] >= 0 && pos[1] < len(graph[pos[0]])
}

type fence byte

const (
	Up fence = iota
	Down
	Left
	Right
)

func Solve2(file string) int {
	field := util.ReadInput(file)
	plots := calcPlots(field)

	areas := make(map[int]int)
	for i := range plots {
		for j := range plots[i] {
			areas[plots[i][j]]++
		}
	}

	fences := [][]mapset.Set[fence]{}
	for i := 0; i < len(field); i++ {
		fences = append(fences, make([]mapset.Set[fence], len(field[i])))
		for j := 0; j < len(field[i]); j++ {
			fences[i][j] = mapset.New[fence]()
			if i == 0 {
				fences[i][j].Put(Up)
			}
			if i == len(field)-1 {
				fences[i][j].Put(Down)
			}
			if j == 0 {
				fences[i][j].Put(Left)
			}
			if j == len(field[i])-1 {
				fences[i][j].Put(Right)
			}
			left := [2]int{i, j - 1}
			if isWithinBounds(left, field) && field[i][j] != field[left[0]][left[1]] {
				fences[i][j].Put(Left)
			}
			right := [2]int{i, j + 1}
			if isWithinBounds(right, field) && field[i][j] != field[right[0]][right[1]] {
				fences[i][j].Put(Right)
			}
			down := [2]int{i + 1, j}
			if isWithinBounds(down, field) && field[i][j] != field[down[0]][down[1]] {
				fences[i][j].Put(Down)
			}
			up := [2]int{i - 1, j}
			if isWithinBounds(up, field) && field[i][j] != field[up[0]][up[1]] {
				fences[i][j].Put(Up)
			}
		}
	}

	segments := make(map[int]int) // plot id -> segments count
	for i := 0; i < len(plots); i++ {
		for j := 0; j < len(plots[i]); j++ {
			if fences[i][j].Has(Up) {
				rightUp := [2]int{i - 1, j + 1}
				if fences[i][j].Has(Right) ||
					isWithinBounds(rightUp, field) && plots[rightUp[0]][rightUp[1]] == plots[i][j] && fences[rightUp[0]][rightUp[1]].Has(Left) {
					segments[plots[i][j]]++
				}
			}
			if fences[i][j].Has(Down) {
				rightDown := [2]int{i + 1, j + 1}
				if fences[i][j].Has(Right) ||
					isWithinBounds(rightDown, field) && plots[rightDown[0]][rightDown[1]] == plots[i][j] && fences[rightDown[0]][rightDown[1]].Has(Left) {
					segments[plots[i][j]]++
				}
			}
		}
	}
	for j := 0; j < len(plots[0]); j++ {
		for i := 0; i < len(plots); i++ {
			if fences[i][j].Has(Left) {
				leftDown := [2]int{i + 1, j - 1}
				if fences[i][j].Has(Down) ||
					isWithinBounds(leftDown, field) && plots[leftDown[0]][leftDown[1]] == plots[i][j] && fences[leftDown[0]][leftDown[1]].Has(Up) {
					segments[plots[i][j]]++
				}
			}
			if fences[i][j].Has(Right) {
				rightDown := [2]int{i + 1, j + 1}
				if fences[i][j].Has(Down) ||
					isWithinBounds(rightDown, field) && plots[rightDown[0]][rightDown[1]] == plots[i][j] && fences[rightDown[0]][rightDown[1]].Has(Up) {
					segments[plots[i][j]]++
				}
			}
		}
	}

	res := 0
	for plotId, area := range areas {
		res += area * segments[plotId]
	}
	return res
}
