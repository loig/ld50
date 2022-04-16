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
)

type level struct {
	area     [][]*levelElement
	movable  []*levelElement
	selected int
}

func initLevel(sizeX, sizeY int) (l level) {
	l.area = make([][]*levelElement, sizeY)
	for i := 0; i < sizeY; i++ {
		l.area[i] = make([]*levelElement, sizeX)
	}
	l.area[0][0] = &levelElement{elementType: persoType, posX: 0, posY: 0}
	l.area[2][1] = &levelElement{elementType: cactusType, posX: 1, posY: 2}
	l.area[3][2] = &levelElement{elementType: snakeType, posX: 2, posY: 3}
	l.area[4][3] = &levelElement{elementType: scorpionType, posX: 3, posY: 4}
	l.area[5][4] = &levelElement{elementType: waterType, posX: 4, posY: 5}
	l.area[6][5] = &levelElement{elementType: foodType, posX: 5, posY: 6}

	l.movable = make([]*levelElement, 0)
	var selectedPos int
	for i := 0; i < sizeY; i++ {
		for j := 0; j < sizeX; j++ {
			if l.area[i][j] != nil {
				if l.area[i][j].elementType == persoType {
					l.selected = selectedPos
				}
				switch l.area[i][j].elementType {
				case persoType, snakeType, scorpionType:
					l.movable = append(l.movable, l.area[i][j])
					selectedPos++
				}
			}
		}
	}
	return
}

func (l *level) Update() (hurt bool) {

	if inpututil.IsKeyJustPressed(ebiten.KeyTab) {
		l.ChangeSelected()
		return
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		hurt = l.MoveSelected(globMoveLeft)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		hurt = l.MoveSelected(globMoveRight)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		hurt = l.MoveSelected(globMoveUp)
	} else if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		hurt = l.MoveSelected(globMoveDown)
	}

	return
}

func (l *level) ChangeSelected() {
	l.selected = (l.selected + 1) % len(l.movable)
}

func (l *level) MoveSelected(direction int) (hurt bool) {
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

	i += moveY
	j += moveX

	for i >= 0 && i < len(l.area) && j >= 0 && j < len(l.area[0]) {
		if l.area[i][j] != nil {
			hurt =
				(toMove.elementType == persoType && l.area[i][j].elementType == scorpionType) ||
					(toMove.elementType == snakeType && l.area[i][j].elementType == persoType)
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

	l.area[i][j] = toMove

	return
}

func (l level) Draw(screen *ebiten.Image) {

	l.DrawBackground(screen)

	for _, line := range l.area {
		for _, element := range line {
			if element != nil {
				element.Draw(screen)
			}
		}
	}

	l.DrawSelected(screen)
}

func (l level) DrawBackground(screen *ebiten.Image) {
	ebitenutil.DrawRect(
		screen,
		float64(globAreaPositionX),
		float64(globAreaPositionY),
		float64(globAreaCellSize*len(l.area[0])),
		float64(globAreaCellSize*len(l.area)),
		color.RGBA{55, 55, 55, 255},
	)
}

func (l level) DrawSelected(screen *ebiten.Image) {
	ebitenutil.DrawRect(
		screen,
		float64(l.movable[l.selected].posX*globAreaCellSize+globAreaPositionX),
		float64(l.movable[l.selected].posY*globAreaCellSize+globAreaPositionY),
		float64(globAreaCellSize),
		float64(globAreaCellSize),
		color.RGBA{255, 255, 255, 255},
	)
}
