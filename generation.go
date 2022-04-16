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
	possibleXPos := make([]int, len(l.area[0])/3-1) // need to be parameterized
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
	possibleYPos := make([]int, len(l.area)/2-2) // need to be parameterized
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

	log.Print("everything: ", xyseq)

	// from positions, get coordinates

	// put obstacles according to the position sequence and the step sequence

	// check that the level is solvable

}
