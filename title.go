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
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) UpdateTitle() {

	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		g.subStep = (g.subStep + 1) % 3
	} else if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		g.subStep = (g.subStep + 2) % 3
	} else if isAnyKeyJustPressed() {
		switch g.subStep {
		case 0:
			g.step = stepLevel
			g.level = initLevel(globLevelX, globLevelY, false, 1)
		case 1:
			g.step = stepTuto
			g.inTuto = true
			g.level = initLevel(globLevelX, globLevelY, true, 1)
		case 2:
			g.step = stepCredits
		}
		g.subStep = 0
	}

}

func (g *Game) DrawTitle(screen *ebiten.Image) {

	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(
		0,
		-5,
	)
	screen.DrawImage(titleImage, &options)

	ebitenutil.DebugPrintAt(screen, "Play", 45, globScreenHeight/2+10)
	ebitenutil.DebugPrintAt(screen, "Learn", 45, globScreenHeight/2+22)
	ebitenutil.DebugPrintAt(screen, "About", 45, globScreenHeight/2+34)

	options = ebiten.DrawImageOptions{}
	options.GeoM.Translate(
		33,
		float64(globScreenHeight/2+12*g.subStep-1+15),
	)
	screen.DrawImage(imageFood, &options)
	//ebitenutil.DebugPrintAt(screen, "o", 35, globScreenHeight/2+12*g.subStep-1+10)

}
