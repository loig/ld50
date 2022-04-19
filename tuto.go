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
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"math/rand"
)

const (
	tutoStepBase int = iota
	tutoStepDead
	tutoStepDone
)

const (
	tutoLearnDeathLevel int = 8
	tutoLearnTabLevel   int = 4
)

var tutoSteps [][]string = [][]string{
	{"Move with arrows", "cactus block you"},
	{"Scorpions block", "and hurt you"},
	{"Watermelon", "heals you"},
	{"Use tab", "then move others"},
	{"Do not fear", "moving scorpions"},
	{"Snakes block", "and don't hurt"},
	{"But moving snakes", "will hurt you"},
	{"Keep an eye", "on your water"},
	{"Water bottles", "will help"},
	{"With space you", "skip levels"},
}

func (g *Game) UpdateTuto() {
	if g.subStep == tutoStepBase {
		hurt, food, water, finished, skip, fromX, toX, fromY, toY, hasMoved := g.level.Update(g.hud.levelNum >= tutoLearnTabLevel)
		if skip && g.hud.levelNum == len(tutoSteps) {
			g.subStep = tutoStepDone
		}
		if hurt {
			g.cameraShake = true
		}
		if hasMoved {
			g.AddParticlesOnGrid(fromX, fromY, toX, toY)
		}
		if water {
			g.AddWaterFoodParticles(toX, toY, true)
		}
		if food {
			g.AddWaterFoodParticles(toX, toY, false)
		}
		dead := g.hud.Update(hurt, food, water, g.hud.levelNum < tutoLearnDeathLevel)
		if dead {
			g.subStep = tutoStepDead
		}
		if finished {
			g.NextLevel(false, true)
			g.step = stepLevelTransition
		}
		return
	}

	if g.subStep == tutoStepDead {
		if isAnyKeyJustPressed() {
			g.subStep = tutoStepBase
			if g.hud.levelNum != tutoLearnDeathLevel {
				g.hud.levelNum--
				g.hud.life = g.hud.lifeMax
			}
			g.NextLevel(false, true)
		}
		return
	}

	if isAnyKeyJustPressed() {
		g.subStep = 0
		g.step = stepTitle
		g.inTuto = false
		g.Reset()
	}
}

func (g Game) DrawTuto(screen *ebiten.Image) {
	if g.subStep == tutoStepBase {
		xshift := 0
		yshift := 0
		if g.cameraShake {
			xshift = rand.Intn(globShakeAmplitude)
			yshift = rand.Intn(globShakeAmplitude)
		}
		g.level.Draw(screen, g.animeFrame, xshift, yshift, 1, false)
		g.hud.Draw(screen, true)
	} else if g.subStep == tutoStepDead {
		ebitenutil.DebugPrintAt(screen, "You died!", 30, 30)
		ebitenutil.DebugPrintAt(screen, "But it's fine", 20, 45)
		ebitenutil.DebugPrintAt(screen, "Press any key", 20, 60)
	} else {
		ebitenutil.DebugPrintAt(screen, "Congrats!", 30, 15)
		ebitenutil.DebugPrintAt(screen, "Let's play", 25, 35)
		ebitenutil.DebugPrintAt(screen, "the real game now", 5, 50)
		ebitenutil.DebugPrintAt(screen, "Press any key", 20, 70)
	}
}

