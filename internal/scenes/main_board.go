package scenes

import (
	"flappy-go/internal/core"
	"flappy-go/internal/entities"
	"flappy-go/internal/ui"

	raylib "github.com/gen2brain/raylib-go/raylib"
)

func MainScene() *core.Scene {
	scene := core.NewScene(nil, "main_scene", []string{}, 0)
	ui := userInterface(scene)
	scene.Add(ui)
	// Add game controller to the game board
	gameController := entities.NewGameController(scene, gameBoard(scene))
	scene.Add(gameController)
	return scene
}

func gameBoard(parent *core.Scene) func() *core.Scene {
	return func() *core.Scene {
		var speed float32 = 100.0
		// Physics now runs in seconds; use player gravity constant (pixels/s^2)
		scene := core.NewPhysicsScene(parent, "game_board", []string{}, 0, raylib.Vector2{X: 0, Y: 800}) // Gravity pointing downwards
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
		// Add the player to the scene
		player := entities.NewPlayer(scene, "blue")
		scene.Add(player)
		return scene
	}
}

func userInterface(parent *core.Scene) *core.Scene {
	scene := core.NewScene(parent, "ui", []string{}, 100)
	// Add the score display to the scene
	scoreDisplay := ui.NewScoreDisplay(scene)
	scene.Add(scoreDisplay)
	//Add The start message
	startMessage := ui.NewStartMessage(scene)
	startMessage.Hide()
	scene.Add(startMessage)
	// Add the game over message
	gameOverMessage := ui.NewGameOverMessage(scene)
	gameOverMessage.Hide()
	scene.Add(gameOverMessage)
	return scene
}
