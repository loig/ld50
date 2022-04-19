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
	"math/rand"
	"time"
)

// Game implements the Game interface from ebiten
type Game struct {
	step             int
	subStep          int
	previousLevel    level
	level            level
	hud              hud
	animeFrame       int
	animeStep        int
	transitionFrame  int
	transitionStep   int
	inTuto           bool
	cameraShake      bool
	cameraShakeFrame int
	particles        []*particle
	lastAlive        int
	audio            soundManager
}

// Game steps
const (
	stepTitle int = iota
	stepCredits
	stepTuto
	stepLevel
	stepLevelTransition
	stepDead
)

func initGame() *Game {
	g := Game{}
	g.hud = initHud()
	g.initAudio()
	rand.Seed(time.Now().UnixNano())
	g.lastAlive = -1
	return &g
}

func (g *Game) NextLevel(skip, inTuto bool) {
	if !skip {
		g.hud.NextLevel(inTuto)
	}
	g.previousLevel = g.level
	g.level = initLevel(globLevelX, globLevelY, inTuto, g.hud.levelNum)
}

func (g *Game) Reset() {
	g.hud.Reset()
}
