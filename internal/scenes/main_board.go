package scenes

import (
	"flappy-go/internal/core"
	"flappy-go/internal/entities"
	"flappy-go/internal/ui"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

func MainScene() *core.Scene {
	scene := core.NewScene(nil, "main_scene", 0)
	gameBoard := gameBoard(scene)
	scene.Add(gameBoard)
	ui := userInterface(scene)
	scene.Add(ui)
	return scene
}

func gameBoard(parent *core.Scene) *core.Scene {
	var speed float32 = 100.0
	// Physics now runs in seconds; use player gravity constant (pixels/s^2)
	scene := core.NewPhysicsScene(parent, "game_board", 0, raylib.Vector2{X: 0, Y: 800}) // Gravity pointing downwards

	// Add the player to the scene
	player := entities.NewPlayer(scene, "red")
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

func userInterface(parent *core.Scene) *core.Scene {
	scene := core.NewScene(parent, "ui", 100)
	// Add the score display to the scene
	scoreDisplay := ui.NewScoreDisplay(scene)
	scene.Add(scoreDisplay)
	//Add The start message
	startMessage := ui.NewInstructionsMessage(scene)
	startMessage.Hide()
	scene.Add(startMessage)
	// Add the game over message
	gameOverMessage := ui.NewGameOverMessage(scene)
	gameOverMessage.Hide()
	scene.Add(gameOverMessage)
	return scene
}
