package main

import (
	"guessgame/game"
	"guessgame/robot"

	"github.com/hudgit2019/leafboot/gameboot"
)

func main() {
	gameboot.StartGame(&game.GuessLogic{}, &robot.GuessRobotLogic{})
}
