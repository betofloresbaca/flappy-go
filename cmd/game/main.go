// Main entry point for the simple-go-game.
// This file initializes and starts the game.
package main

import (
	_ "embed"
	"simple-go-game/internal/core"
	"simple-go-game/internal/scenes"
)

func main() {
	// Create a new game instance
	g := core.NewGame(860, 540, "Simple Go Game", 60)

	// Ensure cleanup when the program exits
	defer g.Cleanup()

	// Initialize the game window and resources
	g.Initialize()

	// Create and set the main scene
	mainScene := scenes.CreateMainBoard()
	g.SetScene(mainScene)

	// Start the main game loop
	g.Run()
}
