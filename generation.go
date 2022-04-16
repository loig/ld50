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
	"log"
	"math/rand"
)

type posCuple struct {
	lr, ud int
}

func (l *level) GenArea() {

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
	log.Print(
		"steps: ", numSteps,
		"\nleft-right: ", numLeftRight,
		"\nup-down: ", numUpDown,
	)

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
	log.Print("left-right: ", xseq)

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
	log.Print("up-down: ", yseq)

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

	log.Print("everything: ", xyseq)

	// from positions, get coordinates
	xshift := rand.Intn(globXDivider)
	lastx := -1
	lasty := -1
	yshift := rand.Intn(globYDivider)
	for i := 0; i < len(xyseq); i++ {
		x := xyseq[i].lr
		if x == lastx {
			xyseq[i].lr = xyseq[i-1].lr
		} else {
			lastx = x
			if x == (len(l.area[0])/globXDivider-1)/2 {
				x = len(l.area[0]) / 2
				xshift = rand.Intn(globXDivider)
			} else {
				x = x*globXDivider + xshift
				if xshift == globXDivider-1 {
					xshift = rand.Intn(globXDivider-1) + 1
				} else {
					xshift = rand.Intn(globXDivider)
				}
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
				y = y*globYDivider + yshift
				if yshift == globYDivider-1 {
					yshift = rand.Intn(globYDivider-1) + 1
				} else {
					yshift = rand.Intn(globYDivider)
				}
			}
			xyseq[i].ud = y
		}
	}

	log.Print("real coordinates: ", xyseq)

	// put cactus according to the position sequence and the step sequence
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

		l.area[xyseq[i].ud][xyseq[i].lr] = &levelElement{
			elementType: foodType,
			posX:        xyseq[i].lr,
			posY:        xyseq[i].ud,
		}

		x := xyseq[i].lr + xmod
		y := xyseq[i].ud + ymod
		if x >= 0 && x < len(l.area[0]) && y >= 0 && y < len(l.area) {
			l.area[y][x] = &levelElement{
				elementType: cactusType,
				posX:        x,
				posY:        y,
			}
		}
	}

	// check that the level is solvable

}
