package main

import (
	"fmt"
	"os"
	"strconv"
)

type rotation int

const (
	rotationNo rotation = iota
	// 90 degrees
	rotationLeft
	// 180 degrees
	rotationDown
	// 270 degrees
	rotationRight
)

type faceDirection int

const (
	facePositiveX faceDirection = iota
	facePositiveY
	facePositiveZ
	faceNegativeX
	faceNegativeY
	faceNegativeZ
)

type beaconScanner struct {
	face    faceDirection
	orn     rotation
	beacons []point
}

// relative convert the p to the relative position (from scanner 0)
func (b *beaconScanner) relativePoint(orn rotation, face faceDirection, p point) point {

	// p.print()
	switch orn {
	case rotationNo:
	case rotationLeft:
		p.y, p.z = -p.z, p.y
	case rotationDown:
		p.y, p.z = -p.y, -p.z
	case rotationRight:
		p.y, p.z = p.z, -p.y
	}

	switch face {
	case facePositiveX:
	case faceNegativeX:
		p.x, p.y = -p.x, -p.y
	case facePositiveY:
		p.x, p.y = p.y, -p.x
	case faceNegativeY:
		p.x, p.y = -p.y, p.x
	case facePositiveZ:
		p.x, p.z = p.z, -p.x
	case faceNegativeZ:
		p.x, p.z = -p.z, p.x
	}

	return p
}

func (b *beaconScanner) turn(beacons []point, i int) []point {
	var np []point
	face := (b.face + faceDirection(i)) % 6
	for _, bc := range beacons {
		np = append(np, b.relativePoint(b.orn, face, bc))
	}
	return np
}

func (b *beaconScanner) rotate(beacons []point, i int) []point {
	var np []point
	orn := (b.orn + rotation(i)) % 4
	for _, bc := range beacons {
		np = append(np, b.relativePoint(orn, b.face, bc))
	}
	return np
}

var emptyPoiont = point{}

func (b *beaconScanner) distanceSet(beacons []point) map[vector]int {
	distanceSet := make(map[vector]int)

	for _, bp := range b.beacons {
		for _, p := range beacons {
			distanceSet[bp.distanceFrom(p)]++
		}
	}

	return distanceSet
}

// calibration calibrates the scanner with scanner 0(the given point)
// and returns the position of the beaconScanner (relative to scanner 0)
func (b *beaconScanner) calibration(beacons []point) point {
	offset := b.distanceSet(beacons)
	for k, v := range offset {
		if v >= 12 {
			return point(k)
		}
	}
	return emptyPoiont
}

type vector point

type point struct {
	x, y, z int
}

func (p point) distanceFrom(p2 point) vector {
	return vector{x: p.x - p2.x, y: p.y - p2.y, z: p.z - p2.z}
}

var scanners []point

func main() {
	data := readInput()
	scanner0 := &beaconScanner{
		face:    facePositiveX,
		orn:     rotationNo,
		beacons: data[0],
	}

	var done = []*beaconScanner{scanner0}
	var undone []*beaconScanner

	beacons := make(map[point]struct{})

	for _, b := range scanner0.beacons {
		beacons[b] = struct{}{}
	}

	for i, s := range data {
		if i == 0 {
			continue
		}
		newScanner := &beaconScanner{
			face:    facePositiveX,
			orn:     rotationNo,
			beacons: s,
		}
		undone = append(undone, newScanner)
	}

	for {
		if len(undone) == 0 {
			break
		}
	LOOP:
		for index, newScanner := range undone {
			for i := 0; i < 6; i++ {
				bs := newScanner.turn(newScanner.beacons, i)
				for j := 0; j < 4; j++ {
					ns := newScanner.rotate(bs, j)
					for _, scanner0 := range done {
						scannerPosition := scanner0.calibration(ns)
						if scannerPosition != emptyPoiont {
							scanners = append(scanners, scannerPosition)
							var afterCalibration []point
							for _, bp := range ns {
								p := point{
									x: bp.x + scannerPosition.x,
									y: bp.y + scannerPosition.y,
									z: bp.z + scannerPosition.z,
								}
								afterCalibration = append(afterCalibration, p)
								beacons[p] = struct{}{}
							}

							done = append(done, &beaconScanner{
								face:    faceDirection(i),
								orn:     rotation(j),
								beacons: afterCalibration,
							})
							// remove index from undone
							undone = append(undone[:index], undone[index+1:]...)
							goto LOOP
						}
					}
				}
			}
		}
	}

	fmt.Fprintf(os.Stdout, "p1: answer: %d\n", len(beacons))

	var max int

	for _, p := range scanners {
		for _, p2 := range scanners {
			if p.manhattan(p2) > max {
				max = p.manhattan(p2)
			}
		}
	}

	fmt.Fprintf(os.Stdout, "p2: max manhattan: %d\n", max)
}

func (p point) manhattan(p2 point) int {
	return abs(p.x-p2.x) + abs(p.y-p2.y) + abs(p.z-p2.z)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func (p point) print() {
	fmt.Printf("(%d, %d, %d)\n", p.x, p.y, p.z)
}

func atoi(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}
