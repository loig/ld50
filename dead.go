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

func (g *Game) UpdateDead() {
	if g.transitionStep < 1 {
		g.transitionFrame++
		if g.transitionFrame > globDeadStep0Frames {
			g.transitionFrame = 0
			g.transitionStep++
		}
	} else if g.transitionStep < 2 {
		g.transitionFrame++
		if g.transitionFrame > globDeadStep1Frames {
			g.transitionFrame = 0
			g.transitionStep++
		}
	} else {
		if isAnyKeyJustPressed() {
			g.transitionStep = 0
			g.transitionFrame = 0
			g.step = stepTitle
			if g.inTuto {
				g.step = stepTuto
				g.subStep = tutoStepBase
				if g.hud.levelNum != tutoLearnDeathLevel {
					g.hud.levelNum--
					g.hud.life = g.hud.lifeMax
				}
				g.NextLevel(false, true)
				return
			}
			g.Reset()
		}
	}
}

func (g *Game) DrawDead(screen *ebiten.Image) {

	if g.transitionStep == 0 {
		g.level.Draw(screen, g.animeFrame, 0, 0, 1, false, true)
	}

	if g.transitionStep == 1 {
		g.level.Draw(screen, g.animeFrame, 0, 0, 1-float64(g.transitionFrame)/float64(globDeadStep1Frames), false, true)
	}

	if g.transitionStep == 2 {
		g.level.Draw(screen, g.animeFrame, 0, 0, 0, false, true)
	}

	g.hud.Draw(screen, g.inTuto, true, g.transitionStep == 2)

	/*
		ebitenutil.DebugPrintAt(screen, "You died!", 30, 30)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("(at level ", g.hud.levelNum, ")"), 20, 45)
	*/

}
