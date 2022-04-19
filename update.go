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

// Update implements the Update method of the ebiten Game interface
func (g *Game) Update() error {

	g.animeStep++
	if g.animeStep >= globNumAnimeSteps {
		g.animeStep = 0
		g.animeFrame = (g.animeFrame + 1) % 2
	}

	if g.cameraShake {
		g.cameraShakeFrame++
		if g.cameraShakeFrame >= globNumCamShakeFrames {
			g.cameraShake = false
			g.cameraShakeFrame = 0
		}
	}

	g.UpdateParticles()

	g.updateMusic()

	switch g.step {
	case stepTitle:
		g.UpdateTitle()
	case stepCredits:
		g.UpdateCredits()
	case stepTuto:
		g.UpdateTuto()
	case stepLevel:
		g.UpdateLevel()
	case stepDead:
		g.UpdateDead()
	case stepLevelTransition:
		g.UpdateLevelTransition()
	case stepRank:
		g.UpdateRank()
	}

	return nil
}
