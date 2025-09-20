package scenes

import (
	"flappy-go/internal/core"
	"flappy-go/internal/entities"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

func GameBoard() *core.Scene {
	var speed float32 = 100.0

	// Physics now runs in seconds; use player gravity constant (pixels/s^2)
	scene := core.NewPhysicsScene(nil, "game_board", 0, raylib.Vector2{X: 0, Y: 800}) // Gravity pointing downwards
	// Add the score display to the scene
	scoreDisplay := entities.NewScoreDisplay(scene)
	scene.Add(scoreDisplay)
	// Add the player to the scene
	player := entities.NewPlayer(scene, "red", scoreDisplay)
	scene.Add(player)
	// Add the background to the scene
	background := entities.NewBackground(scene, "night")
	scene.Add(background)
	// Add the ground to the scene
	ground := entities.NewGround(scene, speed)
	scene.Add(ground)
	// Add a pipe generator to the scene
	pipeGenerator := entities.NewPipeGateGenerator(scene, speed)
	pipeGenerator.Running = true
	scene.Add(pipeGenerator)
	return scene
}
