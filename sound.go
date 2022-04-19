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
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"log"
	"math"
	"time"
)

type soundManager struct {
	audioContext *audio.Context
	musicPlayer  *audio.Player
}

// loop the music
func (g *Game) updateMusic() {
	if g.audio.musicPlayer != nil {
		if !g.audio.musicPlayer.IsPlaying() {
			g.audio.musicPlayer.Rewind()
			g.audio.musicPlayer.Play()
		}
	} else {
		var error error
		g.audio.musicPlayer, error = audio.NewPlayer(g.audio.audioContext, infiniteMusic)
		if error != nil {
			log.Panic("Audio problem:", error)
		}
		g.audio.musicPlayer.Play()
	}
	v := g.audio.musicPlayer.Volume()
	if g.step == stepTitle || g.step == stepCredits {
		if v < 1 {
			g.audio.musicPlayer.SetVolume(v + 0.005)
		}
	} else {
		if v > 0.2 {
			g.audio.musicPlayer.SetVolume(v - 0.005)
		}
	}
}

// stop the music
func (g *Game) stopMusic() {
	if g.audio.musicPlayer != nil && g.audio.musicPlayer.IsPlaying() {
		g.audio.musicPlayer.Pause()
	}
}

// load all audio assets
func (g *Game) initAudio() {

	var error error
	g.audio.audioContext = audio.NewContext(44100)

	// music
	sound, error := mp3.Decode(g.audio.audioContext, bytes.NewReader(music))
	if error != nil {
		log.Panic("Audio problem:", error)
	}
	tduration, _ := time.ParseDuration("1m32s")
	duration := tduration.Seconds()
	theBytes := int64(math.Round(duration * 4 * float64(44100)))
	infiniteMusic = audio.NewInfiniteLoop(sound, theBytes)
}
