package day10

import (
	"github.com/marat-rkh/adventofcode2024/util"
	"github.com/zyedidia/generic/mapset"
)

func Solve1(file string) int {
	input := ReadInput(file)
	res := DoSolve1(input)
	return res
}

func ReadInput(file string) [][]int {
	lines := util.ReadInput(file)
	input := make([][]int, len(lines))
	for i, line := range lines {
		runes := []rune(line)
		input[i] = make([]int, len(runes))
		for j, r := range runes {
			input[i][j] = int(r - '0')
		}
	}
	return input
}

func DoSolve1(input [][]int) int {
	trails := make(map[[2]int]*mapset.Set[[2]int])
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == 0 {
				traverse([2]int{i, j}, input, trails)
			}
		}
	}
	res := 0
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == 0 {
				trailEnds := trails[[2]int{i, j}]
				res += trailEnds.Size()
			}
		}
	}
	return res
}

func traverse(cur [2]int, graph [][]int, trails map[[2]int]*mapset.Set[[2]int]) {
	// Here we rely on the fact that graph is acyclic
	if _, ok := trails[cur]; ok {
		return
	}
	if graph[cur[0]][cur[1]] == 9 {
		trailEnds := mapset.New[[2]int]()
		trailEnds.Put(cur)
		trails[cur] = &trailEnds
		return
	}
	moves := calculateMoves(cur, graph)
	for _, move := range moves {
		traverse(move, graph, trails)
	}
	trailEnds := mapset.New[[2]int]()
	for _, move := range moves {
		if neighborTrailEnds, ok := trails[move]; ok {
			neighborTrailEnds.Each(func(trailEnd [2]int) {
				trailEnds.Put(trailEnd)
			})
		}
	}
	trails[cur] = &trailEnds
}

func calculateMoves(cur [2]int, graph [][]int) [][2]int {
	moves := [][2]int{}
	left := [2]int{cur[0], cur[1] - 1}
	if canMove(cur, left, graph) {
		moves = append(moves, left)
	}
	right := [2]int{cur[0], cur[1] + 1}
	if canMove(cur, right, graph) {
		moves = append(moves, right)
	}
	up := [2]int{cur[0] - 1, cur[1]}
	if canMove(cur, up, graph) {
		moves = append(moves, up)
	}
	down := [2]int{cur[0] + 1, cur[1]}
	if canMove(cur, down, graph) {
		moves = append(moves, down)
	}
	return moves
}

func canMove(cur [2]int, next [2]int, graph [][]int) bool {
	if next[0] >= 0 && next[0] < len(graph) && next[1] >= 0 && next[1] < len(graph[next[0]]) {
		curHeight := graph[cur[0]][cur[1]]
		nextHeight := graph[next[0]][next[1]]
		return nextHeight-curHeight == 1
	}
	return false
}

func Solve2(file string) int {
	input := ReadInput(file)
	res := DoSolve2(input)
	return res
}

func DoSolve2(input [][]int) int {
	trails := make(map[[2]int]int)
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == 0 {
				traverse2([2]int{i, j}, input, trails)
			}
		}
	}
	res := 0
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == 0 {
				if trailsCount, ok := trails[[2]int{i, j}]; ok {
					res += trailsCount
				}
			}
		}
	}
	return res
}

func traverse2(cur [2]int, graph [][]int, trails map[[2]int]int) {
	// Here we rely on the fact that graph is acyclic
	if _, ok := trails[cur]; ok {
		return
	}
	if graph[cur[0]][cur[1]] == 9 {
		trails[cur] = 1
		return
	}
	moves := calculateMoves(cur, graph)
	for _, move := range moves {
		traverse2(move, graph, trails)
	}
	trailsCount := 0
	for _, move := range moves {
		if neighborTrailsCount, ok := trails[move]; ok {
			trailsCount += neighborTrailsCount
		}
	}
	trails[cur] = trailsCount
}
