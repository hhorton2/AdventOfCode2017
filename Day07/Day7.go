package main

import (
	"time"
	"io/ioutil"
	"strings"
	"fmt"
	"strconv"
)

func main() {
	start := time.Now()
	dat, err := ioutil.ReadFile("Day07/input.txt")
	check(err)
	nodes := parseInput(string(dat))
	fmt.Printf("%v\n%v", nodes, time.Since(start))
}
func parseInput(input string) []node {
	nodes := make([]node, 10)
	leafNodeMap := make(map[string]*node)
	for _, line := range strings.Split(input, "\n") {
		instructionSplit := strings.Split(line, ") -> ")
		if len(instructionSplit) > 1 {
			var root node
			root.name = strings.Split(instructionSplit[0], " (")[0]
			root.weight, _ = strconv.Atoi(strings.Split(instructionSplit[0], " (")[1])
			root.above = make([]*node, 0)
			root.below = make([]*node, 0)
			for _,aboveNode := range strings.Split(instructionSplit[1], ", "){
				if _,present := leafNodeMap[aboveNode]; present{
					root.above = append(root.above, leafNodeMap[aboveNode])
					delete(leafNodeMap, root.name)
				}

			}
		} else {
			var leaf node
			leaf.name = strings.Split(line, " ")[0]
			leaf.weight, _ = strconv.Atoi(strings.Replace(strings.Split(line, " (")[1], ")", "", -1))
			leaf.above = make([]*node, 0)
			leaf.below = make([]*node, 0)
			leafNodeMap[leaf.name] = &leaf
		}
	}

	return nodes
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type node struct {
	name   string
	weight int
	above  []*node
	below  []*node
}
