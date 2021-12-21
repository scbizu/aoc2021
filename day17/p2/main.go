package main

import (
	"fmt"
	"os"
)

const (
	leftX   = 117
	rightX  = 164
	topY    = -89
	bottomY = -140
)

type point struct {
	x, y int
}

var pointsMap map[point]struct{} = make(map[point]struct{})

func main() {
	steps := make(map[int][]point)
	var x int
	for i := leftX; i <= rightX; i++ {
		for {
			if x > rightX {
				break
			}
			var xSum int
			for k := x; k > 0; k-- {
				xSum += k
				if xSum <= rightX && xSum >= leftX {
					step := x - k + 1
					if _, ok := steps[step]; !ok {
						steps[step] = []point{{x: x}}
					} else {
						steps[step] = append(steps[step], point{x: x})
					}
					break
				}
			}
			x++
		}
	}

	for step, points := range steps {
		for index := range points {
			for j := bottomY; j <= topY; j++ {
				step := step
				for {
					if step == 1 {
						pointsMap[point{x: points[index].x, y: j}] = struct{}{}
					} else {
						y := ((2*j)/step + step - 1) / 2
						if (y > 0 && -y-1 < bottomY) || sumX(points[index].x, step) > rightX {
							break
						}
						if bottomY <= sumY(y, step) && sumY(y, step) <= topY && sumX(points[index].x, step) >= leftX && sumX(points[index].x, step) <= rightX {
							pointsMap[point{x: points[index].x, y: y}] = struct{}{}
						}
					}
					step++
				}
			}
		}
	}

	fmt.Fprintf(os.Stdout, "pointsMap: %v\n", pointsMap)
	fmt.Fprintf(os.Stdout, "points: %v\n", len(pointsMap))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func sumY(y int, step int) int {
	var sum int

	dst := y - step + 1

	for i := y; i >= dst; i-- {
		sum += i
	}
	return sum
}

func sumX(x int, step int) int {
	var sum int
	dst := x - step + 1
	if dst < 0 {
		dst = 0
	}

	for i := x; i >= dst; i-- {
		sum += i
	}
	return sum
}
