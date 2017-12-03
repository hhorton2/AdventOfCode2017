package main

import (
	"fmt"
	"math"
)

func main() {
	steps := 368078
	right := new(direction)
	up := new(direction)
	left := new(direction)
	down := new(direction)
	right.next = up
	right.count = 0
	up.next = left
	up.count = 0
	left.next = down
	left.count = 0
	down.next = right
	down.count = 0
	currentDirection := right
	partOne := solvePartOne(steps, currentDirection, right, up, left)
	currentDirection = right
	partTwo := solvePartTwo(steps, currentDirection, right, up, left)
	fmt.Println(partOne)
	fmt.Println(partTwo)
}

func solvePartOne(steps int, currentDirection *direction, right *direction, up *direction, left *direction) float64 {
	xCoord := 0
	yCoord := 0
	maxX := 0
	maxY := 0
	minX := 0
	minY := 0
	for i := 1; i < steps; i++ {
		if currentDirection == right {
			xCoord++
			if xCoord > maxX {
				currentDirection = currentDirection.next
				maxX = xCoord
			}
		} else if currentDirection == up {
			yCoord++
			if yCoord > maxY {
				currentDirection = currentDirection.next
				maxY = yCoord
			}
		} else if currentDirection == left {
			xCoord--
			if xCoord < minX {
				currentDirection = currentDirection.next
				minX = xCoord
			}
		} else {
			yCoord--
			if yCoord < minY {
				currentDirection = currentDirection.next
				minY = yCoord
			}
		}
	}
	return math.Abs(float64(xCoord)) + math.Abs(float64(yCoord))
}

func solvePartTwo(max int, currentDirection *direction, right *direction, up *direction, left *direction) int {
	var graph [100][100]int
	graph[50][50] = 1
	xCoord := 50
	yCoord := 50
	maxX := 50
	maxY := 50
	minX := 50
	minY := 50
	for {
		if currentDirection == right {
			xCoord++
			if xCoord > maxX {
				currentDirection = currentDirection.next
				maxX = xCoord
			}
		} else if currentDirection == up {
			yCoord++
			if yCoord > maxY {
				currentDirection = currentDirection.next
				maxY = yCoord
			}
		} else if currentDirection == left {
			xCoord--
			if xCoord < minX {
				currentDirection = currentDirection.next
				minX = xCoord
			}
		} else {
			yCoord--
			if yCoord < minY {
				currentDirection = currentDirection.next
				minY = yCoord
			}
		}
		sum := getNeighborSum(graph, xCoord, yCoord)
		if sum > max {
			return sum
		} else {
			graph[xCoord][yCoord] = sum
		}
	}
}

func getNeighborSum(graph [100][100]int, xCoord int, yCoord int) int {
	sum := 0
	sum += graph[xCoord-1][yCoord]
	sum += graph[xCoord+1][yCoord]
	sum += graph[xCoord][yCoord+1]
	sum += graph[xCoord][yCoord-1]
	sum += graph[xCoord-1][yCoord-1]
	sum += graph[xCoord-1][yCoord+1]
	sum += graph[xCoord+1][yCoord-1]
	sum += graph[xCoord+1][yCoord+1]
	return sum
}

type direction struct {
	count int
	next  *direction
}
