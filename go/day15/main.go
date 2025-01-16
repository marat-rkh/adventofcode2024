package day15

import (
	"github.com/marat-rkh/adventofcode2024/util"
)

const (
	player = '@'
	wall   = '#'
	box    = 'O'
	empty  = '.'
)

func Solve1(file string) int {
	lines := util.ReadInput(file)
	field, commands := parseInput(lines)

	// Convert field to rune slices for easier modification
	fieldRunes := make([][]rune, len(field))
	for i := range field {
		fieldRunes[i] = []rune(field[i])
	}

	// Find player position
	startI, startJ := findPlayer(fieldRunes)

	for _, cmd := range commands {
		switch cmd {
		case '^':
			moveUp(fieldRunes, &startI, &startJ)
		case 'v':
			moveDown(fieldRunes, &startI, &startJ)
		case '<':
			moveLeft(fieldRunes, &startI, &startJ)
		case '>':
			moveRight(fieldRunes, &startI, &startJ)
		}
	}

	return calculateScore(fieldRunes)
}

func parseInput(lines []string) ([]string, string) {
	var field []string
	var commands string

	emptyLineIdx := -1
	for i, line := range lines {
		if line == "" {
			emptyLineIdx = i
			break
		}
		field = append(field, line)
	}

	for i := emptyLineIdx + 1; i < len(lines); i++ {
		commands += lines[i]
	}

	return field, commands
}

func findPlayer(field [][]rune) (int, int) {
	for i := range field {
		for j := range field[i] {
			if field[i][j] == player {
				return i, j
			}
		}
	}
	return 0, 0
}

func moveUp(field [][]rune, startI, startJ *int) {
	if field[*startI-1][*startJ] == wall {
		return
	}

	// Count consecutive boxes
	boxCount := 0
	for i := *startI - 1; field[i][*startJ] == box; i-- {
		boxCount++
	}

	// Check if we can move boxes
	if field[*startI-boxCount-1][*startJ] != wall {
		// Move boxes
		if boxCount > 0 {
			for i := *startI - boxCount; i < *startI; i++ {
				field[i-1][*startJ] = box
			}
		}
		// Move player
		field[*startI][*startJ] = empty
		*startI--
		field[*startI][*startJ] = player
	}
}

func moveDown(field [][]rune, startI, startJ *int) {
	if field[*startI+1][*startJ] == wall {
		return
	}

	// Count consecutive boxes
	boxCount := 0
	for i := *startI + 1; field[i][*startJ] == box; i++ {
		boxCount++
	}

	// Check if we can move boxes
	if field[*startI+boxCount+1][*startJ] != wall {
		// Move boxes
		if boxCount > 0 {
			for i := *startI + boxCount; i > *startI; i-- {
				field[i+1][*startJ] = box
			}
		}
		// Move player
		field[*startI][*startJ] = empty
		*startI++
		field[*startI][*startJ] = player
	}
}

func moveLeft(field [][]rune, startI, startJ *int) {
	if field[*startI][*startJ-1] == wall {
		return
	}

	// Count consecutive boxes
	boxCount := 0
	for j := *startJ - 1; field[*startI][j] == box; j-- {
		boxCount++
	}

	// Check if we can move boxes
	if field[*startI][*startJ-boxCount-1] != wall {
		// Move boxes
		if boxCount > 0 {
			for j := *startJ - boxCount; j < *startJ; j++ {
				field[*startI][j-1] = box
			}
		}
		// Move player
		field[*startI][*startJ] = empty
		*startJ--
		field[*startI][*startJ] = player
	}
}

func moveRight(field [][]rune, startI, startJ *int) {
	if field[*startI][*startJ+1] == wall {
		return
	}

	// Count consecutive boxes
	boxCount := 0
	for j := *startJ + 1; field[*startI][j] == box; j++ {
		boxCount++
	}

	// Check if we can move boxes
	if field[*startI][*startJ+boxCount+1] != wall {
		// Move boxes
		if boxCount > 0 {
			for j := *startJ + boxCount; j > *startJ; j-- {
				field[*startI][j+1] = box
			}
		}
		// Move player
		field[*startI][*startJ] = empty
		*startJ++
		field[*startI][*startJ] = player
	}
}

func calculateScore(field [][]rune) int {
	score := 0
	for i := range field {
		for j := range field[i] {
			if field[i][j] == box {
				score += i*100 + j
			}
		}
	}
	return score
}
