// Source : https://leetcode.com/problems/rectangle-area/
// Author : daoran
// Date   : 2015-04-16

/**********************************************************************************
 *
 * Find the total area covered by two rectilinear rectangles in a 2D plane.
 * Each rectangle is defined by its bottom left corner and top right corner as shown in the figure.
 *
 *                      Y
 *                      ^
 *                      |
 *                      |
 *                      |
 *                      |
 *                      |     (C,D):(3,4)
 *            +------------------+
 *            |         |        |
 *            |         |        |
 *            |         |        |               (G,H):(9,2)
 *            |         +----------------------------+
 *            |         |        |                   |
 *            |         |        |                   |
 *            |         |        |                   |
 *            +---------|--------+-------------------|---------> X
 *      (A,B):(-3,0)    |                            |
 *                      +----------------------------+
 *                  (E,F):(0,-1)
 *
 *
 *
 * Assume that the total area is never beyond the maximum possible value of int.
 *
 * Credits:Special thanks to @mithmatt for adding this problem, creating the above image and all test cases.
 *
 **********************************************************************************/

package main

import "fmt"

type point struct {
	x, y int
}

type rectangle struct {
	topLeft     point
	bottomRight point
}

func (r rectangle) area() int {
	return (r.bottomRight.x - r.topLeft.x) * (r.topLeft.y - r.bottomRight.y)
}

func (r rectangle) inclusiveArea(nr rectangle) int {
	// I include it
	if nr.topLeft.x >= r.topLeft.x && nr.bottomRight.x <= r.bottomRight.x && nr.topLeft.y <= r.topLeft.y && nr.bottomRight.y >= r.bottomRight.y {
		return r.area()
	}

	// it includes me
	if nr.topLeft.x <= r.topLeft.x && nr.bottomRight.x >= r.bottomRight.x && nr.topLeft.y >= r.topLeft.y && nr.bottomRight.y <= r.bottomRight.y {
		return nr.area()
	}

	// 0 : no includive
	return 0
}

func (r rectangle) overlappedArea(nr rectangle) int {
	overlapX := max(0, min(nr.bottomRight.x, r.bottomRight.x)-max(nr.topLeft.x, r.topLeft.x))
	overlapY := max(0, min(nr.topLeft.y, r.topLeft.y)-max(nr.bottomRight.y, r.bottomRight.y))

	return overlapX * overlapY
}

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	r1 := rectangle{topLeft: point{A, D}, bottomRight: point{C, B}}
	r2 := rectangle{topLeft: point{E, H}, bottomRight: point{G, F}}

	area := r1.inclusiveArea(r2)
	if area > 0 {
		return area
	}
	return r1.area() + r2.area() - r1.overlappedArea(r2)
}

func main() {
	// 16
	fmt.Println(computeArea(-1, -1, 1, 1, -2, -2, 2, 2))
	// 16
	fmt.Println(computeArea(-2, -2, 2, 2, -1, -1, 1, 1))
	// 17
	fmt.Println(computeArea(-2, -2, 2, 2, -4, 3, -3, 4))
	// 45
	fmt.Println(computeArea(-3, 0, 3, 4, 0, -1, 9, 2))
	// 24
	fmt.Println(computeArea(-2, -2, 2, 2, -3, -3, 3, -1))
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
