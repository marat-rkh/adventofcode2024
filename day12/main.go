package day12

import (
	"github.com/marat-rkh/adventofcode2024/util"
)

func Solve1(file string) int {
	field := util.ReadInput(file)
	measurements := [][3]int{}
	visited := make(map[[2]int]bool)
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			if visited[[2]int{i, j}] {
				continue
			}
			queue := [][2]int{{i, j}}
			curArea := 0
			curPerimeter := 0
			for len(queue) > 0 {
				cur := queue[0]
				queue = queue[1:]
				if visited[cur] {
					continue
				}
				visited[cur] = true
				curArea++
				moves := [][2]int{
					{cur[0], cur[1] - 1}, // left
					{cur[0], cur[1] + 1}, // right
					{cur[0] - 1, cur[1]}, // up
					{cur[0] + 1, cur[1]}, // down
				}
				for _, move := range moves {
					if !isWithinBounds(move, field) || field[move[0]][move[1]] != field[cur[0]][cur[1]] {
						curPerimeter++
					} else {
						queue = append(queue, move)
					}
				}
			}
			measurements = append(measurements, [3]int{int(field[i][j]), curArea, curPerimeter})
		}
	}
	res := 0
	for _, measurement := range measurements {
		res += measurement[1] * measurement[2]
	}
	return res
}

func isWithinBounds(pos [2]int, graph []string) bool {
	return pos[0] >= 0 && pos[0] < len(graph) && pos[1] >= 0 && pos[1] < len(graph[pos[0]])
}
