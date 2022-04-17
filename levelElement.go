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
	//"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	//"image/color"
)

type levelElement struct {
	elementType           int
	posX                  int
	posY                  int
	doNotMoveInGeneration bool
}

const (
	nilType int = iota
	persoType
	cactusType
	snakeType
	scorpionType
	waterType
	foodType
)

func (e levelElement) Draw(screen *ebiten.Image) {
	if e.elementType != nilType {
		//var c color.Color
		var img *ebiten.Image
		var xshift, yshift int
		switch e.elementType {
		case persoType:
			//c = color.RGBA{255, 0, 255, 255}
			img = imagePerso[0]
			yshift = -3
		case cactusType:
			//c = color.RGBA{0, 255, 0, 255}
			img = imageCactus
			yshift = -4
		case snakeType:
			//c = color.RGBA{0, 0, 255, 255}
			img = imageSnake[0]
		case scorpionType:
			//c = color.RGBA{255, 255, 0, 255}
			img = imageScorpion[0]
		case waterType:
			//c = color.RGBA{0, 255, 255, 255}
			img = imageBottle
		case foodType:
			//c = color.RGBA{255, 0, 0, 255}
			img = imageFood
		}
		/*
			ebitenutil.DrawRect(
				screen,
				float64(e.posX*globAreaCellSize+globAreaPositionX),
				float64(e.posY*globAreaCellSize+globAreaPositionY),
				float64(globAreaCellSize), float64(globAreaCellSize),
				c,
			)
		*/
		options := ebiten.DrawImageOptions{}
		options.GeoM.Translate(
			float64(e.posX*globAreaCellSize+globAreaPositionX+xshift),
			float64(e.posY*globAreaCellSize+globAreaPositionY+yshift),
		)
		screen.DrawImage(img, &options)

	} /*else {
		ebitenutil.DrawRect(
			screen,
			float64(e.posX*globAreaCellSize+globAreaPositionX),
			float64(e.posY*globAreaCellSize+globAreaPositionY),
			float64(globAreaCellSize), float64(globAreaCellSize),
			color.RGBA{100, 100, 100, 255},
		)
	}*/
}

func (e levelElement) IsMovable() bool {
	return e.elementType == persoType || e.elementType == snakeType || e.elementType == scorpionType
}

func (e levelElement) IsCollectible() bool {
	return e.elementType == waterType || e.elementType == foodType
}
