package scenes

import (
	"simple-go-game/internal/core"
	"simple-go-game/internal/entities"
)

func GameBoard() *core.Scene {
	var speed float32 = 100.0

	scene := core.NewScene()

	// Add the score display to the scene
	scoreDisplay := entities.NewScoreDisplay()
	scene.Add(scoreDisplay)

	// Add the player to the scene
	player := entities.NewPlayer("red", scoreDisplay)
	scene.Add(player)

	// Add the background to the scene
	background := entities.NewBackground("night")
	scene.Add(background)

	// Add the ground to the scene
	ground := entities.NewGround(speed)
	ground.Running = true
	scene.Add(ground)

	// Add a pipe generator to the scene
	pipeGenerator := entities.NewPipeGateGenerator(speed)
	pipeGenerator.Running = true
	scene.Add(pipeGenerator)

	return scene
}
