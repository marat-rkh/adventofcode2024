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

	Obstruction      = '#'
	Empty            = '.'
	MarkVert         = '|'
	MarkHoriz        = '-'
	MarkCross        = '+'
	AddedObstruction = 'O'
)

func Solve() {
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

	closestObstructions := calcClosestObstructions(field)
	i, j := i0, j0
	direction := direction0
	trace := make(map[[2]int]*mapset.Set[rune])
	addedObstructions := mapset.New[[2]int]()
	for 0 <= i && i < len(field) && 0 <= j && j < len(field[0]) {
		switch direction {
		case Up:
			i = moveUp(field, i, j, closestObstructions, trace, &addedObstructions)
			if i >= 0 {
				direction = Right
			}
		case Right:
			j = moveRight(field, i, j, closestObstructions, trace, &addedObstructions)
			if j < len(field[i]) {
				direction = Down
			}
		case Down:
			i = moveDown(field, i, j, closestObstructions, trace, &addedObstructions)
			if i < len(field) {
				direction = Left
			}
		case Left:
			j = moveLeft(field, i, j, closestObstructions, trace, &addedObstructions)
			if j >= 0 {
				direction = Up
			}
		default:
			panic("invalid direction")
		}
	}
	logField(field, trace, &addedObstructions)

	visitedCount := 0
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			if _, ok := trace[[2]int{i, j}]; ok {
				visitedCount++
			}
		}
	}
	fmt.Printf("Visited: %d\n", visitedCount)
	// Incorrect answers: 531 (too low)
	fmt.Printf("Added obstructions: %d\n", addedObstructions.Size())
}

type obstructions struct {
	above      [2]int
	toTheRight [2]int
	toTheLeft  [2]int
	below      [2]int
}

func calcClosestObstructions(field [][]rune) [][]*obstructions {
	closestObstructions := make([][]*obstructions, len(field))
	for i := 0; i < len(field); i++ {
		closestObstructions[i] = make([]*obstructions, len(field[i]))
		for j := 0; j < len(field[i]); j++ {
			closestObstructions[i][j] = &obstructions{}
		}
	}
	for i := 0; i < len(field); i++ {
		iObs, jObs := -1, -1
		for j := 0; j < len(field[i]); j++ {
			if field[i][j] == Obstruction {
				iObs, jObs = i, j
			}
			closestObstructions[i][j].toTheLeft = [2]int{iObs, jObs}
		}
	}
	for i := 0; i < len(field); i++ {
		iObs, jObs := -1, -1
		for j := len(field[i]) - 1; j >= 0; j-- {
			if field[i][j] == Obstruction {
				iObs, jObs = i, j
			}
			closestObstructions[i][j].toTheRight = [2]int{iObs, jObs}
		}
	}
	for j := 0; j < len(field[0]); j++ {
		iObs, jObs := -1, -1
		for i := 0; i < len(field); i++ {
			if field[i][j] == Obstruction {
				iObs, jObs = i, j
			}
			closestObstructions[i][j].above = [2]int{iObs, jObs}
		}
	}
	for j := 0; j < len(field[0]); j++ {
		iObs, jObs := -1, -1
		for i := len(field) - 1; i >= 0; i-- {
			if field[i][j] == Obstruction {
				iObs, jObs = i, j
			}
			closestObstructions[i][j].below = [2]int{iObs, jObs}
		}
	}
	return closestObstructions
}

func moveUp(field [][]rune, i, j int, closestObs [][]*obstructions, trace map[[2]int]*mapset.Set[rune], addedObs *mapset.Set[[2]int]) int {
	for ; i >= 0; i-- {
		if field[i][j] == Obstruction {
			i++
			break
		}
		if i-1 >= 0 && field[i-1][j] != Obstruction {
			obsToTheRight := closestObs[i][j].toTheRight
			iObs, jObs := obsToTheRight[0], obsToTheRight[1]
			if iObs != -1 && jObs != -1 {
				if marks, ok := trace[[2]int{iObs, jObs - 1}]; ok && marks.Has(Right) {
					addedObs.Put([2]int{i - 1, j})
				}
			}
		}
		getOrInit(trace, i, j).Put(Up)
	}
	return i
}

func moveRight(field [][]rune, i, j int, closestObs [][]*obstructions, trace map[[2]int]*mapset.Set[rune], addedObs *mapset.Set[[2]int]) int {
	for ; j < len(field[i]); j++ {
		if field[i][j] == Obstruction {
			j--
			break
		}
		if j+1 < len(field[i]) && field[i][j+1] != Obstruction {
			obsBelow := closestObs[i][j].below
			iObs, jObs := obsBelow[0], obsBelow[1]
			if iObs != -1 && jObs != -1 {
				if marks, ok := trace[[2]int{iObs - 1, jObs}]; ok && marks.Has(Down) {
					addedObs.Put([2]int{i, j + 1})
				}
			}
		}
		getOrInit(trace, i, j).Put(Right)
	}
	return j
}

func moveDown(field [][]rune, i, j int, closestObs [][]*obstructions, trace map[[2]int]*mapset.Set[rune], addedObs *mapset.Set[[2]int]) int {
	for ; i < len(field); i++ {
		if field[i][j] == Obstruction {
			i--
			break
		}
		if i+1 < len(field) && field[i+1][j] != Obstruction {
			obsToTheLeft := closestObs[i][j].toTheLeft
			iObs, jObs := obsToTheLeft[0], obsToTheLeft[1]
			if iObs != -1 && jObs != -1 {
				if marks, ok := trace[[2]int{iObs, jObs + 1}]; ok && marks.Has(Left) {
					addedObs.Put([2]int{i + 1, j})
				}
			}
		}
		getOrInit(trace, i, j).Put(Down)
	}
	return i
}

func moveLeft(field [][]rune, i, j int, closestObs [][]*obstructions, trace map[[2]int]*mapset.Set[rune], addedObs *mapset.Set[[2]int]) int {
	for ; j >= 0; j-- {
		if field[i][j] == Obstruction {
			j++
			break
		}
		if j-1 >= 0 && field[i][j-1] != Obstruction {
			obsAbove := closestObs[i][j].above
			iObs, jObs := obsAbove[0], obsAbove[1]
			if iObs != -1 && jObs != -1 {
				if marks, ok := trace[[2]int{iObs + 1, jObs}]; ok && marks.Has(Up) {
					addedObs.Put([2]int{i, j - 1})
				}
			}
		}
		getOrInit(trace, i, j).Put(Left)
	}
	return j
}

func getOrInit(trace map[[2]int]*mapset.Set[rune], i, j int) *mapset.Set[rune] {
	if _, ok := trace[[2]int{i, j}]; !ok {
		marks := mapset.New[rune]()
		trace[[2]int{i, j}] = &marks
	}
	return trace[[2]int{i, j}]
}

func logField(field [][]rune, trace map[[2]int]*mapset.Set[rune], addedObs *mapset.Set[[2]int]) {
	var fieldWithTrace strings.Builder
	fieldWithTrace.WriteRune('\n')
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[i]); j++ {
			if addedObs.Has([2]int{i, j}) {
				fieldWithTrace.WriteRune(AddedObstruction)
			} else if field[i][j] == Empty {
				if marks, ok := trace[[2]int{i, j}]; ok {
					if (marks.Has(Up) || marks.Has(Down)) && (marks.Has(Left) || marks.Has(Right)) {
						fieldWithTrace.WriteRune(MarkCross)
					} else if marks.Has(Up) || marks.Has(Down) {
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
