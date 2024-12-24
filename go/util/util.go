package util

import (
	"fmt"
	"os"
	"strings"
)

func ReadInput(inputPath string) []string {
	data, err := os.ReadFile(inputPath)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return strings.Split(string(data), "\n")
}
