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
	"math/rand"
)

type particle struct {
	posX, posY  float64
	vX, vY      float64
	alpha       float64
	age, maxAge int
	dead        bool
	ptype       int // 0 sand, 1 water, 2 watermelon
}

func (p *particle) Draw(screen *ebiten.Image) {
	if !p.dead {
		options := ebiten.DrawImageOptions{}
		options.GeoM.Translate(
			float64(p.posX),
			float64(p.posY),
		)
		red := 1.0
		green := 1.0
		blue := 1.0
		if p.ptype == 1 {
			red = 0.667
			green = 0.761
			blue = 0.871
		}
		if p.ptype == 2 {
			red = 0.882
			green = 0.384
			blue = 0.263
		}
		options.ColorM.Scale(red, green, blue, p.alpha)
		screen.DrawImage(imageParticle, &options)
	}
}

func (g *Game) DrawParticles(screen *ebiten.Image) {
	for i := 0; i <= g.lastAlive; i++ {
		g.particles[i].Draw(screen)
	}
}

func (p *particle) Update() {
	if !p.dead {
		p.posX += p.vX
		p.posY += p.vY
		p.alpha = 1 - float64(p.age)/float64(p.maxAge)
		p.age++
		if p.age > p.maxAge {
			p.dead = true
		}
	}
}

func (g *Game) UpdateParticles() {
	for i := 0; i <= g.lastAlive; i++ {
		g.particles[i].Update()
		if g.particles[i].dead {
			g.particles[i], g.particles[g.lastAlive] = g.particles[g.lastAlive], g.particles[i]
			g.lastAlive--
			i--
		}
	}
}

func (g *Game) AddParticle(x, y float64, ptype int) {
	maxAge := globParticleLifespan
	if ptype == 1 {
		maxAge = rand.Intn(20) + globWaterParticleLifespan
	}
	p := particle{
		posX: x, posY: y,
		vX: 0, vY: -0.05,
		alpha: 1,
		age:   0, maxAge: maxAge,
		dead:  false,
		ptype: ptype,
	}
	g.lastAlive++
	if g.lastAlive < len(g.particles) {
		*(g.particles[g.lastAlive]) = p
	} else {
		g.particles = append(g.particles, &p)
	}
}

func (g *Game) AddParticlesOnGrid(fromX, fromY, toX, toY int) {

	maxNum := fromX - toX + fromY - toY
	if maxNum < 0 {
		maxNum = -maxNum
	}

	num := maxNum
	if maxNum > globMaxNumParticles {
		num = globMaxNumParticles
	}

	if fromX != toX {
		inc := 1
		if fromX < toX {
			inc = -1
		}
		for i := num; i >= 1; i-- {
			toX += inc
			for j := num - i; j < num; j++ {
				g.AddOneParticleOnGrid(toX, fromY)
			}
		}
		return
	}

	if fromY != toY {
		inc := 1
		if fromY < toY {
			inc = -1
		}
		for i := num; i >= 1; i-- {
			toY += inc
			for j := num - i; j < num; j++ {
				g.AddOneParticleOnGrid(fromX, toY)
			}
		}
	}
}

func (g *Game) AddOneParticleOnGrid(x, y int) {
	g.AddParticle(
		float64(globAreaPositionX+x*globAreaCellSize+3+rand.Intn(3)-1),
		float64(globAreaPositionY+y*globAreaCellSize+3-rand.Intn(3)),
		0,
	)
}

func (g *Game) AddWaterFoodParticles(x, y int, water bool) {
	ptype := 1
	if !water {
		ptype = 2
	}
	for i := 0; i < globNumWaterParticles; i++ {
		g.AddParticle(
			float64(globAreaPositionX+x*globAreaCellSize+rand.Intn(7)),
			float64(globAreaPositionY+y*globAreaCellSize-3-rand.Intn(7)),
			ptype,
		)
	}
}
