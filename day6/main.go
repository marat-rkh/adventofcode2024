package day6

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/zyedidia/generic/mapset"

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
	direction0 := field[i0][j0]

	i, j := i0, j0
	direction := direction0
	trace := make(map[[2]int]mapset.Set[rune])
	for 0 <= i && i < len(field) && 0 <= j && j < len(field[0]) {
		switch direction {
		case Up:
			i = moveUp(field, i, j, trace)
			if i >= 0 {
				direction = Right
			}
		case Right:
			j = moveRight(field, i, j, trace)
			if j < len(field[i]) {
				direction = Down
			}
		case Down:
			i = moveDown(field, i, j, trace)
			if i < len(field) {
				direction = Left
			}
		case Left:
			j = moveLeft(field, i, j, trace)
			if j >= 0 {
				direction = Up
			}
		default:
			panic("invalid direction")
		}
	}
	logField(field, trace)

	visitedCount := 0
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			if _, ok := trace[[2]int{i, j}]; ok {
				visitedCount++
			}
		}
	}
	fmt.Println(visitedCount)
}

func moveUp(field [][]rune, i, j int, trace map[[2]int]mapset.Set[rune]) int {
	for ; i >= 0; i-- {
		if field[i][j] == Obstruction {
			i++
			break
		}
		getOrInit(trace, i, j).Put(Up)
	}
	return i
}

func moveRight(field [][]rune, i, j int, trace map[[2]int]mapset.Set[rune]) int {
	for ; j < len(field[i]); j++ {
		if field[i][j] == Obstruction {
			j--
			break
		}
		getOrInit(trace, i, j).Put(Right)
	}
	return j
}

func moveDown(field [][]rune, i, j int, trace map[[2]int]mapset.Set[rune]) int {
	for ; i < len(field); i++ {
		if field[i][j] == Obstruction {
			i--
			break
		}
		getOrInit(trace, i, j).Put(Down)
	}
	return i
}

func moveLeft(field [][]rune, i, j int, trace map[[2]int]mapset.Set[rune]) int {
	for ; j >= 0; j-- {
		if field[i][j] == Obstruction {
			j++
			break
		}
		getOrInit(trace, i, j).Put(Left)
	}
	return j
}

func getOrInit(trace map[[2]int]mapset.Set[rune], i, j int) mapset.Set[rune] {
	if _, ok := trace[[2]int{i, j}]; !ok {
		trace[[2]int{i, j}] = mapset.New[rune]()
	}
	return trace[[2]int{i, j}]
}

func logField(field [][]rune, trace map[[2]int]mapset.Set[rune]) {
	var fieldWithTrace strings.Builder
	fieldWithTrace.WriteRune('\n')
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			if field[i][j] == Empty {
				if trace, ok := trace[[2]int{i, j}]; ok {
					if (trace.Has(Up) || trace.Has(Down)) && (trace.Has(Left) || trace.Has(Right)) {
						fieldWithTrace.WriteRune(MarkCross)
					} else if trace.Has(Up) || trace.Has(Down) {
						fieldWithTrace.WriteRune(MarkVert)
					} else {
						// trace.Has(Left) || trace.Has(Right)
						fieldWithTrace.WriteRune(MarkHoriz)
					}
				} else {
					fieldWithTrace.WriteRune(field[i][j])
				}
			} else {
				fieldWithTrace.WriteRune(field[i][j])
			}
		}
		fieldWithTrace.WriteRune('\n')
	}
	log.Println(fieldWithTrace.String())
}
