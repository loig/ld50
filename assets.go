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
	"bytes"
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	_ "image/png"
	"log"
)

//go:embed title.png
var titleBytes []byte
var titleImage *ebiten.Image

//go:embed sprites.png
var spritesBytes []byte
var spritesImage *ebiten.Image

var imagePerso [2]*ebiten.Image
var imageBottle *ebiten.Image
var imageCactus *ebiten.Image
var imageSelect *ebiten.Image
var imageSnake [2]*ebiten.Image
var imageScorpion [2]*ebiten.Image
var imageFood *ebiten.Image
var imageGoal *ebiten.Image
var imageLife [2]*ebiten.Image
var imageWater *ebiten.Image
var imageDesert [2]*ebiten.Image
var imageWaterBar [2]*ebiten.Image
var imageDead *ebiten.Image

var imageParticle *ebiten.Image

func loadAssets() {
	var err error

	titleDecoded, _, err := image.Decode(bytes.NewReader(titleBytes))
	if err != nil {
		log.Fatal(err)
	}
	titleImage = ebiten.NewImageFromImage(titleDecoded)

	spritesDecoded, _, err := image.Decode(bytes.NewReader(spritesBytes))
	if err != nil {
		log.Fatal(err)
	}
	spritesImage = ebiten.NewImageFromImage(spritesDecoded)

	imagePerso[0] = spritesImage.SubImage(image.Rect(0, 0, 9, 9)).(*ebiten.Image)
	imagePerso[1] = spritesImage.SubImage(image.Rect(9, 0, 18, 9)).(*ebiten.Image)
	imageBottle = spritesImage.SubImage(image.Rect(18, 0, 27, 9)).(*ebiten.Image)
	imageCactus = spritesImage.SubImage(image.Rect(27, 0, 36, 9)).(*ebiten.Image)
	imageSelect = spritesImage.SubImage(image.Rect(45, 0, 54, 9)).(*ebiten.Image)
	imageSnake[0] = spritesImage.SubImage(image.Rect(0, 9, 9, 18)).(*ebiten.Image)
	imageSnake[1] = spritesImage.SubImage(image.Rect(9, 9, 18, 18)).(*ebiten.Image)
	imageScorpion[0] = spritesImage.SubImage(image.Rect(18, 9, 27, 18)).(*ebiten.Image)
	imageScorpion[1] = spritesImage.SubImage(image.Rect(27, 9, 36, 18)).(*ebiten.Image)
	imageFood = spritesImage.SubImage(image.Rect(36, 9, 45, 18)).(*ebiten.Image)
	imageGoal = spritesImage.SubImage(image.Rect(45, 9, 54, 18)).(*ebiten.Image)
	imageDesert[0] = spritesImage.SubImage(image.Rect(27, 18, 36, 27)).(*ebiten.Image)
	imageDesert[1] = spritesImage.SubImage(image.Rect(36, 18, 45, 27)).(*ebiten.Image)
	imageLife[0] = spritesImage.SubImage(image.Rect(0, 18, 9, 27)).(*ebiten.Image)
	imageLife[1] = spritesImage.SubImage(image.Rect(9, 18, 18, 27)).(*ebiten.Image)
	imageWaterBar[0] = spritesImage.SubImage(image.Rect(0, 27, 45, 36)).(*ebiten.Image)
	imageWaterBar[1] = spritesImage.SubImage(image.Rect(0, 36, 45, 45)).(*ebiten.Image)
	imageWater = spritesImage.SubImage(image.Rect(18, 18, 27, 27)).(*ebiten.Image)
	imageDead = spritesImage.SubImage(image.Rect(45, 36, 54, 45)).(*ebiten.Image)

	imageParticle = spritesImage.SubImage(image.Rect(45, 18, 48, 21)).(*ebiten.Image)
}
