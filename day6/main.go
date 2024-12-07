package day6

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/samber/lo"
)

const (
	Up    = '^'
	Right = '>'
	Down  = 'v'
	Left  = '<'

	Obstruction = '#'
	Empty       = '.'
	MarkVert    = '|'
	MarkHoriz   = '-'
	MarkCross   = '+'
)

func Solve1() {
	log.SetOutput(os.Stdout)
	data, err := os.ReadFile("day6/in0.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	lines := strings.Split(string(data), "\n")
	field := lo.Map(lines, func(line string, _ int) []rune {
		return []rune(line)
	})

	i0, j0 := 0, 0
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			if field[i][j] == Up || field[i][j] == Right || field[i][j] == Down || field[i][j] == Left {
				i0, j0 = i, j
			}
		}
	}
	direction0 := field[i0][j0]

	i, j := i0, j0
	direction := direction0
	for 0 <= i && i < len(field) && 0 <= j && j < len(field[0]) {
		switch direction {
		case Up:
			i = moveUp(field, i, j)
			if i >= 0 {
				direction = Right
			}
		case Right:
			j = moveRight(field, i, j)
			if j < len(field[i]) {
				direction = Down
			}
		case Down:
			i = moveDown(field, i, j)
			if i < len(field) {
				direction = Left
			}
		case Left:
			j = moveLeft(field, i, j)
			if j >= 0 {
				direction = Up
			}
		default:
			panic("invalid direction")
		}
	}
	field[i0][j0] = direction0
	logField(field)

	markedCount := 0
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			if field[i][j] != Empty && field[i][j] != Obstruction {
				markedCount++
			}
		}
	}
	fmt.Println(markedCount)
}

func moveUp(field [][]rune, i, j int) int {
	for ; i >= 0; i-- {
		if field[i][j] == Obstruction {
			i++
			field[i][j] = MarkCross
			break
		}
		if field[i][j] == MarkHoriz {
			field[i][j] = MarkCross
		} else {
			field[i][j] = MarkVert
		}
	}
	return i
}

func moveRight(field [][]rune, i, j int) int {
	j0 := j
	for ; j < len(field[i]); j++ {
		if field[i][j] == Obstruction {
			j--
			break
		}
		if j == j0 || field[i][j] == MarkVert {
			field[i][j] = MarkCross
		} else {
			field[i][j] = MarkHoriz
		}
	}
	return j
}

func moveDown(field [][]rune, i, j int) int {
	i0 := i
	for ; i < len(field); i++ {
		if field[i][j] == Obstruction {
			i--
			break
		}
		if i == i0 || field[i][j] == MarkHoriz {
			field[i][j] = MarkCross
		} else {
			field[i][j] = MarkVert
		}
	}
	return i
}

func moveLeft(field [][]rune, i, j int) int {
	j0 := j
	for ; j >= 0; j-- {
		if field[i][j] == Obstruction {
			j++
			break
		}
		if j == j0 || field[i][j] == MarkVert {
			field[i][j] = MarkCross
		} else {
			field[i][j] = MarkHoriz
		}
	}
	return j
}

func logField(field [][]rune) {
	for i := 0; i < len(field); i++ {
		log.Println(string(field[i]))
	}
}
