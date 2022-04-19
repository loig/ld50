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
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type rank struct {
	pos, numPos             int
	firstn, secondn, thirdn string
	firstl, secondl, thirdl int
	ok                      bool
}

const (
	rankStepFirstLetter int = iota
	rankStepSecondLetter
	rankStepThirdLetter
	rankStepRequest
	rankStepDisplay
)

func (g *Game) UpdateRank() {

	if g.subStep > rankStepThirdLetter {
		if isAnyKeyJustPressed() {
			g.step = stepTitle
			g.subStep = 0
			g.Reset()
			return
		}
	}

	if g.subStep == rankStepRequest {
		select {
		case r := <-g.therank:
			g.finalrank = r
			g.subStep++
		default:
		}
	}

	if g.subStep <= rankStepThirdLetter {
		if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
			g.IncreaseLetter(g.subStep, 1)
		} else if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
			g.IncreaseLetter(g.subStep, -1)
		} else if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
			g.subStep = (g.subStep + 1) % 3
		} else if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
			g.subStep = (g.subStep + 2) % 3
		} else if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			go getRank(g.GetName(), g.hud.levelNum, g.therank)
			g.subStep = rankStepRequest
		}
	}

}

func (g *Game) IncreaseLetter(pos, inc int) {
	g.name[pos] = (g.name[pos] + 26 + inc) % 26
}

func (g *Game) GetName() string {
	return fmt.Sprint(string(g.name[0]+65), string(g.name[1]+65), string(g.name[2]+65))
}

func (g *Game) DrawRank(screen *ebiten.Image) {

	ebitenutil.DebugPrintAt(screen, "Ranking", globLevelNumPosX-5, globLevelNumPosY)

	switch g.subStep {
	case rankStepRequest:
		ebitenutil.DebugPrintAt(screen, "Waiting for server", globLevelNumPosX, globLevelNumPosY+20)
	case rankStepDisplay:
		if !g.finalrank.ok {
			ebitenutil.DebugPrintAt(screen, "Can't reach Server", globLevelNumPosX-5, globLevelNumPosY+25)
			return
		}
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("You ranked ", g.finalrank.pos), globLevelNumPosX, globLevelNumPosY+25)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("among ", g.finalrank.numPos), globLevelNumPosX, globLevelNumPosY+35)

		ebitenutil.DebugPrintAt(screen, fmt.Sprint("1. ", g.finalrank.firstn, " lvl ", g.finalrank.firstl), globLevelNumPosX, globLevelNumPosY+60)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("2. ", g.finalrank.secondn, " lvl ", g.finalrank.secondl), globLevelNumPosX, globLevelNumPosY+70)
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("3. ", g.finalrank.thirdn, " lvl ", g.finalrank.thirdl), globLevelNumPosX, globLevelNumPosY+80)
	case rankStepFirstLetter, rankStepSecondLetter, rankStepThirdLetter:
		ebitenutil.DebugPrintAt(screen, fmt.Sprint("Name: ", g.GetName()), globLevelNumPosX, globLevelNumPosY+25)
		ebitenutil.DebugPrintAt(screen, "Use arrows to", globLevelNumPosX, globLevelNumPosY+50)
		ebitenutil.DebugPrintAt(screen, "write your name.", globLevelNumPosX, globLevelNumPosY+60)
		ebitenutil.DebugPrintAt(screen, "Press Enter when", globLevelNumPosX, globLevelNumPosY+75)
		ebitenutil.DebugPrintAt(screen, "You are done.", globLevelNumPosX, globLevelNumPosY+85)
	}

}

func getRank(name string, level int, c chan rank) {
	response, err := http.PostForm("http://localhost:8081/", url.Values{
		"uname": {name},
		"level": {fmt.Sprint(level)},
	})
	if err != nil {
		c <- rank{ok: false}
		return
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		c <- rank{ok: false}
		return
	}

	parts := strings.Split(string(body), ":")
	if len(parts) != 5 {
		c <- rank{ok: false}
	}

	pos, err := strconv.Atoi(parts[0])
	if err != nil {
		c <- rank{ok: false}
		return
	}

	firstn := parts[1][1:4]
	firstl, err := strconv.Atoi(parts[1][5 : len(parts[1])-1])
	if err != nil {
		c <- rank{ok: false}
		return
	}

	secondn := parts[2][1:4]
	secondl, err := strconv.Atoi(parts[2][5 : len(parts[2])-1])
	if err != nil {
		c <- rank{ok: false}
		return
	}

	thirdn := parts[3][1:4]
	thirdl, err := strconv.Atoi(parts[3][5 : len(parts[3])-1])
	if err != nil {
		c <- rank{ok: false}
		return
	}

	numPos, err := strconv.Atoi(parts[4])
	if err != nil {
		c <- rank{ok: false}
		return
	}

	c <- rank{
		ok:  true,
		pos: pos, numPos: numPos,
		firstn: firstn, firstl: firstl,
		secondn: secondn, secondl: secondl,
		thirdn: thirdn, thirdl: thirdl,
	}
	return
}
