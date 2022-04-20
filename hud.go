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
	"image"
	//"image/color"
)

type hud struct {
	life      int
	lifeMax   int
	water     int
	waterStep int
	waterMax  int
	levelNum  int
	//score     int
}

func initHud() (h hud) {
	h.lifeMax = 3
	h.waterMax = 1000
	h.Reset()
	return
}

func (h *hud) Reset() {
	h.life = h.lifeMax
	h.waterMax = 1000
	h.water = h.waterMax
	h.waterStep = 1
	h.levelNum = 1
}

func (h *hud) NextLevel(inTuto bool) {
	//h.score += (h.water * h.life * h.levelNum) / 100
	h.levelNum++
	if h.waterMax > 150 {
		h.waterMax -= 10
	}
	h.water = h.waterMax
	if inTuto {
		h.waterStep = 2
	}
}

func (h *hud) Update(hurt, food, water, infiniteWater bool) (dead bool) {
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

	if !infiniteWater {
		h.water -= h.waterStep
		if h.water < 0 {
			h.water = 0
		}
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

func (h hud) Draw(screen *ebiten.Image, inTuto, inDeath, deathDone bool) {

	h.DrawLife(screen)

	h.DrawWater(screen)

	h.DrawLevelNum(screen, inTuto, inDeath, deathDone)

	//h.DrawScore(screen)

}

func (h hud) DrawLife(screen *ebiten.Image) {
	for i := 0; i < h.lifeMax; i++ {
		/*
			ebitenutil.DrawRect(
				screen,
				float64(i*(globLifeSize+globLifeSep)+globLifePositionX),
				float64(globLifePositionY),
				float64(globLifeSize),
				float64(globLifeSize),
				color.RGBA{255, 0, 0, 255},
			)
		*/
		options := ebiten.DrawImageOptions{}
		options.GeoM.Translate(
			float64(i*(globLifeSize+globLifeSep)+globLifePositionX),
			float64(globLifePositionY),
		)
		isok := 1
		if i >= h.life {
			isok = 0
		}
		screen.DrawImage(imageLife[isok], &options)
	}
}

func (h hud) DrawWater(screen *ebiten.Image) {
	/*
		ebitenutil.DrawRect(
			screen,
			float64(globWaterPositionX),
			float64(globWaterPositionY),
			float64(globWaterWidth*h.water/h.waterMax),
			float64(globWaterHeight),
			color.RGBA{0, 255, 255, 255},
		)
	*/
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(
		float64(globWaterPositionX),
		float64(globWaterPositionY),
	)
	screen.DrawImage(imageWater, &options)

	options.GeoM.Translate(
		float64(globWaterSep+globAreaCellSize),
		1,
	)
	screen.DrawImage(imageWaterBar[1], &options)
	screen.DrawImage(imageWaterBar[0].SubImage(image.Rect(0, 27, (45*h.water)/h.waterMax, 36)).(*ebiten.Image), &options)
}

func (h hud) DrawLevelNum(screen *ebiten.Image, inTuto, inDeath, deathDone bool) {
	if !inTuto {
		if inDeath {
			ebitenutil.DebugPrintAt(screen, "You died!", globLevelNumPosX, globLevelNumPosY-8)
			if deathDone {
				ebitenutil.DebugPrintAt(screen, fmt.Sprint("(at level ", h.levelNum, ")"), globLevelNumPosX, globLevelNumPosY+2)
			}
		} else {
			ebitenutil.DebugPrintAt(screen, fmt.Sprint("Level ", h.levelNum), globLevelNumPosX, globLevelNumPosY)
		}
	} else {
		if inDeath {
			ebitenutil.DebugPrintAt(screen, "You died!", globLevelNumPosX, globLevelNumPosY-8)
			if deathDone {
				ebitenutil.DebugPrintAt(screen, "Press any key", globLevelNumPosX, globLevelNumPosY+2)
			}
		} else {
			for i, s := range tutoSteps[h.levelNum-1] {
				ebitenutil.DebugPrintAt(screen, s, globLevelNumPosX, globLevelNumPosY-10*(len(tutoSteps[h.levelNum-1])-1-i)+2*(len(tutoSteps[h.levelNum-1])-1))
			}
		}
	}
}

/*
func (h hud) DrawScore(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, fmt.Sprint("Score: ", h.score), globScorePosX, globScorePosY)
}
*/
