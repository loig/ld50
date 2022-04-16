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

func (l *level) GenArea() {

	// gen a sequence of steps (up, down, left, right) with at least three steps
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

	// 0 up, 1 down
	upDowns := make([]int, numUpDown)
	upDowns[numUpDown-1] = 0 // always ends with up
	upDowns[0] = 0           // always starts with up
	for i := 1; i < numUpDown-1; i++ {
		upDowns[i] = rand.Intn(2)
	}

	// 2 left, 3 right
	leftRights := make([]int, numLeftRight)
	left := false  // must have at least one step left
	right := false // must have at least one step right
	if numLeftRight == 2 {
		leftRights[0] = rand.Intn(2)
		leftRights[1] = 1 - leftRights[0]
		leftRights[0] += 2
		leftRights[1] += 2
	}
	for i := 0; i < numLeftRight; i++ {
		if !left && i == numLeftRight-1 {
			leftRights[i] = 2
		} else if !right && i == numLeftRight-1 {
			leftRights[i] = 3
		} else {
			leftRights[i] = rand.Intn(2) + 2
			if leftRights[i] == 2 {
				left = true
			} else {
				right = true
			}
		}
	}

	// merge the step sequences
	steps := make([]int, numSteps)
	// if numUpDown > numLeftRight or if numUpDown == numLeftRight and random is 1
	first := upDowns
	second := leftRights
	if numUpDown < numLeftRight || (numUpDown == numLeftRight && rand.Intn(2) == 0) {
		first = leftRights
		second = upDowns
	}
	for i := 0; i < numSteps; i++ {
		log.Print(i, i%2, i/2, len(first), len(second))
		if i%2 == 0 {
			steps[i] = first[i/2]
		} else {
			steps[i] = second[i/2]
		}
	}

	log.Print(steps)

	// transform the step sequence in a sequence of positions on the grid

	// put obstacles according to the position sequence

}
