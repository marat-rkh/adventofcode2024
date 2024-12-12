package day9

import (
	"fmt"

	"github.com/marat-rkh/adventofcode2024/util"
)

func Solve1() {
	input := ReadInput("day9/in1.txt")
	res := DoSolve1(input)
	fmt.Println(res)
}

func ReadInput(file string) []rune {
	return []rune(util.ReadInput(file)[0])
}

func DoSolve1(input []rune) int {
	res := 0
	resIndex := 0
	lastFileIndex := len(input) - 1
	if len(input)%2 == 0 {
		lastFileIndex--
	}
	for i := 0; i < len(input); i++ {
		if i%2 == 0 {
			fileSize := int(input[i] - '0')
			fileId := i / 2
			for j := 0; j < fileSize; j++ {
				res += resIndex * fileId
				resIndex++
			}
		} else {
			empSpaceSize := int(input[i] - '0')
			for empSpaceSize > 0 && lastFileIndex > i {
				lastFileSize := int(input[lastFileIndex] - '0')
				lastFileId := lastFileIndex / 2
				moveSize := min(empSpaceSize, lastFileSize)
				for j := 0; j < moveSize; j++ {
					res += resIndex * lastFileId
					resIndex++
				}
				empSpaceSize -= moveSize
				input[i] = rune(empSpaceSize + '0') // only for debug
				lastFileSize -= moveSize
				input[lastFileIndex] = rune(lastFileSize + '0')
				if lastFileSize == 0 {
					lastFileIndex -= 2
				}
			}
		}
	}
	return res
}
