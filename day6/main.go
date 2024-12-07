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
	Marked      = 'X'
)

func Solve1() {
	log.SetOutput(os.Stdout)
	data, err := os.ReadFile("day6/in1.txt")
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
	i, j := i0, j0
	for 0 <= i && i < len(field) && 0 <= j && j < len(field[0]) {
		switch field[i][j] {
		case Up:
			i = moveUp(field, i, j)
			if i >= 0 {
				field[i][j] = Right
			}
		case Right:
			j = moveRight(field, i, j)
			if j < len(field[i]) {
				field[i][j] = Down
			}
		case Down:
			i = moveDown(field, i, j)
			if i < len(field) {
				field[i][j] = Left
			}
		case Left:
			j = moveLeft(field, i, j)
			if j >= 0 {
				field[i][j] = Up
			}
		default:
			panic("invalid direction")
		}
	}
	logField(field)

	markedCount := 0
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			if field[i][j] == Marked {
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
			break
		}
		field[i][j] = Marked
	}
	return i
}

func moveRight(field [][]rune, i, j int) int {
	for ; j < len(field[i]); j++ {
		if field[i][j] == Obstruction {
			j--
			break
		}
		field[i][j] = Marked
	}
	return j
}

func moveDown(field [][]rune, i, j int) int {
	for ; i < len(field); i++ {
		if field[i][j] == Obstruction {
			i--
			break
		}
		field[i][j] = Marked
	}
	return i
}

func moveLeft(field [][]rune, i, j int) int {
	for ; j >= 0; j-- {
		if field[i][j] == Obstruction {
			j++
			break
		}
		field[i][j] = Marked
	}
	return j
}

func logField(field [][]rune) {
	for i := 0; i < len(field); i++ {
		log.Println(string(field[i]))
	}
}
