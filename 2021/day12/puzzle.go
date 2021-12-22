package main

import (
	"fmt"
	"strings"

	"github.com/oscarh/aoc/util"
)

func loadConnections() map[string][]string {
	connections := map[string][]string{}
	for _, connection := range util.LoadInput() {
		startEnd := strings.Split(connection, "-")
		s := startEnd[0]
		e := startEnd[1]
		connections[s] = append(connections[s], e)
		connections[e] = append(connections[e], s)
	}
	return connections
}

func small(c string) bool {
	return c != "start" && c != "end" && strings.ToLower(c) == c
}

func visit(cave string, path []string, connections map[string][]string, visited []string, smallVisitedTwice bool) int {
	if cave == "start" {
		return 0
	}

	if small(cave) {
		for _, v := range visited {
			if cave == v {
				if smallVisitedTwice {
					return 0
				} else {
					smallVisitedTwice = true
				}
			}
		}
		visited = append(visited, cave)
	}

	path = append(path, cave)
	if cave == "end" {
		fmt.Println(strings.Join(path, ","))
		return 1
	}

	paths := 0
	for _, next := range connections[cave] {
		paths += visit(next, path, connections, visited, smallVisitedTwice)
	}

	return paths
}

func countPaths(connections map[string][]string, mayVisitSmallTwice bool) int {
	count := 0
	visited := []string{}
	path := []string{"start"}
	for _, next := range connections["start"] {
		count += visit(next, path, connections, visited, !mayVisitSmallTwice)
	}
	return count
}

func part1() int {
	conns := loadConnections()
	fmt.Println(conns)
	return countPaths(conns, false)
}

func part2() int {
	conns := loadConnections()
	fmt.Println(conns)
	return countPaths(conns, true)
}

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}
