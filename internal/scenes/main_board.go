package scenes

import (
	"simple-go-game/internal/core"
	"simple-go-game/internal/entities"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

func GameBoard() *core.Scene {
	var speed float32 = 100.0

	// Physics now runs in seconds; use player gravity constant (pixels/s^2)
	scene := core.NewPhysicsScene(0, raylib.Vector2{X: 0, Y: 800}) // Gravity pointing downwards
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
