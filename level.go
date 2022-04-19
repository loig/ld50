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
	"image/color"
	"math/rand"
)

type level struct {
	area         [][]*levelElement
	movable      []*levelElement
	perso        *levelElement
	goalX, goalY int
	selected     int
	num          int
}

func initLevel(sizeX, sizeY int, inTuto bool, levelNum int) (l level) {

	l.num = levelNum

	l.area = make([][]*levelElement, sizeY)
	for i := 0; i < sizeY; i++ {
		l.area[i] = make([]*levelElement, sizeX)
	}

	persoX := sizeX / 2
	persoY := sizeY - 1
	l.perso = &levelElement{elementType: persoType, posX: persoX, posY: persoY}
	l.area[persoY][persoX] = l.perso

	l.goalX = persoX
	l.goalY = 0

	if inTuto {

		l.SetTutoArea(levelNum)

	} else {

		withSnakes := levelNum > globLevelUnlockSnake
		withScorpions := levelNum > globLevelUnlockScorpion
		withWater := levelNum > globLevelUnlockWater
		withFood := withWater && (withSnakes || withScorpions)

		// gen level
		l.GenArea(withSnakes, withScorpions, withFood, withWater)
	}

	l.movable = make([]*levelElement, 0)
	var selectedPos int
	for i := 0; i < sizeY; i++ {
		for j := 0; j < sizeX; j++ {
			if l.area[i][j] != nil {
				if l.area[i][j].elementType == persoType {
					l.selected = selectedPos
				}
				if l.area[i][j].IsMovable() {
					l.movable = append(l.movable, l.area[i][j])
					selectedPos++
				}
			}
		}
	}

	return
}

func (l *level) Update(allowTab bool) (hurt, food, water, finished, skip bool, fromX, toX, fromY, toY int, hasMoved, waterp, foodp bool) {

	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		skip = true
		return
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyTab) && allowTab {
		l.ChangeSelected()
		return
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		hasMoved = true
		hurt, food, water, fromX, toX, fromY, toY, waterp, foodp = l.MoveSelected(globMoveLeft)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		hasMoved = true
		hurt, food, water, fromX, toX, fromY, toY, waterp, foodp = l.MoveSelected(globMoveRight)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		hasMoved = true
		hurt, food, water, fromX, toX, fromY, toY, waterp, foodp = l.MoveSelected(globMoveUp)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		hasMoved = true
		hurt, food, water, fromX, toX, fromY, toY, waterp, foodp = l.MoveSelected(globMoveDown)
	}

	finished = l.perso.posX == l.goalX && l.perso.posY == l.goalY

	return
}

func (l *level) ChangeSelected() {
	l.selected = (l.selected + 1) % len(l.movable)
}

func (l *level) MoveSelected(direction int) (hurt, food, water bool, fromX, toX, fromY, toY int, foodp, waterp bool) {
	var moveX, moveY int
	switch direction {
	case globMoveUp:
		moveY = -1
	case globMoveDown:
		moveY = 1
	case globMoveLeft:
		moveX = -1
	case globMoveRight:
		moveX = 1
	default:
		return
	}

	toMove := l.movable[l.selected]
	i := toMove.posY
	j := toMove.posX

	fromX = toMove.posX
	fromY = toMove.posY

	i += moveY
	j += moveX

	for i >= 0 && i < len(l.area) && j >= 0 && j < len(l.area[0]) {
		if l.area[i][j] != nil && l.area[i][j].elementType != nilType {
			hurt =
				(toMove.elementType == persoType && l.area[i][j].elementType == scorpionType) ||
					(toMove.elementType == snakeType && l.area[i][j].elementType == persoType)
			if l.area[i][j].IsCollectible() {
				foodp = l.area[i][j].elementType == foodType
				waterp = l.area[i][j].elementType == waterType
				if toMove.elementType == persoType {
					food = l.area[i][j].elementType == foodType
					water = l.area[i][j].elementType == waterType
				}
				l.area[i][j] = nil
				i += moveY
				j += moveX
			}
			break
		}

		i += moveY
		j += moveX
	}

	i -= moveY
	j -= moveX

	l.area[toMove.posY][toMove.posX] = nil

	toMove.posX = j
	toMove.posY = i

	toX = toMove.posX
	toY = toMove.posY

	l.area[i][j] = toMove

	return
}

