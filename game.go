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

// Game implements the Game interface from ebiten
type Game struct {
	level level
	hud   hud
}

func initGame() *Game {
	g := Game{}
	g.level = initLevel(globLevelX, globLevelY)
	g.hud = initHud()
	return &g
}

func (g *Game) NextLevel() {
	g.hud.NextLevel()
	g.level = initLevel(globLevelX, globLevelY)
}
