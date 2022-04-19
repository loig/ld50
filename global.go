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

const (
	globScreenWidth           = 117
	globScreenHeight          = 117
	globWindowWidth           = globScreenWidth * 4
	globWindowHeight          = globScreenHeight * 4
	globTitle                 = "ld50"
	globAreaCellSize          = 9
	globLevelX                = 11
	globLevelY                = 8
	globAreaPositionX         = (globScreenWidth - globLevelX*globAreaCellSize) / 2
	globAreaPositionY         = (globScreenHeight-globLevelY*globAreaCellSize)/2 + 4
	globMoveLeft              = 0
	globMoveUp                = 1
	globMoveRight             = 2
	globMoveDown              = 3
	globLifeSize              = 9
	globLifePositionX         = globAreaPositionX
	globLifePositionY         = globScreenHeight - (globAreaPositionY-4+globLifeSize)/2
	globLifeSep               = 2
	globWaterHeight           = globLifeSize
	globWaterWidth            = 56
	globWaterPositionX        = globAreaPositionX + globAreaCellSize*globLevelX - globWaterWidth
	globWaterPositionY        = globScreenHeight - (globAreaPositionY-4+globWaterHeight)/2
	globWaterSep              = 2
	globWaterDrink            = 25
	globNumSteps              = 5
	globXDivider              = 3
	globYDivider              = 2
	globNumCactus             = 6
	globNumSnakes             = 2
	globNumScorpions          = 2
	globNumFood               = 1
	globNumWater              = 1
	globProbaWaterOnPath      = 8
	globProbaFoodOnPath       = 8
	globProbaMoveSnake        = 3
	globNumMoving             = 3
	globLevelNumPosX          = globAreaPositionX
	globLevelNumPosY          = 5
	globScorePosX             = 0
	globScorePosY             = 0
	globNumAnimeSteps         = 15
	globNumFramesTransStep0   = 60
	globNumFramesTransStep1   = 50
	globNumFramesTransStep2   = 40
	globNumCamShakeFrames     = 10
	globShakeAmplitude        = 5
	globParticleLifespan      = 30
	globMaxNumParticles       = 5
	globNumWaterParticles     = 4
	globWaterParticleLifespan = 30
	globLevelUnlockWater      = 5
	globLevelUnlockScorpion   = 10
	globLevelUnlockSnake      = 15
)
