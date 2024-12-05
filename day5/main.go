package day5

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/samber/lo"
)

func Solve1() {
	data, err := os.ReadFile("day5/in1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	lines := strings.Split(string(data), "\n")
	separatorIndex := lo.IndexOf(lines, "")
	rules := lines[:separatorIndex]
	updates := lines[separatorIndex+1:]

	// Initial graph of rules
	graph := make(map[string]map[string]bool)
	for _, rule := range rules {
		rulePair := strings.Split(rule, "|")
		neighbors, ok := graph[rulePair[0]]
		if !ok {
			neighbors = make(map[string]bool)
			graph[rulePair[0]] = neighbors
		}
		neighbors[rulePair[1]] = true
	}

	// Find root nodes
	allNodes, endNodes := make(map[string]bool), make(map[string]bool)
	for _, rule := range rules {
		rulePair := strings.Split(rule, "|")
		allNodes[rulePair[0]] = true
		allNodes[rulePair[1]] = true
		endNodes[rulePair[1]] = true
	}
	roots := make(map[string]bool)
	for node := range allNodes {
		if _, ok := endNodes[node]; !ok {
			roots[node] = true
		}
	}

	// DFS to build a transitive closure
	visited := make(map[string]bool)
	var dfs func(node string)
	dfs = func(node string) {
		visited[node] = true
		for neighbor := range graph[node] {
			if !visited[neighbor] {
				dfs(neighbor)
			}
		}
		transNeighbors := make(map[string]bool)
		for neighbor := range graph[node] {
			transNeighbors[neighbor] = true
			for transNeighbor := range graph[neighbor] {
				transNeighbors[transNeighbor] = true
			}
		}
		graph[node] = transNeighbors
	}
	for root := range roots {
		if !visited[root] {
			dfs(root)
		}
	}

	// Validate updates
	validUpdates := make([][]string, 0)
	for _, update := range updates {
		nodes := strings.Split(update, ",")
		isValid := true
		for i := 0; i < len(nodes)-1; i++ {
			if _, ok := graph[nodes[i+1]][nodes[i]]; ok {
				isValid = false
				break
			}
		}
		if isValid {
			validUpdates = append(validUpdates, nodes)
		}
	}
	sum := 0
	for _, update := range validUpdates {
		if len(update)%2 == 0 {
			panic("Expecting the length of the update to be odd")
		}
		midNode := update[len(update)/2]
		value, _ := strconv.Atoi(midNode)
		sum += value
	}
	fmt.Println(sum)
}

func Solve2() {
}