func (l *level) SetTutoArea(levelNum int) {

	mid := len(l.area[0]) / 2

	switch levelNum {
	case 1:
		l.area[len(l.area)/2][len(l.area[0])/2] = &levelElement{
			elementType: cactusType,
			posX:        len(l.area[0]) / 2,
			posY:        len(l.area) / 2,
		}
		l.area[len(l.area)/2][len(l.area[0])/2-1] = &levelElement{
			elementType: cactusType,
			posX:        len(l.area[0])/2 - 1,
			posY:        len(l.area) / 2,
		}
		l.area[len(l.area)/2][len(l.area[0])/2+1] = &levelElement{
			elementType: cactusType,
			posX:        len(l.area[0])/2 + 1,
			posY:        len(l.area) / 2,
		}
		l.area[len(l.area)-1][len(l.area[0])/2-2] = &levelElement{
			elementType: cactusType,
			posX:        len(l.area[0])/2 - 2,
			posY:        len(l.area) - 1,
		}
		l.area[len(l.area)-1][len(l.area[0])/2+2] = &levelElement{
			elementType: cactusType,
			posX:        len(l.area[0])/2 + 2,
			posY:        len(l.area) - 1,
		}
		l.area[len(l.area)/2+1][len(l.area[0])/2-2] = &levelElement{
			elementType: cactusType,
			posX:        len(l.area[0])/2 - 2,
			posY:        len(l.area)/2 + 1,
		}
		l.area[len(l.area)/2+2][len(l.area[0])-1] = &levelElement{
			elementType: cactusType,
			posX:        len(l.area[0]) - 1,
			posY:        len(l.area)/2 + 2,
		}
		l.area[0][len(l.area[0])/2-1] = &levelElement{
			elementType: cactusType,
			posX:        len(l.area[0])/2 - 1,
			posY:        0,
		}
	case 2:
		l.area[len(l.area)/2][len(l.area[0])/2] = &levelElement{
			elementType: scorpionType,
			posX:        len(l.area[0]) / 2,
			posY:        len(l.area) / 2,
		}
		l.area[len(l.area)/2][len(l.area[0])/2-1] = &levelElement{
			elementType: cactusType,
			posX:        len(l.area[0])/2 - 1,
			posY:        len(l.area) / 2,
		}
		l.area[len(l.area)/2][len(l.area[0])/2+1] = &levelElement{
			elementType: cactusType,
			posX:        len(l.area[0])/2 + 1,
			posY:        len(l.area) / 2,
		}
		l.area[len(l.area)-1][len(l.area[0])/2-2] = &levelElement{
			elementType: scorpionType,
			posX:        len(l.area[0])/2 - 2,
			posY:        len(l.area) - 1,
		}
		l.area[len(l.area)-1][len(l.area[0])/2+2] = &levelElement{
			elementType: scorpionType,
			posX:        len(l.area[0])/2 + 2,
			posY:        len(l.area) - 1,
		}
		l.area[len(l.area)/2+1][len(l.area[0])/2-2] = &levelElement{
			elementType: cactusType,
			posX:        len(l.area[0])/2 - 2,
			posY:        len(l.area)/2 + 1,
		}
		l.area[len(l.area)/2+2][len(l.area[0])-1] = &levelElement{
			elementType: cactusType,
			posX:        len(l.area[0]) - 1,
			posY:        len(l.area)/2 + 2,
		}
		l.area[0][len(l.area[0])/2-1] = &levelElement{
			elementType: cactusType,
			posX:        len(l.area[0])/2 - 1,
			posY:        0,
		}
	case 3:
		for i := 0; i < len(l.area); i++ {
			l.area[i][len(l.area[0])/2-1] = &levelElement{
				elementType: cactusType,
				posX:        len(l.area[0])/2 - 1,
				posY:        i,
			}
			l.area[i][len(l.area[0])/2+1] = &levelElement{
				elementType: cactusType,
				posX:        len(l.area[0])/2 + 1,
				posY:        i,
			}
		}
		l.area[len(l.area)/2][len(l.area[0])/2] = &levelElement{
			elementType: foodType,
			posX:        len(l.area[0]) / 2,
			posY:        len(l.area) / 2,
		}
	case 4:
		l.addElementAt(mid-1, 0, cactusType)
		l.addElementAt(mid+2, 0, cactusType)
		l.addElementAt(mid, 1, cactusType)
		l.addElementAt(mid+2, 1, cactusType)
		l.addElementAt(mid, 2, cactusType)
		l.addElementAt(mid+2, 2, cactusType)
		l.addElementAt(mid-1, 3, cactusType)
		l.addElementAt(mid, 3, cactusType)
		l.addElementAt(mid+2, 3, cactusType)
		l.addElementAt(mid-2, 4, cactusType)
		l.addElementAt(mid+2, 4, cactusType)
		l.addElementAt(mid-1, 5, cactusType)
		l.addElementAt(mid+1, 5, cactusType)
		l.addElementAt(mid-1, 6, cactusType)
		l.addElementAt(mid+1, 6, cactusType)
		l.addElementAt(mid-1, 7, cactusType)
		l.addElementAt(mid+1, 7, cactusType)
		l.addElementAt(mid, 4, scorpionType)
	case 5:
		l.addElementAt(mid-1, 0, cactusType)
		l.addElementAt(mid+2, 0, cactusType)
		l.addElementAt(mid, 1, cactusType)
		l.addElementAt(mid+2, 1, cactusType)
		l.addElementAt(mid, 2, cactusType)
		l.addElementAt(mid+2, 2, cactusType)
		l.addElementAt(mid, 3, cactusType)
		l.addElementAt(mid+2, 3, cactusType)
		l.addElementAt(mid-1, 4, cactusType)
		l.addElementAt(mid+2, 4, cactusType)
		l.addElementAt(mid-1, 5, cactusType)
		l.addElementAt(mid+1, 5, cactusType)
		l.addElementAt(mid-2, 6, cactusType)
		l.addElementAt(mid+1, 6, cactusType)
		l.addElementAt(mid-1, 7, cactusType)
		l.addElementAt(mid+1, 7, cactusType)
		l.addElementAt(mid, 4, scorpionType)
	case 7:
		l.addElementAt(mid-1, 0, cactusType)
		l.addElementAt(mid+2, 0, cactusType)
		l.addElementAt(mid, 1, cactusType)
		l.addElementAt(mid+2, 1, cactusType)
		l.addElementAt(mid, 2, cactusType)
		l.addElementAt(mid+2, 2, cactusType)
		l.addElementAt(mid, 3, cactusType)
		l.addElementAt(mid+2, 3, cactusType)
		l.addElementAt(mid-1, 4, cactusType)
		l.addElementAt(mid+2, 4, cactusType)
		l.addElementAt(mid-1, 5, cactusType)
		l.addElementAt(mid+1, 5, cactusType)
		l.addElementAt(mid-2, 6, cactusType)
		l.addElementAt(mid+1, 6, cactusType)
		l.addElementAt(mid-1, 7, cactusType)
		l.addElementAt(mid+1, 7, cactusType)
		l.addElementAt(mid, 4, snakeType)
	case 6:
		l.addElementAt(mid-1, 0, cactusType)
		l.addElementAt(mid+3, 0, cactusType)
		l.addElementAt(mid, 1, cactusType)
		l.addElementAt(mid+1, 1, cactusType)
		l.addElementAt(mid+3, 1, cactusType)
		l.addElementAt(mid+1, 2, cactusType)
		l.addElementAt(mid+3, 2, cactusType)
		l.addElementAt(mid, 3, cactusType)
		l.addElementAt(mid+1, 3, cactusType)
		l.addElementAt(mid+3, 3, cactusType)
		l.addElementAt(mid-1, 4, cactusType)
		l.addElementAt(mid+1, 4, cactusType)
		l.addElementAt(mid+3, 4, cactusType)
		l.addElementAt(mid-1, 5, cactusType)
		l.addElementAt(mid+3, 5, cactusType)
		l.addElementAt(mid-1, 6, cactusType)
		l.addElementAt(mid+1, 6, cactusType)
		l.addElementAt(mid+2, 6, cactusType)
		l.addElementAt(mid-1, 7, cactusType)
		l.addElementAt(mid+1, 7, cactusType)
		l.addElementAt(mid, 4, snakeType)
	case 8:
		l.addElementAt(mid, 6, cactusType)
		l.addElementAt(mid-1, 7, cactusType)
		l.addElementAt(mid+1, 7, cactusType)
	case 9:
		for i := 0; i < len(l.area); i++ {
			l.area[i][len(l.area[0])/2-1] = &levelElement{
				elementType: cactusType,
				posX:        len(l.area[0])/2 - 1,
				posY:        i,
			}
			l.area[i][len(l.area[0])/2+1] = &levelElement{
				elementType: cactusType,
				posX:        len(l.area[0])/2 + 1,
				posY:        i,
			}
		}
		l.area[len(l.area)/2][len(l.area[0])/2] = &levelElement{
			elementType: waterType,
			posX:        len(l.area[0]) / 2,
			posY:        len(l.area) / 2,
		}
	case 10:
		l.addElementAt(mid, 6, cactusType)
		l.addElementAt(mid-1, 7, cactusType)
		l.addElementAt(mid+1, 7, cactusType)
	}

}
