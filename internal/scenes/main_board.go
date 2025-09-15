package scenes

import (
	"simple-go-game/internal/core"
	"simple-go-game/internal/entities"
)

func CreateMainBoard() *core.Scene {
	var speed float32 = 100.0

	scene := core.NewScene()
	// Add the player to the scene
	player := entities.NewPlayer(40, 30, "blue")
	scene.Add(player)
	// Add the background to the scene
	background := entities.NewBackground()
	scene.Add(background)
	// Add the ground to the scene
	ground := entities.NewGround(440, speed)
	ground.Running = true
	scene.Add(ground)
	// Add a pipe gate to the scene
	pipeGate := entities.NewPipeGate(800, 100, 100, speed)
	pipeGate.Running = true
	scene.Add(pipeGate)
	// Additional scene setup can be done here
	return scene
}
