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
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
)

type hud struct {
	life      int
	lifeMax   int
	water     int
	waterStep int
	waterMax  int
	levelNum  int
	score     int
}

func initHud() (h hud) {
	h.lifeMax = 3
	h.waterMax = 1000
	h.Reset()
	return
}

func (h *hud) Reset() {
	h.life = h.lifeMax
	h.water = h.waterMax
	h.waterStep = 1
	h.levelNum = 1
}

func (h *hud) NextLevel() {
	h.score += (h.water * h.life * h.levelNum) / 100
	h.levelNum++
	h.water = h.waterMax
}

func (h *hud) Update(hurt, food, water bool) (dead bool) {
	if hurt {
		h.life--
		if h.life < 0 {
			h.life = 0
		}
	}
	if food {
		h.life++
		if h.life > h.lifeMax {
			h.life = h.lifeMax
		}
	}

	h.water -= h.waterStep
	if h.water < 0 {
		h.water = 0
	}
	if water {
		h.water += globWaterDrink * h.waterMax / 100
		if h.water > h.waterMax {
			h.water = h.waterMax
		}
	}

	dead = h.life <= 0 || h.water <= 0

	return
}

func (h hud) Draw(screen *ebiten.Image) {

	h.DrawLife(screen)

	h.DrawWater(screen)

	h.DrawLevelNum(screen)

	//h.DrawScore(screen)

}

func (h hud) DrawLife(screen *ebiten.Image) {
	for i := 0; i < h.life; i++ {
		ebitenutil.DrawRect(
			screen,
			float64(i*(globLifeSize+globLifeSep)+globLifePositionX),
			float64(globLifePositionY),
			float64(globLifeSize),
			float64(globLifeSize),
			color.RGBA{255, 0, 0, 255},
		)
	}
}

func (h hud) DrawWater(screen *ebiten.Image) {
	ebitenutil.DrawRect(
		screen,
		float64(globWaterPositionX),
		float64(globWaterPositionY),
		float64(globWaterWidth*h.water/h.waterMax),
		float64(globWaterHeight),
		color.RGBA{0, 255, 255, 255},
	)
}

func (h hud) DrawLevelNum(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, fmt.Sprint("Level ", h.levelNum), globLevelNumPosX, globLevelNumPosY)
}

func (h hud) DrawScore(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, fmt.Sprint("Score: ", h.score), globScorePosX, globScorePosY)
}
