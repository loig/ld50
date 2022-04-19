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
	"github.com/hajimehoshi/ebiten/v2"
)

func (g *Game) UpdateLevelTransition() {
	g.transitionFrame++
	if g.transitionStep == 0 {
		if g.transitionFrame > globNumFramesTransStep0 {
			g.transitionStep = 1
			g.transitionFrame = 0
			return
		}
	}
	if g.transitionStep == 1 {
		if g.transitionFrame > globNumFramesTransStep1 {
			g.transitionStep = 2
			g.transitionFrame = 0
		}
	}
	if g.transitionStep == 2 {
		if g.transitionFrame > globNumFramesTransStep2 {
			g.transitionStep = 0
			g.transitionFrame = 0
			if g.inTuto {
				g.step = stepTuto
			} else {
				g.step = stepLevel
			}
		}
	}
}

func (g *Game) DrawLevelTransition(screen *ebiten.Image) {
	if g.transitionStep == 0 {
		g.previousLevel.Draw(screen, g.animeFrame, 0, 0, 1-float64(g.transitionFrame)/float64(globNumFramesTransStep0), true)
	} else if g.transitionStep == 1 {
		yshift := ((globLevelY - 1) * globAreaCellSize * g.transitionFrame) / globNumFramesTransStep1
		g.level.Draw(screen, g.animeFrame, 0, yshift-(globLevelY-1)*globAreaCellSize, 0, true)
		g.previousLevel.Draw(screen, g.animeFrame, 0, yshift, 0, true)
	} else {
		g.level.Draw(screen, g.animeFrame, 0, 0, 1*float64(g.transitionFrame)/float64(globNumFramesTransStep2), true)
	}
}
