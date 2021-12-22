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

func small(h string) bool {
	return h != "start" && h != "end" && strings.ToLower(h) == h
}

func visit(hoele string, path []string, connections map[string][]string, visited []string) int {
	if hoele == "start" {
		return 0
	}

	if small(hoele) {
		for _, v := range visited {
			if hoele == v {
				return 0
			}
		}
		visited = append(visited, hoele)
	}

	path = append(path, hoele)
	if hoele == "end" {
		fmt.Println(strings.Join(path, ","))
		return 1
	}

	paths := 0
	for _, next := range connections[hoele] {
		paths += visit(next, path, connections, visited)
	}

	return paths
}

func countPaths(connections map[string][]string) int {
	count := 0
	visited := []string{}
	path := []string{"start"}
	for _, next := range connections["start"] {
		count += visit(next, path, connections, visited)
	}
	return count
}

func part1() int {
	conns := loadConnections()
	fmt.Println(conns)
	return countPaths(conns)
}

func part2() int {
	return 0
}

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}
