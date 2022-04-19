/*LD50, a game for Ludum Dare 50
  Copyright (C) 2022  Loïg Jezequel

  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.

  You should have received a copy of the GNU General Public License
  along with this program.  If not, see https://www.gnu.org/licenses/.
*/
package main

import (
	//"log"
	"math/rand"
)

type posCuple struct {
	lr, ud int
}

func (l *level) GenArea(withSnakes, withScorpions, withFood, withWater bool) {

	// choose a number of steps
	numSteps := rand.Intn(globNumSteps-2) + 3
	numUpDown := numSteps / 2
	numLeftRight := numUpDown
	if numUpDown+numLeftRight < numSteps {
		if numLeftRight < 2 {
			numLeftRight++
		} else {
			if rand.Intn(2) == 0 {
				numLeftRight++
			} else {
				numUpDown++
			}
		}
	}
	//log.Print(
	//	"steps: ", numSteps,
	//	"\nleft-right: ", numLeftRight,
	//	"\nup-down: ", numUpDown,
	//)

	// gen a sequence of left/right positions according to the number of steps
	// positions are not coordinates, they represent sets of x coordinates
	possibleXPos := make([]int, len(l.area[0])/globXDivider-1) // need to be parameterized
	for i := 0; i < len(possibleXPos); i++ {
		if i < len(possibleXPos)/2 {
			possibleXPos[i] = i
		} else {
			possibleXPos[i] = i + 1
		}
	}

	xseq := make([]int, numLeftRight+1)
	xseq[0] = len(possibleXPos) / 2
	for i := 1; i < len(xseq)-1; i++ {
		id := rand.Intn(len(possibleXPos))
		xseq[i] = possibleXPos[id]
		possibleXPos = append(possibleXPos[:id], possibleXPos[id+1:]...)
	}
	xseq[numLeftRight] = xseq[0]
	//log.Print("left-right: ", xseq)

	// gen a sequence of up/down coordinates according to the number of steps
	// positions are not coordinates, they represent sets of y coordinates
	possibleYPos := make([]int, len(l.area)/globYDivider-2) // need to be parameterized
	for i := 0; i < len(possibleYPos); i++ {
		possibleYPos[i] = i + 1
	}

	yseq := make([]int, numUpDown+1)
	yseq[0] = len(possibleYPos) + 1
	for i := 1; i < len(yseq)-1; i++ {
		id := rand.Intn(len(possibleYPos))
		yseq[i] = possibleYPos[id]
		possibleYPos = append(possibleYPos[:id], possibleYPos[id+1:]...)
	}
	yseq[numUpDown] = 0
	//log.Print("up-down: ", yseq)

	// merge the two sequences of positions to get couples of positions (x,y)
	xyseq := make([]posCuple, numSteps+1)

	if numUpDown > numLeftRight || (numUpDown == numLeftRight && rand.Intn(2) == 0) {
		for i := 0; i <= numSteps; i++ {
			xyseq[i] = posCuple{
				lr: xseq[i/2],
				ud: yseq[(i+1)/2],
			}
		}
	} else {
		for i := 0; i <= numSteps; i++ {
			xyseq[i] = posCuple{
				lr: xseq[(i+1)/2],
				ud: yseq[i/2],
			}
		}
	}

	// if the second last position is below the second one, do a symmetry
	if xyseq[1].lr == xyseq[len(xyseq)-2].lr && xyseq[1].ud < xyseq[len(xyseq)-2].ud {
		for i := 1; i < (len(xyseq)+1)/2; i++ {
			xyseq[i].ud, xyseq[len(xyseq)-i-1].ud = xyseq[len(xyseq)-i-1].ud, xyseq[i].ud
		}
	}

	//log.Print("everything: ", xyseq)

	// from positions, get coordinates
	lastx := -1
	lasty := -1
	for i := 0; i < len(xyseq); i++ {
		x := xyseq[i].lr
		if x == lastx {
			xyseq[i].lr = xyseq[i-1].lr
		} else {
			lastx = x
			if x == (len(l.area[0])/globXDivider-1)/2 {
				x = len(l.area[0]) / 2
			} else if x > lastx {
				x = x*globXDivider + rand.Intn(globXDivider-1)
			} else {
				x = x*globXDivider + 1 + rand.Intn(globXDivider-1)
			}
			xyseq[i].lr = x
		}

		y := xyseq[i].ud
		if y == lasty {
			xyseq[i].ud = xyseq[i-1].ud
		} else {
			lasty = y
			if y == yseq[0] {
				y = len(l.area) - 1
			} else if y != 0 {
				if y > lasty {
					y = y*globYDivider + rand.Intn(globYDivider-1)
				} else {
					y = y*globYDivider + 1 + rand.Intn(globYDivider-1)
				}
			}
			xyseq[i].ud = y
		}
	}

	//log.Print("real coordinates: ", xyseq)

	// put cactus (or water or food) according to the position sequence and the step sequence
	numWater := rand.Intn(globNumWater) + 1
	numFood := rand.Intn(globNumFood) + 1
	for i := 1; i < len(xyseq); i++ {
		xmod := 0
		xdiff := xyseq[i].lr - xyseq[i-1].lr
		if xdiff > 0 {
			xmod = 1
		} else if xdiff < 0 {
			xmod = -1
		}

		ymod := 0
		ydiff := xyseq[i].ud - xyseq[i-1].ud
		if ydiff > 0 {
			ymod = 1
		} else if ydiff < 0 {
			ymod = -1
		}

		isWater := false
		isFood := false

		if xyseq[i].lr != len(l.area[0])/2 || (xyseq[i].ud != 0 && xyseq[i].ud != len(l.area)-1) {
			if withWater && numWater > 0 && rand.Intn(globProbaWaterOnPath) == 0 {
				isWater = true
				numWater--
			}

			if !isWater && withFood && numFood > 0 && rand.Intn(globProbaFoodOnPath) == 0 {
				isFood = true
				numFood--
			}
		}

		if isWater || isFood {
			xmod = 0
			ymod = 0
		}

		x := xyseq[i].lr + xmod
		y := xyseq[i].ud + ymod
		if x >= 0 && x < len(l.area[0]) && y >= 0 && y < len(l.area) {
			l.area[y][x] = &levelElement{
				elementType: cactusType,
				posX:        x,
				posY:        y,
			}

			if isWater {
				l.area[y][x].elementType = waterType
			}

			if isFood {
				l.area[y][x].elementType = foodType
			}
		}
	}

	// get the expected path
	path := make([][]bool, len(l.area))
	pathlen := -2
	for i := 0; i < len(path); i++ {
		path[i] = make([]bool, len(l.area[0]))
	}

	for i := 1; i < len(xyseq); i++ {
		if xyseq[i].ud != xyseq[i-1].ud {
			start := xyseq[i].ud
			end := xyseq[i-1].ud
			if start > end {
				start, end = end, start
			}
			for j := start; j <= end; j++ {
				path[j][xyseq[i].lr] = true
				pathlen++
			}
		} else {
			start := xyseq[i].lr
			end := xyseq[i-1].lr
			if start > end {
				start, end = end, start
			}
			for j := start; j <= end; j++ {
				path[xyseq[i].ud][j] = true
				pathlen++
			}
		}
	}

	// add cactus on center path if needed
	hasCactus := false
	for y := 0; y < len(l.area); y++ {
		hasCactus = hasCactus || (l.area[y][len(l.area[y])/2] != nil) && (l.area[y][len(l.area[y])/2].elementType == cactusType)
	}
	if !hasCactus {
		cacRand := rand.Intn(len(l.area))
		cacPos := 0
		for cacRand >= 0 {
			if !path[cacPos][len(l.area[cacPos])/2] {
				cacRand--
			}
			if cacRand >= 0 {
				cacPos = (cacPos + 1) % len(l.area)
			}
		}
		l.area[cacPos][len(l.area[cacPos])/2] = &levelElement{
			elementType: cactusType,
			posX:        len(l.area[cacPos]) / 2,
			posY:        cacPos,
		}
	}

	// add cactus on left path if needed
	hasCactus = false
	possiblePos := make([]posCuple, 0)
	x := len(l.area[0]) / 2
	mid := x
	y := 0
	for x != mid || y != len(l.area)-1 {
		if y == 0 && x != 0 {
			x--
			hasCactus = hasCactus || (l.area[y][x] != nil && l.area[y][x].elementType == cactusType)
			if !path[y][x] {
				possiblePos = append(possiblePos, posCuple{lr: x, ud: y})
			}
			continue
		}

		if x == 0 && y != len(l.area)-1 {
			y++
			hasCactus = hasCactus || (l.area[y][x] != nil && l.area[y][x].elementType == cactusType)
			if !path[y][x] {
				possiblePos = append(possiblePos, posCuple{lr: x, ud: y})
			}
			continue
		}

		x++
		hasCactus = hasCactus || (l.area[y][x] != nil && l.area[y][x].elementType == cactusType)
		if !path[y][x] {
			possiblePos = append(possiblePos, posCuple{lr: x, ud: y})
		}
	}

	//log.Print(hasCactus, possiblePos)

	if !hasCactus && len(possiblePos) > 0 {
		choice := rand.Intn(len(possiblePos))
		l.area[possiblePos[choice].ud][possiblePos[choice].lr] = &levelElement{
			elementType: cactusType,
			posX:        possiblePos[choice].lr,
			posY:        possiblePos[choice].ud,
		}
	}

	// add cactus on right path if needed
	hasCactus = false
	possiblePos = make([]posCuple, 0)
	x = len(l.area[0]) / 2
	y = 0
	for x != mid || y != len(l.area)-1 {
		if y == 0 && x != len(l.area[0])-1 {
			x++
			hasCactus = hasCactus || (l.area[y][x] != nil && l.area[y][x].elementType == cactusType)
			if !path[y][x] {
				possiblePos = append(possiblePos, posCuple{lr: x, ud: y})
			}
			continue
		}

		if x == len(l.area[0])-1 && y != len(l.area)-1 {
			y++
			hasCactus = hasCactus || (l.area[y][x] != nil && l.area[y][x].elementType == cactusType)
			if !path[y][x] {
				possiblePos = append(possiblePos, posCuple{lr: x, ud: y})
			}
			continue
		}

		x--
		hasCactus = hasCactus || (l.area[y][x] != nil && l.area[y][x].elementType == cactusType)
		if !path[y][x] {
			possiblePos = append(possiblePos, posCuple{lr: x, ud: y})
		}
	}

	//log.Print(hasCactus, possiblePos)

	if !hasCactus && len(possiblePos) > 0 {
		choice := rand.Intn(len(possiblePos))
		l.area[possiblePos[choice].ud][possiblePos[choice].lr] = &levelElement{
			elementType: cactusType,
			posX:        possiblePos[choice].lr,
			posY:        possiblePos[choice].ud,
		}
	}

	// add a few more cactus
	isFree := func(i, j int) bool {
		if l.area[i][j] != nil || path[i][j] {
			return false
		}

		if i > 0 && l.area[i-1][j] != nil {
			return false
		}

		if j > 0 && l.area[i][j-1] != nil {
			return false
		}

		if i < len(l.area)-2 && l.area[i+1][j] != nil {
			return false
		}

		if j < len(l.area[0])-2 && l.area[i][j+1] != nil {
			return false
		}

		return true
	}

	numCactus := 0
	numFree := 0
	for i := 0; i < len(l.area); i++ {
		for j := 0; j < len(l.area); j++ {
			if l.area[i][j] != nil && l.area[i][j].elementType == cactusType {
				numCactus++
			} else if isFree(i, j) {
				numFree++
			}
		}
	}

	toAdd := rand.Intn(globNumCactus) + 1 - numCactus
	if toAdd < numFree {
		for toAdd > 0 {
			addPos := rand.Intn(numFree)
		OneCactusLoop:
			for i := 0; i < len(l.area); i++ {
				for j := 0; j < len(l.area[0]); j++ {
					if isFree(i, j) {
						addPos--
						if addPos < 0 {
							l.area[i][j] = &levelElement{
								elementType: cactusType,
								posX:        j,
								posY:        i,
							}
							break OneCactusLoop
						}
					}
				}
			}
			numFree--
			toAdd--
			numCactus++
		}
	}

	// add more water and food if needed
	if withWater && numWater > 0 {
		for numWater > 0 && numFree > 0 {
			pos := rand.Intn(numFree + pathlen)
		OneWaterLoop:
			for i := 0; i < len(path); i++ {
				for j := 0; j < len(path[0]); j++ {
					if (path[i][j] || isFree(i, j)) && (j != mid || (i != 0 && i != len(path)-1)) {
						if pos == 0 {
							if isFree(i, j) {
								numFree--
							}
							l.area[i][j] = &levelElement{
								elementType: waterType,
								posX:        j,
								posY:        i,
							}
							numWater--
							break OneWaterLoop
						}
						pos--
					}
				}
			}
		}
	}

	if withFood && numFood > 0 {
		for numFood > 0 && numFree > 0 {
			pos := rand.Intn(numFree + pathlen)
		OneFoodLoop:
			for i := 0; i < len(path); i++ {
				for j := 0; j < len(path[0]); j++ {
					if (path[i][j] || isFree(i, j)) && (j != mid || (i != 0 && i != len(path)-1)) {
						if pos == 0 {
							if isFree(i, j) {
								numFree--
							}
							l.area[i][j] = &levelElement{
								elementType: foodType,
								posX:        j,
								posY:        i,
							}
							numFood--
							break OneFoodLoop
						}
						pos--
					}
				}
			}
		}
	}

	maxNumMoving := globNumMoving

	// add a few scorpions
	if withScorpions {
		numScorpions := rand.Intn(globNumScorpions + 1)
		if !withSnakes && numScorpions == 0 {
			numScorpions = 1
		}
		if numScorpions > maxNumMoving {
			numScorpions = maxNumMoving
		}
		for numScorpions > 0 {
			if rand.Intn(2) == 0 {
				// on path
				pos := rand.Intn(pathlen)
			OneScorpionLoopA:
				for i := 0; i < len(path); i++ {
					for j := 0; j < len(path[0]); j++ {
						if path[i][j] && (j != mid || (i != 0 && i != len(path)-1)) {
							if pos <= 0 && l.area[i][j] == nil {
								l.area[i][j] = &levelElement{
									elementType: scorpionType,
									posX:        j,
									posY:        i,
								}
								maxNumMoving--
								break OneScorpionLoopA
							}
							pos--
						}
					}
				}
			} else {
				// not necessarily on path
				pos := rand.Intn(pathlen + numFree)
			OneScorpionLoopB:
				for i := 0; i < len(path); i++ {
					for j := 0; j < len(path[0]); j++ {
						if (path[i][j] || isFree(i, j)) && (j != mid || (i != 0 && i != len(path)-1)) {
							if pos == 0 && l.area[i][j] == nil {
								if isFree(i, j) {
									numFree--
								}
								l.area[i][j] = &levelElement{
									elementType: scorpionType,
									posX:        j,
									posY:        i,
								}
								maxNumMoving--
								break OneScorpionLoopB
							}
							pos--
						}
					}
				}
			}
			numScorpions--
		}
	}

	// add a few snakes to replace some cactus
	if withSnakes {
		numSnakes := rand.Intn(globNumSnakes) + 1
		if numSnakes > maxNumMoving {
			numSnakes = maxNumMoving
		}
		for numSnakes > 0 {
			pos := rand.Intn(numCactus)
		OneSnakeLoop:
			for i := 0; i < len(path); i++ {
				for j := 0; j < len(path[0]); j++ {
					if l.area[i][j] != nil && l.area[i][j].elementType == cactusType {
						if pos <= 0 &&
							j != mid && j != 0 && j != len(l.area[0])-1 &&
							((i != 0 && i != len(l.area)-1) || j == mid-1 || j == mid+1) {
							l.area[i][j].elementType = snakeType
							numCactus--
							break OneSnakeLoop
						}
						pos--
					}
				}
			}
			numSnakes--
		}
	}

	// move snakes
	for i := 0; i < len(l.area); i++ {
		for j := 0; j < len(l.area[i]); j++ {
			if l.area[i][j] != nil && l.area[i][j].elementType == snakeType && !l.area[i][j].doNotMoveInGeneration {
				directions := make([]int, 0) // 0 up, 1 right, 2 down, 3 left
				// possible to move snake down or up
				if i > 0 && i < len(l.area)-1 {
					if l.area[i-1][j] == nil && l.area[i+1][j] == nil {
						// down
						if !path[i-1][j] {
							directions = append(directions, 2)
						}
						// up
						if !path[i+1][j] {
							directions = append(directions, 0)
						}
					}
				}
				// possible to move snake left or right
				if j > 0 && j < len(l.area[0])-1 {
					if l.area[i][j-1] == nil && l.area[i][j+1] == nil {
						// left
						if !path[i][j+1] {
							directions = append(directions, 3)
						}
						if !path[i][j-1] {
							directions = append(directions, 1)
						}
					}
				}

				// choose direction
				if len(directions) > 0 {
					dirId := rand.Intn(len(directions))
					xmod := 0
					ymod := 0
					switch directions[dirId] {
					case 0:
						ymod = -1
					case 1:
						xmod = 1
					case 2:
						ymod = 1
					case 3:
						xmod = -1
					}
					// put a cactus to stop the snake
					l.area[i-ymod][j-xmod] = &levelElement{
						elementType: cactusType,
						posX:        j - xmod,
						posY:        i - ymod,
					}
					// move the snake
					x := j + xmod
					y := i + ymod
					for rand.Intn(globProbaMoveSnake) == 0 &&
						x+xmod >= 0 && x+xmod < len(l.area[0]) &&
						y+ymod >= 0 && y+ymod < len(l.area) &&
						l.area[y+ymod][x+xmod] == nil {
						x += xmod
						y += ymod
					}
					l.area[y][x] = l.area[i][j]
					l.area[y][x].posX = x
					l.area[y][x].posY = y
					l.area[y][x].doNotMoveInGeneration = true
					l.area[i][j] = nil
				}
			}
		}
	}

	// for debug only
	/*
		for i := 0; i < len(l.area); i++ {
			for j := 0; j < len(l.area[0]); j++ {
				if path[i][j] {
					if l.area[i][j] == nil {
						l.area[i][j] = &levelElement{
							elementType: nilType,
							posX:        j,
							posY:        i,
						}
					} else {
						log.Print("Warning, found ", l.area[i][j], " on path")
					}
				}
			}
		}
	*/

}
