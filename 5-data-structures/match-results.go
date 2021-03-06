/*
	Program to analyze sport match results and calculate the points of the given team.
	The match results are stored in an array called results.
	Each match has one of the following results:
		"w" - won
		"l" - lost
		"d" - draw

	A win adds three points, a draw adds one point, and a lost match does not add any points.
*/
package main

import "fmt"

func cmp(x string) int {
	var pts int
	switch x {
	case "w":
		pts = 3
	case "l":
		pts = 0
	case "d":
		pts = 1
	}
	return pts
}

func main() {
	results := []string{"w", "l", "w", "d", "w", "l", "l", "l", "d", "d", "w", "l", "w", "d"}
	var pts int = 0
	for _, res := range results {
		pts += cmp(res)
	}

	var n string
	fmt.Scanln(&n)
	pts += cmp(n)

	fmt.Println(pts)
}
