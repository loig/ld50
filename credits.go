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
)

func (g *Game) UpdateCredits() {
	if isAnyKeyJustPressed() {
		g.step = stepTitle
	}
}

func (g *Game) DrawCredits(screen *ebiten.Image) {

	ebitenutil.DebugPrintAt(screen, "A game made in", 15, 0)
	ebitenutil.DebugPrintAt(screen, "more or less 48h", 10, 15)
	ebitenutil.DebugPrintAt(screen, "for LD50", 35, 30)
	ebitenutil.DebugPrintAt(screen, "Thanks to Cécile", 10, 50)
	ebitenutil.DebugPrintAt(screen, "for the graphics!", 8, 65)
	ebitenutil.DebugPrintAt(screen, "Code at github.com:", 2, 85)
	ebitenutil.DebugPrintAt(screen, "loig/ld50", 32, 100)
}
