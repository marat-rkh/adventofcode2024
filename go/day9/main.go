package day9

import (
	"fmt"

	"github.com/marat-rkh/adventofcode2024/util"
	"github.com/zyedidia/generic/list"
	"github.com/zyedidia/generic/mapset"
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

func Solve2() {
	input := ReadInput("day9/in1.txt")
	res := DoSolve2(input)
	fmt.Println(res)
}

func DoSolve2(input []rune) int {
	moved := make(map[int]*list.List[[2]int], len(input)/2)
	allMoved := mapset.New[int]()
	lastFileIndex := len(input) - 1
	if len(input)%2 == 0 {
		lastFileIndex--
	}
	for i := lastFileIndex; i >= 0; i -= 2 {
		fileSize := int(input[i] - '0')
		fileId := i / 2
		for j := 1; j < i; j += 2 {
			empSpaceSize := int(input[j] - '0')
			if empSpaceSize >= fileSize {
				if _, ok := moved[j]; !ok {
					moved[j] = list.New[[2]int]()
				}
				moved[j].PushBack([2]int{fileId, fileSize})
				input[j] = rune(empSpaceSize - fileSize + '0')
				allMoved.Put(fileId)
				break
			}
		}
	}
	res := 0
	resIndex := 0
	for i := 0; i < len(input); i++ {
		if i%2 == 0 {
			fileSize := int(input[i] - '0')
			fileId := i / 2
			isMoved := allMoved.Has(fileId)
			for j := 0; j < fileSize; j++ {
				if !isMoved {
					res += resIndex * fileId
				}
				resIndex++
			}
		} else {
			if movedList, ok := moved[i]; ok {
				movedList.Front.Each(func(fileInfo [2]int) {
					fileId := fileInfo[0]
					fileSize := fileInfo[1]
					for j := 0; j < fileSize; j++ {
						res += resIndex * fileId
						resIndex++
					}
				})
			}
			empSpaceSize := int(input[i] - '0')
			for j := 0; j < empSpaceSize; j++ {
				resIndex++
			}
		}
	}
	return res
}

func DoSolve2Faster(input []rune) int {
	// O (n log n) solution
	// 1. We put all empty blocks into a binary search tree. Besides the size of the block,
	// we also store its index in the original array.
	// 2. For each element in the tree, we calculate the element with the minimum index in its subtree.
	// This can be done in one depth-first traversal.
	// 3. When we need to find the leftmost empty block for a file, we can find the subtree in our tree
	// where all empty blocks are greater than or equal to the size of the file in logarithmic time.
	// We also check which element in this subtree has the minimum index (the leftmost).
	// 4. Then, we remove this element from the tree and insert a new one with a size equal to "original size - file size",
	// while the index remains the same. Insertion into the binary tree is done in O(log n).
	// Restoring the minimum elements in the subtrees after insertion can be done in one pass
	// through all the ancestors of the inserted element, which is also O(log n).
	// 5. To implemet this we need a custom self balancing binary search tree.
	return 0
}
