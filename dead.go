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
)

func (g *Game) UpdateDead() {
	if isAnyKeyJustPressed() {
		g.step = stepTitle
		g.Reset()
	}
}

func (g *Game) DrawDead(screen *ebiten.Image) {

	ebitenutil.DebugPrintAt(screen, "You died!", 30, 30)
	ebitenutil.DebugPrintAt(screen, fmt.Sprint("(at level ", g.hud.levelNum, ")"), 20, 45)

}
