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

	graph := calcDirectDependenciesGraph(rules)
	validUpdates := make([][]string, 0)
	for _, update := range updates {
		nodes := strings.Split(update, ",")
		isValid := isValidUpdate(nodes, graph)
		if isValid {
			validUpdates = append(validUpdates, nodes)
		}
	}
	sum := sumMiddleNodes(validUpdates)
	fmt.Println(sum)
}

func calcDirectDependenciesGraph(rules []string) map[string]map[string]bool {
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
	return graph
}

func calcTransitiveClosureGraph(rules []string) map[string]map[string]bool {
	graph := calcDirectDependenciesGraph(rules)

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
	return graph
}

func isValidUpdate(update []string, relations map[string]map[string]bool) bool {
	for i := 0; i < len(update)-1; i++ {
		for j := i + 1; j < len(update); j++ {
			if !isGreaterByRelations(update[i], update[j], relations) {
				return false
			}
		}
	}
	return true
}

func isGreaterByRelations(x, y string, relations map[string]map[string]bool) bool {
	if _, ok := relations[y][x]; ok {
		return false
	}
	return true
}

func sumMiddleNodes(nodes [][]string) int {
	sum := 0
	for _, update := range nodes {
		if len(update)%2 == 0 {
			panic("Expecting the length of the update to be odd")
		}
		midNode := update[len(update)/2]
		value, _ := strconv.Atoi(midNode)
		sum += value
	}
	return sum
}

func Solve2() {
	data, err := os.ReadFile("day5/in1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	lines := strings.Split(string(data), "\n")
	separatorIndex := lo.IndexOf(lines, "")
	rules := lines[:separatorIndex]
	updates := lines[separatorIndex+1:]

	graph := calcDirectDependenciesGraph(rules)
	invalidUpdates := make([][]string, 0)
	for _, update := range updates {
		nodes := strings.Split(update, ",")
		isValid := isValidUpdate(nodes, graph)
		if !isValid {
			invalidUpdates = append(invalidUpdates, nodes)
		}
	}
	for _, invalidUpdate := range invalidUpdates {
		// Bubble sort in descending order
		for i := 0; i < len(invalidUpdate); i++ {
			for j := 0; j < len(invalidUpdate)-i-1; j++ {
				if !isGreaterByRelations(invalidUpdate[j], invalidUpdate[j+1], graph) {
					invalidUpdate[j], invalidUpdate[j+1] = invalidUpdate[j+1], invalidUpdate[j]
				}
			}
		}
	}
	sum := sumMiddleNodes(invalidUpdates)
	fmt.Println(sum)
}