func (l level) Draw(screen *ebiten.Image, frame int, xshift, yshift int, alpha float64, inTransition bool) {

	l.DrawBackground(screen, xshift, yshift, inTransition)

	l.DrawGoal(screen, xshift, yshift, alpha)

	if !inTransition {
		l.DrawSelected(screen, alpha)
	}

	for _, line := range l.area {
		for _, element := range line {
			if element != nil {
				element.Draw(screen, frame, xshift, yshift, alpha)
			}
		}
	}

}

func (l level) DrawBackground(screen *ebiten.Image, xshift, yshift int, inTransition bool) {

	for i := 0; i < len(l.area); i++ {
		for j := 0; j < len(l.area[0]); j++ {
			options := ebiten.DrawImageOptions{}
			options.GeoM.Translate(
				float64(j*globAreaCellSize+globAreaPositionX+xshift),
				float64(i*globAreaCellSize+globAreaPositionY+yshift),
			)
			screen.DrawImage(imageDesert[(i*len(l.area[0])+j+l.num)%2], &options)
		}
	}

	if inTransition && yshift < 0 {
		ebitenutil.DrawRect(
			screen,
			float64(globAreaPositionX),
			float64(0),
			float64(globAreaCellSize*len(l.area[0])),
			float64(globAreaPositionY),
			color.RGBA{0, 0, 0, 255},
		)
	}

	if inTransition && yshift > 0 {
		ebitenutil.DrawRect(
			screen,
			float64(globAreaPositionX),
			float64(globAreaPositionY+globAreaCellSize*len(l.area)),
			float64(globAreaCellSize*len(l.area[0])),
			float64(globScreenHeight),
			color.RGBA{0, 0, 0, 255},
		)
	}
}

func (l level) DrawGoal(screen *ebiten.Image, xshift, yshift int, alpha float64) {
	/*
		ebitenutil.DrawRect(
			screen,
			float64(l.goalX*globAreaCellSize+globAreaPositionX),
			float64(l.goalY*globAreaCellSize+globAreaPositionY),
			float64(globAreaCellSize),
			float64(globAreaCellSize),
			color.RGBA{155, 55, 55, 255},
		)
	*/
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(
		float64(l.goalX*globAreaCellSize+globAreaPositionX+xshift),
		float64(l.goalY*globAreaCellSize+globAreaPositionY+yshift),
	)
	options.ColorM.Scale(1, 1, 1, alpha)
	screen.DrawImage(imageGoal, &options)
}

func (l level) DrawSelected(screen *ebiten.Image, alpha float64) {
	/*
		ebitenutil.DrawRect(
			screen,
			float64(l.movable[l.selected].posX*globAreaCellSize+globAreaPositionX),
			float64(l.movable[l.selected].posY*globAreaCellSize+globAreaPositionY),
			float64(globAreaCellSize),
			float64(globAreaCellSize),
			color.RGBA{255, 255, 255, 255},
		)
	*/
	options := ebiten.DrawImageOptions{}
	options.GeoM.Translate(
		float64(l.movable[l.selected].posX*globAreaCellSize+globAreaPositionX),
		float64(l.movable[l.selected].posY*globAreaCellSize+globAreaPositionY),
	)
	options.ColorM.Scale(1, 1, 1, alpha)
	screen.DrawImage(imageSelect, &options)
}

func (g *Game) UpdateLevel() {
	hurt, food, water, finished, skip, fromX, toX, fromY, toY, hasMoved, foodp, waterp := g.level.Update(true)
	if skip {
		g.NextLevel(skip, false)
	}
	if hurt {
		g.cameraShake = true
	}
	if hasMoved {
		g.AddParticlesOnGrid(fromX, fromY, toX, toY)
	}
	if waterp {
		g.AddWaterFoodParticles(toX, toY, true)
	}
	if foodp {
		g.AddWaterFoodParticles(toX, toY, false)
	}
	dead := g.hud.Update(hurt, food, water, false)
	if dead {
		g.step = stepDead
	}
	if finished {
		g.NextLevel(false, false)
		g.step = stepLevelTransition
	}
}

func (g Game) DrawLevel(screen *ebiten.Image) {
	xshift := 0
	yshift := 0
	if g.cameraShake {
		xshift = rand.Intn(globShakeAmplitude)
		yshift = rand.Intn(globShakeAmplitude)
	}
	g.level.Draw(screen, g.animeFrame, xshift, yshift, 1, false)
	g.hud.Draw(screen, false)
}
